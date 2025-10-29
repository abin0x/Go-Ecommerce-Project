package util

import (
	"encoding/json"
	"net/http"
)

// senData sends a JSON response with the given data and status code
func SendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
