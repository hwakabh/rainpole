package rainpole

import (
	"encoding/json"
	"net/http"
)

// TODO: implement with gqlgen packages
func GraphqlRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// TODO: replace with using type structs
	json.NewEncoder(w).Encode(RestResponse{
		Body: "GraphQL API root!",
	})
}
