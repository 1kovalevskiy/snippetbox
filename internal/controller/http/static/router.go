package static

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(router *chi.Mux, _ *slog.Logger) {
	fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static/")))
	router.Handle("/static/*", fileServer)
}
