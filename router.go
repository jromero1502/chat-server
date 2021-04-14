package server

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) AddRoute(path string, method string, handler http.HandlerFunc, middlewares ...Middleware) {
	if r.rules[path] == nil {
		r.rules[path] = make(map[string]http.HandlerFunc)
	}

	h := handler
	for _, m := range middlewares {
		h = m(h)
	}

	r.rules[path][method] = h
}

func (r *Router) FindHandler(path string, method string) http.HandlerFunc {
	handler, exists := r.rules[path][method]

	if !exists {
		return nil
	}

	return handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler := r.FindHandler(request.URL.Path, request.Method)
	if handler == nil {
		w.WriteHeader(404)
	} else {
		handler(w, request)
	}
}
