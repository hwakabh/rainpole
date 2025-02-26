package main

import (
	"encoding/json"
	"net/http"
)

type RestResponse struct {
	Body string `json:"content"`
}

func RestRoute(w http.ResponseWriter, r *http.Request) {
	resp := RestResponse{
		Body: "API root!",
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(resp)
}
