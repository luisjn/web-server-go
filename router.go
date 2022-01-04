package main

import (
	"net/http"
)

// Router contains what routes can handle the server
type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, ok := r.rules[path]
	handler, ok2 := r.rules[path][method]
	return handler, ok, ok2
}

func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	handler, ok, ok2 := r.FindHandler(req.URL.Path, req.Method)
	if !ok {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	if !ok2 {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(rw, req)
}
