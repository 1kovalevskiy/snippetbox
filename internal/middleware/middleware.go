package middleware

import (
	"log/slog"

	"github.com/1kovalevskiy/snippetbox/internal/middleware/logger"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/chi/v5"
)

func NewMiddleware(router *chi.Mux, l *slog.Logger) {
	router.Use(middleware.RedirectSlashes)
	router.Use(logger.New(l))
}