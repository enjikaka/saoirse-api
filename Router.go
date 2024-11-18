package main

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

type router struct {
	*httprouter.Router
}

func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		context.Set(r, "params", ps)
		h.ServeHTTP(w, r)
	}
}

func (r *router) Get(path string, handler http.Handler) {
	r.GET(path, wrapHandler(handler))
}

// NewRouter : Creates a new router instance
func NewRouter() *router {
	return &router{httprouter.New()}
}
