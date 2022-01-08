package main

import (
	"log"
	"os"

	"github.com/hellzxmaker/go-gif/games"
	"github.com/hellzxmaker/go-gif/homepage"
	"github.com/hellzxmaker/go-gif/players"
	"github.com/hellzxmaker/go-gif/server"

	mux "github.com/gorilla/mux"
)

func main() {
	// TODO: Get logger props from env var
	// TODO: Get SSL props from env var
	// TODO: Get server props from env var
	// Create a new logger instance
	logger := log.New(
		os.Stdout,
		"go-gif-api: ",
		log.LstdFlags|log.Lshortfile,
	)

	// Create handlers for each API endpoint
	h := homepage.NewHandlers(logger)
	p := players.NewHandlers(logger)
	g := games.NewHandlers(logger)

	// TODO: Add TLS config support cloudflare expose go to internet
	// Create the router and setup the routes
	r := mux.NewRouter()
	h.SetupRoutes(r)
	p.SetupRoutes(r)
	g.SetupRoutes(r)

	// Create the server
	srv := server.NewServer(r)
	logger.Println("Starting server on port 8080...")

	// Main server loop
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start on port %v: ", err)
	}
}
