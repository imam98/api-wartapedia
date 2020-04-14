package main

import (
	"log"
	"net/http"
)

func allowOnlyGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		w.Header().Set("Content-Type", "application/json")
		if method != "GET" {
			handleError(w, http.StatusBadRequest, "Bad request method")
			log.Printf("%s %q %v", r.Method, r.URL.String(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func checkQueryString(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/api/news" {
			if q.Get("cat") == "" || q.Get("src") == "" {
				handleError(w, http.StatusBadRequest, "Query params must not be empty")
				log.Printf("%s %q %v", r.Method, r.URL.String(), http.StatusBadRequest)
				return
			}
		} else if r.URL.Path == "/api/list/source" {
			if q.Get("cat") == "" {
				handleError(w, http.StatusBadRequest, "Query params must not be empty")
				log.Printf("%s %q %v", r.Method, r.URL.String(), http.StatusBadRequest)
				return
			}
		}

		log.Printf("GET %q", r.URL.String())
		next.ServeHTTP(w, r)
	})
}
