package main

import (
	"log"
	"net/http"
)

func allowOnlyGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		if method != "GET" {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			log.Printf("%s %q %v", r.Method, r.URL.String(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func checkQueryString(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if q.Get("cat") == "" || q.Get("pub") == "" {
			http.Error(w, "Query string must not be empty", http.StatusBadRequest)
			log.Printf("%s %q %v", r.Method, r.URL.String(), http.StatusBadRequest)
			return
		}

		log.Printf("GET %q", r.URL.String())
		next.ServeHTTP(w, r)
	})
}
