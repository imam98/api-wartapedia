package main

import (
	"encoding/json"
	es "github.com/elastic/go-elasticsearch/v8"
	"github.com/gorilla/mux"
	"github.com/imam98/api-wartapedia/pkg/domain"
	"github.com/imam98/api-wartapedia/pkg/domain/entity"
	"github.com/imam98/api-wartapedia/pkg/infrastructure/news_fetcher"
	"github.com/imam98/api-wartapedia/pkg/infrastructure/persistence/elasticsearch"
	"github.com/imam98/api-wartapedia/pkg/listing"
	"github.com/imam98/api-wartapedia/pkg/querying"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"time"
)

type response struct {
	Status  int           `json:"status"`
	Message string        `json:"message,omitempty"`
	Data    []entity.News `json:"data,omitempty"`
}

func newsHandler(w http.ResponseWriter, r *http.Request) {
	fetcher := news_fetcher.NewFetcher()
	service := listing.NewService(fetcher)
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")

	q := r.URL.Query()
	flags := genRepoFlags(q.Get("cat"), q.Get("src"))
	data, err := service.GetNews(flags)
	if err != nil {
		handleError(w, http.StatusNotFound, "News source not found")
		return
	}

	resp := response{
		Status: http.StatusOK,
		Data:   data,
	}
	if err := encoder.Encode(resp); err != nil {
		handleError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func sourceListHandler(w http.ResponseWriter, r *http.Request) {
	fetcher := news_fetcher.NewFetcher()
	service := listing.NewService(fetcher)
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")

	q := r.URL.Query()
	catFlag := genRepoFlags(q.Get("cat"), "")
	source, err := service.GetSourcesFromCategory(catFlag)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := struct {
		Status int      `json:"status"`
		Data   []string `json:"data"`
	}{
		Status: http.StatusOK,
		Data:   source,
	}
	if err := encoder.Encode(resp); err != nil {
		handleError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func main() {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC822Z,
	}
	logger := zerolog.New(output).With().Timestamp().Logger()

	esClient, err := es.NewDefaultClient()
	if err != nil {
		logger.Error().Err(err).Msg("Failed to create es client")
	}
	repo := elasticsearch.NewRepository(elasticsearch.Config{
		Client:    esClient,
		IndexName: "wartapedia",
	})
	queryer := querying.NewService(repo)

	r := mux.NewRouter()
	r.HandleFunc("/api/news", newsHandler)
	r.HandleFunc("/api/list/source", sourceListHandler)
	r.HandleFunc("/api/search", searchQueryHandler(queryer))
	handler := http.Handler(r)
	handler = checkQueryString(handler)
	handler = allowOnlyGet(handler)

	logger.Info().Msg("Service listening to port 3000")
	http.ListenAndServe(":3000", handler)
}

func searchQueryHandler(service domain.QueryService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)
		w.Header().Set("content-type", "application/json")

		q := r.URL.Query()
		searchQuery := q.Get("q")
		results, err := service.Query(searchQuery, 50)
		if err != nil {
			handleError(w, http.StatusInternalServerError, err.Error())
			return
		}

		resp := response{
			Status: http.StatusOK,
			Data:   results,
		}
		if err := encoder.Encode(resp); err != nil {
			handleError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
}

func genRepoFlags(category, source string) domain.RepoFlag {
	var catFlag domain.RepoFlag
	switch category {
	case "nasional":
		catFlag = domain.CAT_NASIONAL
	case "dunia":
		catFlag = domain.CAT_DUNIA
	case "tekno":
		catFlag = domain.CAT_TEKNO
	default:
		catFlag = domain.RepoFlag(0)
	}

	var srcFlag domain.RepoFlag
	switch source {
	case "antaranews":
		srcFlag = domain.ANTARANEWS
	case "bbc":
		srcFlag = domain.BBC
	case "detik":
		srcFlag = domain.DETIK
	case "okezone":
		srcFlag = domain.OKEZONE
	case "republika":
		srcFlag = domain.REPUBLIKA
	default:
		srcFlag = domain.RepoFlag(0)
	}

	return catFlag | srcFlag
}

func handleError(w http.ResponseWriter, status int, message string) {
	responseErr := response{
		Status:  status,
		Message: message,
	}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(responseErr)
}

func makeLog(req *http.Request) *zerolog.Logger {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC822Z,
	}
	logger := zerolog.New(output).With().
		Timestamp().
		Str("method", req.Method).
		Str("uri", req.URL.String()).
		Str("ip", req.RemoteAddr).
		Str("referer", req.Referer()).
		Logger()

	return &logger
}
