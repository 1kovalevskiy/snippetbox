package v1

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"strconv"
)


type routes struct {
	logger *slog.Logger
}

func NewRoutes(l *slog.Logger) *routes {
	return &routes{l}
}

func (rt *routes) home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		rt.serverError(w, err)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		rt.serverError(w, err)
	}
}
 
func (rt *routes) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		rt.notFound(w)
        return
	}
	if id < 1 {
		rt.logger.Warn("Unexpected id", "id", id)
		rt.notFound(w)
        return

	}
	fmt.Fprintf(w, "Отображение выбранной заметки с ID %d...", id)
}
 
func (rt *routes) createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Форма для создания новой заметки..."))
}
