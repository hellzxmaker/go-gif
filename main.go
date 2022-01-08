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
	// Create a new logger instance
	logger := log.New(
		os.Stdout,
		"go-gif-api: ",
		log.LstdFlags|log.Lshortfile,
	)

	h := homepage.NewHandlers(logger)
	p := players.NewHandlers(logger)
	g := games.NewHandlers(logger)

	// TODO: Add TLS config support cloudflare expose go to internet
	// mux := http.NewServeMux()
	r := mux.NewRouter()
	h.SetupRoutes(r)
	p.SetupRoutes(r)
	g.SetupRoutes(r)

	srv := server.NewServer(r)
	logger.Println("Starting server on port 8080...")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start on port %v: ", err)
	}
}
