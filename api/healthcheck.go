package api

import (
	"encoding/json"
	"net/http"
)

// HealthMessage type
type HealthMessage struct {
	Message string `json:"message"`
}

// HealthCheckFunc responds with 200
func HealthCheckFunc(w http.ResponseWriter, r *http.Request) {
	msg := HealthMessage{Message: "I'm alive!"}
	ToJSON, _ := json.Marshal(msg)
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(ToJSON)
}
