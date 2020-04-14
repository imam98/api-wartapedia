package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/imam98/api-wartapedia/pkg/infrastructure/news_fetcher"
	"github.com/imam98/api-wartapedia/pkg/listing"
	"github.com/imam98/api-wartapedia/pkg/news"
	"log"
	"net/http"
)

type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    []news.News `json:"data,omitempty"`
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
	r := mux.NewRouter()
	r.HandleFunc("/api/news", newsHandler)
	r.HandleFunc("/api/list/source", sourceListHandler)
	handler := http.Handler(r)
	handler = checkQueryString(handler)
	handler = allowOnlyGet(handler)
	log.Println("Service listening to port 3000")
	http.ListenAndServe(":3000", handler)
}

func genRepoFlags(category, source string) news.RepoFlag {
	var catFlag news.RepoFlag
	switch category {
	case "nasional":
		catFlag = news.CAT_NASIONAL
	case "dunia":
		catFlag = news.CAT_DUNIA
	case "tekno":
		catFlag = news.CAT_TEKNO
	default:
		catFlag = news.RepoFlag(0)
	}

	var srcFlag news.RepoFlag
	switch source {
	case "antaranews":
		srcFlag = news.ANTARANEWS
	case "bbc":
		srcFlag = news.BBC
	case "detik":
		srcFlag = news.DETIK
	case "okezone":
		srcFlag = news.OKEZONE
	case "republika":
		srcFlag = news.REPUBLIKA
	default:
		srcFlag = news.RepoFlag(0)
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
