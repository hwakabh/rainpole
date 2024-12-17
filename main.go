package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Starting web server ...")
	// Instantiate multiplexer
	mux := http.NewServeMux()

	// Register mappings between URLs and handler
	mux.HandleFunc("/", DefaultRoute)
	mux.HandleFunc("/healthz", HealthCheckRoute)

	// Using struct in net/http with instantiated multiplexer
	server := http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()

}

// TODO: return as JSON with code 200
func DefaultRoute(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, r.Method, r.URL)
	// fmt.Fprintln(w, r.Proto)
	fmt.Fprintf(w, "hello, developers")
}

// TODO: return as JSON {"status": "ok"} with code 200
func HealthCheckRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}
