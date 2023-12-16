package middleware

import (
	mlogger "github.com/1kovalevskiy/snippetbox/internal/middleware/logger"
	"github.com/1kovalevskiy/snippetbox/pkg/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewMiddleware(router *chi.Mux, l logger.Interface) error {
	router.Use(mlogger.New(l))
	router.Use(middleware.RedirectSlashes)
	return nil
}
