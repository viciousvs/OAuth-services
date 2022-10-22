package blog

import (
	"github.com/viciousvs/OAuth-services/gateway/utils/httpUtils"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	httpUtils.NewErrorResponse(w, http.StatusNotFound, "not found")
	return
}
