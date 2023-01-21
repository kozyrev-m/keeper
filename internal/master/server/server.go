package server

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/kozyrev-m/keeper/internal/master/config"
	"github.com/kozyrev-m/keeper/internal/master/server/httpserver"
	"github.com/kozyrev-m/keeper/internal/master/storage"
)

// StartServer starts server.
func StartServer(config *config.Config) error {
	store, err := storage.CreateStore(config)
	if err != nil {
		return err
	}

	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))

	srv := httpserver.New(store, sessionStore)

	httpServer := &http.Server{
		Addr: config.Address,
		Handler: srv,
	}

	return httpServer.ListenAndServe()
}