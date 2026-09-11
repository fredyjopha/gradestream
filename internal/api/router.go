package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/ingest", handleIngest)
	return r
}
