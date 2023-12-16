package v1

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/1kovalevskiy/snippetbox/internal/entity"
)

func (rt *routes) home(w http.ResponseWriter, r *http.Request) {
	s, err := rt.usecase.GetTenLatestSnippet(r.Context())
	if err != nil {
		rt.serverError(w, r, err)
		return
	}

	rt.render(w, r, "home.page.tmpl", &templateData{
		Snippets: s,
	})
}

func (rt *routes) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		rt.notFound(w, r)
		return
	}
	if id < 1 {
		rt.notFound(w, r)
		return

	}

	s, err := rt.usecase.GetSnippet(r.Context(), id)
	if err != nil {
		if errors.Is(err, entity.ErrNoRecord) {
			rt.notFound(w, r)
		} else {
			rt.serverError(w, r, err)
		}
		return
	}

	rt.render(w, r, "show.page.tmpl", &templateData{
		Snippet: s,
	})
}

func (rt *routes) createSnippetPage(w http.ResponseWriter, r *http.Request) {
	rt.render(w, r, "create.page.tmpl", &templateData{})
}

func (rt *routes) createSnippet(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	content := r.FormValue("content")
	expires := r.FormValue("expires")
	num, err := strconv.Atoi(expires)
	if err != nil || num < 1 {
		rt.clientError(w, r, 400, "Время жизни заметки должно быть положительным числом")
		return
	}

	request := entity.SnippetCreate{
		Title:   title,
		Content: content,
		Expires: expires,
	}
	id, err := rt.usecase.CreateSnippet(r.Context(), request)
	if err != nil {
		rt.serverError(w, r, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
