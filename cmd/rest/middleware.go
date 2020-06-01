package main

import (
	"net/http"
)

func allowOnlyGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		w.Header().Set("Content-Type", "application/json")
		if method != "GET" {
			handleError(w, http.StatusBadRequest, "Bad request method")
			makeLog(r).Error().
				Int("status", http.StatusBadRequest).
				Msg("Bad request method")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func checkQueryString(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/api/domain" {
			if q.Get("cat") == "" || q.Get("src") == "" {
				handleError(w, http.StatusBadRequest, "Query params must not be empty")
				makeLog(r).Error().
					Int("status", http.StatusBadRequest).
					Msg("Empty query params")
				return
			}
		} else if r.URL.Path == "/api/list/source" {
			if q.Get("cat") == "" {
				handleError(w, http.StatusBadRequest, "Query params must not be empty")
				makeLog(r).Error().
					Int("status", http.StatusBadRequest).
					Msg("Empty query params")
				return
			}
		}

		makeLog(r).Info().
			Msg("Incoming request")
		next.ServeHTTP(w, r)
	})
}
