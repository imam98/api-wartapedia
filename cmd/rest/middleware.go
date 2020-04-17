package main

import (
	"github.com/rs/zerolog"
	"net/http"
)

func allowOnlyGet(next http.Handler, logger zerolog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		w.Header().Set("Content-Type", "application/json")
		if method != "GET" {
			handleError(w, http.StatusBadRequest, "Bad request method")
			logger.Error().
				Str("method", method).
				Str("uri", r.URL.String()).
				Int("status", http.StatusBadRequest).
				Str("ip", r.RemoteAddr).
				Str("referer", r.Referer()).
				Msg("Bad request method")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func checkQueryString(next http.Handler, logger zerolog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/api/news" {
			if q.Get("cat") == "" || q.Get("src") == "" {
				handleError(w, http.StatusBadRequest, "Query params must not be empty")
				logger.Error().
					Str("method", "GET").
					Str("uri", r.URL.String()).
					Int("status", http.StatusBadRequest).
					Str("ip", r.RemoteAddr).
					Str("referer", r.Referer()).
					Msg("Empty query params")
				return
			}
		} else if r.URL.Path == "/api/list/source" {
			if q.Get("cat") == "" {
				handleError(w, http.StatusBadRequest, "Query params must not be empty")
				logger.Error().
					Str("method", "GET").
					Str("uri", r.URL.String()).
					Int("status", http.StatusBadRequest).
					Str("ip", r.RemoteAddr).
					Str("referer", r.Referer()).
					Msg("Empty query params")
				return
			}
		}

		logger.Info().
			Str("method", "GET").
			Str("uri", r.URL.String()).
			Str("ip", r.RemoteAddr).
			Str("referer", r.Referer()).
			Msg("Incoming request")
		next.ServeHTTP(w, r)
	})
}
