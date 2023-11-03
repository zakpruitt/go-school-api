package handlers

import (
	"encoding/json"
	"net/http"
)

// JSONResponse sends a JSON-encoded response with the given status code.
func JSONResponse[T any](w http.ResponseWriter, data T, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		JSONErrorResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// JSONErrorResponse sends a JSON-encoded error message and status code.
func JSONErrorResponse(w http.ResponseWriter, message string, status int) {
	JSONResponse(w, map[string]string{"error": message}, status)
}
