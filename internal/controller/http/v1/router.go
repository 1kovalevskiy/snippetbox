package v1

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
)



func NewRouter(router *chi.Mux, l *slog.Logger) {
	r := NewRoutes(l)
	router.Get("/", r.home)
	router.Get("/snippet", r.showSnippet)
	router.Post("/snippet/create", r.createSnippet)
}
