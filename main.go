package rainpole

import (
	"encoding/json"
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
	mux.HandleFunc("/health", HealthCheckRoute)
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
