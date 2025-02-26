package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type RestResponse struct {
	Body string `json:"content"`
}

type UUID string

// Using receiver for implementing functions to user-defined types
func (u UUID) validate() error {
	if len(string(u)) != 36 {
		return fmt.Errorf("UUID should be 34 chars including its hyphen.")
	}

	elm := strings.Split(string(u), "-")
	// validation logics
	if len(elm) != 5 {
		return fmt.Errorf("UUID should consist 8-4-4-4-12 pattern")
	}
	if len(elm[0]) != 8 {
		return fmt.Errorf("First part of UUID should be 8 char")
	}
	if len(elm[1]) != 4 {
		return fmt.Errorf("Second part of UUID should be 4 char")
	}
	if len(elm[2]) != 4 {
		return fmt.Errorf("Third part of UUID should be 4 char")
	}
	if len(elm[3]) != 4 {
		return fmt.Errorf("Fourth part of UUID should be 4 char")
	}
	if len(elm[4]) != 12 {
		return fmt.Errorf("Last part of UUID should be 12 char")
	}
	return nil
}

func RestRoute(w http.ResponseWriter, r *http.Request) {
	resp := RestResponse{
		Body: "API root!",
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(resp)
}

// Generate random UUID
func GetRandomUuid(w http.ResponseWriter, r *http.Request) {
	var random_uuid UUID = "114bbea4-059f-483d-a604-f2c5ef688e5a"
	fmt.Printf("Generated UUID: %s \n", random_uuid)

	if err := random_uuid.validate(); err != nil {
		fmt.Println(err)
	}
}

// func GetCompany(w http.ResponseWriter, r *http.Request) {

// }
