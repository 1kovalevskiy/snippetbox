package v1

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (rt *routes) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	rt.logger.Error(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (rt *routes) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (rt *routes) notFound(w http.ResponseWriter) {
	rt.clientError(w, http.StatusNotFound)
}