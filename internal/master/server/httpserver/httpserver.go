// Package httpserver provides http-server for storing private information.
//
// server - lightweight implementation of our service.
// This implementation can only process the request that enters it.
// It can also implement the http.Handler interface (for http.ListenAndServe()).
// This will allow you to directly pipe into the http.ListenAndServe function.
package httpserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Server - lightweight server implementation for flexibility and independence.
type Server struct {
	router *mux.Router
}

// New creates a http-server instance.
func New() *Server {
	s := &Server {
		router: mux.NewRouter(),
	}

	s.configureRouter()

	return s
}

// ServeHTTP delegates the work of the router that runs ServeHTTP.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// configureRouter prepares endpoints and middlewares.
func (s *Server) configureRouter() {
	// s.router.HandleFunc("/users", s.handleCreateUser()).Methods("POST")
}