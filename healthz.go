package rainpole

import (
	"encoding/json"
	"net/http"
)

// Typings of response body from /healthz
type HealthStatusResponse struct {
	StatusCode    int    `json:"code"`
	StatusMessage string `json:"status"`
}

func HealthCheckRoute(w http.ResponseWriter, r *http.Request) {
	resp := HealthStatusResponse{
		// control API status-code with http.StatusXXX
		StatusCode:    http.StatusOK,
		StatusMessage: "ok",
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(resp)
}
