package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func NewServer(r *mux.Router) *http.Server {
	// Create the server with some default props
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return srv
}
