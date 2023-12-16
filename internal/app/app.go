package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/1kovalevskiy/snippetbox/config"
	static "github.com/1kovalevskiy/snippetbox/internal/controller/http/static"
	v1 "github.com/1kovalevskiy/snippetbox/internal/controller/http/v1"
	middleware "github.com/1kovalevskiy/snippetbox/internal/middleware"
	"github.com/1kovalevskiy/snippetbox/internal/usecase"
	repo "github.com/1kovalevskiy/snippetbox/internal/usecase/repo_sqlite"
	"github.com/1kovalevskiy/snippetbox/pkg/httpserver"
	"github.com/1kovalevskiy/snippetbox/pkg/logger"
	sqlite_ "github.com/1kovalevskiy/snippetbox/pkg/sqlite"

	"github.com/go-chi/chi/v5"
)

func Run(cfg *config.Config) {
	l := logger.NewLogger()

	sqlite, err := sqlite_.New(cfg.SQL.URL, cfg.SQL.Timeout)
	if err != nil {
		l.Error("app - Run - sql.New", err.Error())
		return
	}
	defer sqlite.Close()

	snippetUseCase := usecase.New(
		repo.New(sqlite),
	)

	router := chi.NewRouter()

	middleware.NewMiddleware(router, l)
	err = v1.NewRouter(router, l, snippetUseCase)
	if err != nil {
		l.Error(err.Error())
		return
	}
	static.NewRouter(router)

	httpServer := httpserver.New(router, httpserver.Port(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err).Error())
	}

	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err).Error())
	}
}
