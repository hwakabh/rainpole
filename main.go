package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Typings of response body from /healthz
type HealthStatusResponse struct {
	StatusCode    int    `json:"code"`
	StatusMessage string `json:"status"`
}

type RestMessage struct {
	Body string `json:"content"`
}

func main() {
	fmt.Println("Starting web server ...")
	// Instantiate multiplexer
	mux := http.NewServeMux()

	// Register mappings between URLs and handler
	mux.HandleFunc("/", DefaultRoute)
	mux.HandleFunc("/healthz", HealthCheckRoute)
	mux.HandleFunc("/api/v1/", RestRoute)
	mux.HandleFunc("/graphql", GraphqlRoute)

	// Using struct in net/http with instantiated multiplexer
	server := http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("listening server port [ " + server.Addr + " ] ...")
	server.ListenAndServe()

}

func DefaultRoute(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, r.Method, r.URL)
	// fmt.Fprintln(w, r.Proto)
	// fmt.Fprintf(w, "hello, developers")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(RestMessage{
		Body: "hello, developers",
	})
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

func RestRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(RestMessage{
		Body: "API root!",
	})
}

// TODO: implement with gqlgen packages
func GraphqlRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// TODO: replace with using type structs
	json.NewEncoder(w).Encode(RestMessage{
		Body: "GraphQL API root!",
	})
}
