package main

import (
	"context"
	"log"
	"net/http"

	"github.com/rodrigodiez/zorro/pkg/generator/uuid"
)

type contextKey int

const (
	xRequestID contextKey = iota
)

func (a *app) requestID(f http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rID := r.Header.Get("X-Request-ID")

		if rID == "" {
			rID = uuid.NewV4().Generate("")
		}

		ctx := context.WithValue(r.Context(), xRequestID, rID)
		w.Header().Set("X-Request-ID", rID)

		f.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *app) logger(f http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("--> %s - %s", r.Method, r.URL.String())
		f.ServeHTTP(w, r)
	})
}
