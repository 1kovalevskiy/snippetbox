package v1

import (
	"github.com/go-chi/chi/v5"

	"github.com/1kovalevskiy/snippetbox/internal/usecase"
	"github.com/1kovalevskiy/snippetbox/pkg/logger"
)

func NewRouter(router *chi.Mux, l logger.Interface, s usecase.Snippet) error {
	r, err := NewRoutes(l, s)
	if err != nil {
		return err
	}
	router.Get("/", r.home)
	router.Get("/snippet", r.showSnippet)
	router.Get("/snippet/create", r.createSnippetPage)
	router.Post("/snippet/create", r.createSnippet)
	return nil
}
