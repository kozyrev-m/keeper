package server

import (
	"net/http"

	"github.com/kozyrev-m/keeper/internal/master/server/httpserver"
)

// StartServer starts server.
func StartServer() error {
	srv := httpserver.New()

	httpServer := &http.Server{
		Addr: ":8081",
		Handler: srv,
	}

	return httpServer.ListenAndServe()
}