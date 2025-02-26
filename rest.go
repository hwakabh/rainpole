package main

import (
	"encoding/json"
	"net/http"
)

type RestMessage struct {
	Body string `json:"content"`
}

func RestRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(RestMessage{
		Body: "API root!",
	})
}
