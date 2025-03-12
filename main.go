package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type RootResponse struct {
	Body string `json:"content"`
}

func main() {
	// Fetch value from envar, and the zero-value of string is ""
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Failed to fetch envar with PORT, using default tcp/8080 ...")
		port = "8080"
	}

	fmt.Println("Starting web server ...")
	// Instantiate multiplexer
	mux := http.NewServeMux()

	// --- Register mappings between URLs and handler
	// load staticfiles
	fileServer := http.FileServer(http.Dir("./public"))
	mux.Handle("/", http.StripPrefix("/", fileServer))

	// health check
	mux.HandleFunc("/health", HealthCheckRoute)

	// REST-APIs endpoints
	mux.HandleFunc("/api/v1/", RestRoute)
	mux.HandleFunc("/api/v1/ip", GetSourceIpAddress)
	mux.HandleFunc("/api/v1/uuid", GetRandomUuid)
	mux.HandleFunc("/api/v1/companies", MultipleCompanyHandler)
	mux.HandleFunc("/api/v1/companies/{id}", SingleCompanyHandler)

	// GraphQL endpoints
	mux.HandleFunc("/graphql", GraphqlRoute)

	// Using struct in net/http with instantiated multiplexer
	server := http.Server{
		Addr:           ":" + port,
		Handler:        mux,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("listening server port [ " + server.Addr + " ] ...")
	server.ListenAndServe()
}
