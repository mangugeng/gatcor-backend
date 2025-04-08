package handler

import (
	"encoding/json"
	"net/http"
)

// Handler handles all requests
func Handler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "Hello from Gatcor API!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
