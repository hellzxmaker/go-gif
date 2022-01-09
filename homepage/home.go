package homepage

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const message = "Go Gif Urself API"

type Handlers struct {
	logger *log.Logger
}

// GET /
// Returns a root message
// Status Code: 200 OK
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(message))
}

func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Println("HOME: Request processed in: ", time.Since(startTime))
		next(w, r)
	}
}

func (h *Handlers) SetupRoutes(r *mux.Router) {
	r.HandleFunc("/", h.Logger(h.Home)).Methods("GET")
}

// Constructor
func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}
