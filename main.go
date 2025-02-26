package main

import (
	"fmt"
	"net/http"
	"time"
)

type RootResponse struct {
	Body string `json:"content"`
}

func main() {
	fmt.Println("Starting web server ...")
	// Instantiate multiplexer
	mux := http.NewServeMux()

	// --- Register mappings between URLs and handler
	// load staticfiles
	fileServer := http.FileServer(http.Dir("./web"))
	mux.Handle("/", http.StripPrefix("/", fileServer))

	// health check
	mux.HandleFunc("/health", HealthCheckRoute)

	// REST-APIs endpoints
	mux.HandleFunc("/api/v1/", RestRoute)
	mux.HandleFunc("/api/v1/uuid", GetRandomUuid)
	// mux.HandleFunc("/api/v1/companies", GetCompany)

	// GraphQL endpoints
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
