package v1

import (
	"fmt"
	"html/template"
	"net/http"
	"runtime/debug"

	"github.com/1kovalevskiy/snippetbox/internal/usecase"
	"github.com/1kovalevskiy/snippetbox/pkg/logger"
)

type routes struct {
	logger    logger.Interface
	usecase   usecase.Snippet
	templates map[string]*template.Template
}

func NewRoutes(l logger.Interface, s usecase.Snippet) (*routes, error) {
	templates, err := newTemplateCache("./ui/html")
	if err != nil {
		return nil, err
	}
	return &routes{l, s, templates}, nil
}

type Errors struct {
	Code    int
	Message string
}

func (rt *routes) serverError(w http.ResponseWriter, r *http.Request, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	rt.logger.Error(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (rt *routes) clientError(w http.ResponseWriter, r *http.Request, status int, message string) {
	rt.render(w, r, "4xx.page.tmpl", &templateData{Errors: &Errors{Code: status, Message: message}})
}

func (rt *routes) notFound(w http.ResponseWriter, r *http.Request) {
	rt.render(w, r, "404.page.tmpl", &templateData{})
}

func (rt *routes) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := rt.templates[name]
	if !ok {
		rt.serverError(w, r, fmt.Errorf("Шаблон %s не существует!", name))
		return
	}

	err := ts.Execute(w, td)
	if err != nil {
		rt.serverError(w, r, err)
	}
}
