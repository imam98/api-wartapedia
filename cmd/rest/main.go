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

func newsHandler(w http.ResponseWriter, r *http.Request) {
	fetcher := news_fetcher.NewFetcher()
	service := listing.NewService(fetcher)

	q := r.URL.Query()
	flags := genSourceFlags(q.Get("cat"), q.Get("pub"))
	data, err := service.GetNews(flags)
	if err != nil {
		http.Error(w, "News source not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/news", newsHandler)
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
