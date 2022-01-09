package games

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/hellzxmaker/go-gif/players"
)

// TODO: Remove me!
var fake_game_info = GameInfo{
	Round:         1,
	ActivePlayers: 2,
}

// TODO: Remove me!
var fake_game_session = GameSession{
	GameJoinUid: "12345",
	GameJoinUrl: "http://localhost:8080/game/12345",
	Players: []players.Player{
		{
			PlayerUid:  "12345",
			PlayerName: "Fake Player",
			ActiveSessions: []string{
				"12345",
				"54321",
			},
		},
		{
			PlayerUid:  "54321",
			PlayerName: "Fake Player 2",
			ActiveSessions: []string{
				"12345",
				"54321",
			},
		},
	},
	GameInfo: fake_game_info,
	Questions: []Question{
		{
			Q:    "What is the capital of France?",
			Type: "normal",
		},
		{
			Q:    "What is the capital of Germany?",
			Type: "normal",
		},
		{
			Q:    "What is the capital of Italy?",
			Type: "normal",
		},
		{
			Q:    "What is the capital of Spain?",
			Type: "normal",
		},
		{
			Q:    "What is the capital of the United States?:imageLinkHere",
			Type: "image",
		},
	},
}

type Handlers struct {
	logger *log.Logger
}

// GET /games/[id]
// Return a game session by id
// Status Code: 200 OK
// Status Code: 404 Not Found
func (h *Handlers) GetGameSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	game_id := vars["id"]

	// Fake logic to interpret the ID
	if fake_game_session.GameJoinUid == game_id {
		payload, err := json.Marshal(fake_game_session)
		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write(payload)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			defer h.logger.Println("games.GetGameSession: Error marshalling JSON: ", err)
		}
	}
}

// GET /games/[id]/info
// Returns the GameInfo object for a game
// Status Code: 200 OK
// Status Code: 404 Not Found
func (h *Handlers) GetGameSessionGameInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	game_id := vars["id"]

	// Fake logic to interpret the ID
	if fake_game_session.GameJoinUid == game_id {
		payload, err := json.Marshal(fake_game_session.GameInfo)
		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write(payload)
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			defer h.logger.Println("games.GetGameSessionGameInfo: Error marshalling JSON: ", err)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		defer h.logger.Println("games.GetGameSessionGameInfo: Game not found: ", game_id)
	}
}

// GET /games
// Returns a list of all game sessions
// Status Code: 200 OK
// Status Code: 404 NOT FOUND
// TODO: Implmenent properly
func (h *Handlers) GetGameSessions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	payload, err := json.Marshal(fake_game_session)
	if err == nil {
		w.Write(payload)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		defer h.logger.Println("games.GetGameSessions: Error marshalling JSON: ", err)
	}
}

// POST /games
// Create a new game session
// Status Code: 201 CREATED
// Status Code: 400 BAD REQUEST
func (h *Handlers) CreateGameSession(w http.ResponseWriter, r *http.Request) {
	// parse the incoming request
	var payload GameSession
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		err1 := json.Unmarshal(body, &payload)
		if err1 == nil {
			w.WriteHeader(http.StatusCreated)
			defer h.logger.Println("games.CreateGameSession: Received payload: ", payload)
		}
	} else {
		// set the status code to failed
		w.WriteHeader(http.StatusBadRequest)
		defer h.logger.Println("games.CreateGameSession: Error reading request body: ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
}

// GET /games/[id]/players
// Get the players in a game
// Status Code: 200 OK
// Status Code: 404 GAME NOT FOUND
// Status Code: 400 BAD REQUEST
func (h *Handlers) GetGameSessionPlayers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	game_session_id := vars["id"]

	// Fake logic to interpret the ID
	if fake_game_session.GameJoinUid == game_session_id {
		payload, err := json.Marshal(fake_game_session.Players)
		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write(payload)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			defer h.logger.Println("games.GetGameSessionPlayers: Error marshalling JSON: ", err)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		defer h.logger.Println("games.GetGameSessionPlayers: Game not found: ", game_session_id)
	}
}

// GET /games/[id]/questions
// Get the questions in a game
// Status Code: 200 OK
// Status Code: 404 GAME NOT FOUND
// Status Code: 400 BAD REQUEST
func (h *Handlers) GetGameSessionQuestions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	game_session_id := vars["id"]

	// Fake logic to interpret the ID
	if fake_game_session.GameJoinUid == game_session_id {
		payload, err := json.Marshal(fake_game_session.Questions)
		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write(payload)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			defer h.logger.Println("games.GetGameSessionQuestions: Error marshalling JSON: ", err)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		defer h.logger.Println("games.GetGameSessionQuestions: Game not found: ", game_session_id)
	}
}

func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Println("games: Request processed in: ", time.Since(startTime))
		next(w, r)
	}
}

func (h *Handlers) SetupRoutes(r *mux.Router) {
	r.HandleFunc("/games", h.Logger(h.GetGameSessions)).Methods("GET")
	r.HandleFunc("/games/{id}", h.Logger(h.GetGameSession)).Methods("GET")
	r.HandleFunc("/games/{id}/info", h.Logger(h.GetGameSessionGameInfo)).Methods("GET")
	r.HandleFunc("/games", h.Logger(h.CreateGameSession)).Methods("POST")
	r.HandleFunc("/games/{id}/players", h.Logger(h.GetGameSessionPlayers)).Methods("GET")
	r.HandleFunc("/games/{id}/questions", h.Logger(h.GetGameSessionQuestions)).Methods("GET")
}

// Constructor
func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}
