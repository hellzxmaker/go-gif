package players

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Player struct {
	// Unique ID of for the player. Pulled from their local app instance
	PlayerUid string `json:"playerUid"`
	// The player's name
	PlayerName string `json:"playerName"`
	// A list of active session IDs
	ActiveSessions []string `json:"activeSessions"`
}

// TODO: Remove me!
var fake_player = Player{
	PlayerUid:  "12345",
	PlayerName: "Fake Player",
	ActiveSessions: []string{
		"12345",
		"54321",
	},
}

type Handlers struct {
	logger *log.Logger
}

// GET /players
// Returns a list of all players
// Status Code: 200 OK
func (h *Handlers) GetPlayers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	payload, err := json.Marshal(fake_player)
	if err == nil {
		w.Write(payload)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		defer h.logger.Println("players.Players: Error marshalling JSON: ", err)
	}
}

// POST /players
// Creates a new player
// Status Code: 201 CREATED
// Status Code: 400 BAD REQUEST
func (h *Handlers) PostPlayers(w http.ResponseWriter, r *http.Request) {
	// parse the incoming request
	var payload Player
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		err1 := json.Unmarshal(body, &payload)
		if err1 == nil {
			w.WriteHeader(http.StatusCreated)
			defer h.logger.Println("players.PostPlayers: Received payload: ", payload)
		}
	} else {
		// set the status code to failed
		w.WriteHeader(http.StatusBadRequest)
		defer h.logger.Println("players.PostPlayers: Error reading request body: ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
}

func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Println("players: Request processed in: ", time.Since(startTime))
		next(w, r)
	}
}

func (h *Handlers) SetupRoutes(r *mux.Router) {
	r.HandleFunc("/players", h.Logger(h.GetPlayers)).Methods("GET")
	r.HandleFunc("/players", h.Logger(h.PostPlayers)).Methods("POST")
}

// Constructor
func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}
