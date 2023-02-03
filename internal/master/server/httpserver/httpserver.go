// Package httpserver provides http-server for storing private information.
//
// server - lightweight implementation of our service.
// This implementation can only process the request that enters it.
// It can also implement the http.Handler interface (for http.ListenAndServe()).
// This will allow you to directly pipe into the http.ListenAndServe function.
package httpserver

import (
	"mime/multipart"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/kozyrev-m/keeper/internal/master/model/datamodel"
	"github.com/kozyrev-m/keeper/internal/master/model/usermodel"
	"github.com/kozyrev-m/keeper/internal/master/storage/store/sqlstore"
)

// Store is store iterface.
type Store interface {
	CreateUser(*usermodel.User) error
	FindUserByLogin(string) (*usermodel.User, error)
	FindUserByID(int) (*usermodel.User, error)

	CreateDataRecord(sqlstore.Content) error
	FindTextsByOwner(int) ([]datamodel.Text, error)

	CreateFile(int, string, string, multipart.File) error
	GetFileList(int) ([]datamodel.File, error)
}

// Server - lightweight server implementation for flexibility and independence.
type Server struct {
	router       *mux.Router
	store        Store
	sessionStore sessions.Store
}

// New creates a http-server instance.
func New(store Store, sessionStore sessions.Store) *Server {
	s := &Server{
		router:       mux.NewRouter(),
		store:        store,
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
	s.router.HandleFunc("/users", s.handleRegisterUser()).Methods(http.MethodPost)
	s.router.HandleFunc("/sessions", s.handleCreateSession()).Methods(http.MethodPost)

	private := s.router.PathPrefix("/private").Subrouter()
	private.Use(s.authenticateUser)
	private.HandleFunc("/whoami", s.handleWhoami()).Methods(http.MethodGet)

	private.HandleFunc("/text", s.handleCreateText()).Methods(http.MethodPost)
	private.HandleFunc("/text", s.handleGetTexts()).Methods(http.MethodGet)

	private.HandleFunc("/pair", s.handleAddPair()).Methods(http.MethodPost)

	private.HandleFunc("/file", s.handleSaveFile()).Methods(http.MethodPost)
	private.HandleFunc("/file", s.handleFileList()).Methods(http.MethodGet)
	private.HandleFunc("/file/{filename}", s.handleDownloadFile()).Methods(http.MethodGet)
}
