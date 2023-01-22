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
	"github.com/gorilla/sessions"
	"github.com/kozyrev-m/keeper/internal/master/model"
)

// Store is store iterface.
type Store interface {
	CreateUser(*model.User) error
	FindUserByLogin(string) (*model.User, error)
}

// Server - lightweight server implementation for flexibility and independence.
type Server struct {
	router *mux.Router
	store  Store
	sessionStore sessions.Store
}

// New creates a http-server instance.
func New(store Store, sessionStore sessions.Store) *Server {
	s := &Server {
		router: mux.NewRouter(),
		store: store,
		sessionStore: sessionStore,
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
	s.router.HandleFunc("/users", s.handleRegisterUser()).Methods("POST")
	s.router.HandleFunc("/sessions", s.handleCreateSession()).Methods("POST")
}