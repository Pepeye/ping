package ping

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/lucsky/cuid"
)

// Ping type
type Ping struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

// F handles http requests
func F(w http.ResponseWriter, r *http.Request) {
	id := cuid.New()
	m := Ping{id, "Hello World!", time.Now()}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}
