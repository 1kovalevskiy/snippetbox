package static

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(router *chi.Mux) error {
	fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static/")))
	router.Handle("/static/*", fileServer)
	return nil
}
