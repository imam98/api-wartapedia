package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func allowOnlyGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		if method != "GET" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)

			resp := response{
				Status:  http.StatusBadRequest,
				Message: "Bad request method",
			}
			json.NewEncoder(w).Encode(resp)

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
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)

			resp := response{
				Status:  http.StatusBadRequest,
				Message: "Query params must not be empty",
			}
			json.NewEncoder(w).Encode(resp)
			log.Printf("%s %q %v", r.Method, r.URL.String(), http.StatusBadRequest)
			return
		}

		log.Printf("GET %q", r.URL.String())
		next.ServeHTTP(w, r)
	})
}
