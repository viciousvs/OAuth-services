package home

import (
	"encoding/json"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	val := r.Context().Value("userUUID").(string)
	_ = json.NewEncoder(w).Encode(map[string]string{"uuid": val})
}
