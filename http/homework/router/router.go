package router

import (
	"net/http"
	"site/middleware"
)

type Route struct {
	Path        string
	Handler     http.Handler
	Middlewares []func(http.Handler) http.Handler
}

type Router struct {
	routes []Route
}

// NewRouter creates a new router
func NewRouter() *Router {
	return &Router{}
}

// Handle registers a new route with middlewares
func (r *Router) Handle(path string, handler http.HandlerFunc, middlewares ...func(http.Handler) http.Handler) {
	wrapped := middleware.Chain(handler, middlewares...)
	r.routes = append(r.routes, Route{
		Path:    path,
		Handler: wrapped,
	})
}

// ServeHTTP implements http.Handler
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.Path == req.URL.Path {
			route.Handler.ServeHTTP(w, req)
			return
		}
	}
	http.NotFound(w, req)
}
