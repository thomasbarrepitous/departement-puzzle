package utils

import (
	"encoding/json"
	"net/http"
)

// RespondJSON sends a JSON response with the specified status code and data
func JSONRespond(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
