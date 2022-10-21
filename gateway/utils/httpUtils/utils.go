package httpUtils

import (
	"encoding/json"
	"net/http"
)

func NewErrorResponse(w http.ResponseWriter, statusCode int, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(map[string]string{"message": msg})
}
