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
	flags := genSourceFlags(q.Get("cat"), q.Get("pub"))
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

func publisherListHandler(w http.ResponseWriter, r *http.Request) {
	fetcher := news_fetcher.NewFetcher()
	service := listing.NewService(fetcher)
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")

	q := r.URL.Query()
	catFlag := genSourceFlags(q.Get("cat"), "")
	publishers, err := service.GetPublishersFromCategory(catFlag)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := struct {
		Status int      `json:"status"`
		Data   []string `json:"data"`
	}{
		Status: http.StatusOK,
		Data:   publishers,
	}
	if err := encoder.Encode(resp); err != nil {
		handleError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/news", newsHandler)
	r.HandleFunc("/api/list/publisher", publisherListHandler)
	handler := http.Handler(r)
	handler = checkQueryString(handler)
	handler = allowOnlyGet(handler)
	log.Println("Service listening to port 3000")
	http.ListenAndServe(":3000", handler)
}

func genSourceFlags(category, publisher string) news.SourceFlag {
	var catFlag news.SourceFlag
	switch category {
	case "nasional":
		catFlag = news.CAT_NASIONAL
	case "dunia":
		catFlag = news.CAT_DUNIA
	case "tekno":
		catFlag = news.CAT_TEKNO
	default:
		catFlag = news.SourceFlag(0)
	}

	var pubFlag news.SourceFlag
	switch publisher {
	case "antaranews":
		pubFlag = news.ANTARANEWS
	case "bbc":
		pubFlag = news.BBC
	case "detik":
		pubFlag = news.DETIK
	case "okezone":
		pubFlag = news.OKEZONE
	case "republika":
		pubFlag = news.REPUBLIKA
	default:
		pubFlag = news.SourceFlag(0)
	}

	return catFlag | pubFlag
}

func handleError(w http.ResponseWriter, status int, message string) {
	responseErr := response{
		Status:  status,
		Message: message,
	}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(responseErr)
}
