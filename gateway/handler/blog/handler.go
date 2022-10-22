package blog

import (
	"encoding/json"
	"github.com/viciousvs/OAuth-services/gateway/utils/httpUtils"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	val, ok := r.Context().Value("userUUID").(string)
	if !ok {
		httpUtils.NewErrorResponse(w, http.StatusUnprocessableEntity, "unprocessable entity")
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(map[string]string{"uuid": val})
	if err != nil {
		httpUtils.NewErrorResponse(w, http.StatusUnprocessableEntity, "unprocessable entity")
		return
	}
}
