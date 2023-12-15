package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/1kovalevskiy/snippetbox/config"
	middleware "github.com/1kovalevskiy/snippetbox/internal/middleware"
	v1 "github.com/1kovalevskiy/snippetbox/internal/controller/http/v1"
	static "github.com/1kovalevskiy/snippetbox/internal/controller/http/static"
	"github.com/1kovalevskiy/snippetbox/pkg/logger"

	"github.com/go-chi/chi/v5"
)


func Run(cfg *config.Config) {
	l := logger.NewLogger()

	router := chi.NewRouter()

	middleware.NewMiddleware(router, l)
	v1.NewRouter(router, l)
	static.NewRouter(router, l)

    l.Info(fmt.Sprintf("Запуск веб-сервера на http://127.0.0.1:%s\n", cfg.HTTP.Port))
	addr := fmt.Sprintf(":%s", cfg.HTTP.Port)
    err := http.ListenAndServe(addr, router)
    log.Fatal(err)
}