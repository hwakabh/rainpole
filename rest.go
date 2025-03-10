package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
)

const l = "abcdefghijklmnopqrstuvwxyz" + "0123456789"

type RestResponse struct {
	Body string `json:"content"`
}

type UUID string

// Using receiver for implementing functions to user-defined types
func (u UUID) validate() error {
	if len(string(u)) != 36 {
		return fmt.Errorf("UUID should be 36 chars including its hyphen.")
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

func GenerateRandomUuid() UUID {
	fmt.Printf("Generating random UUID...\n")

	var first_part string
	for i := 0; i < 8; i++ {
		first_part += string(l[rand.Intn(len(l))])
	}
	var second_part string
	for i := 0; i < 4; i++ {
		second_part += string(l[rand.Intn(len(l))])
	}
	var third_part string
	for i := 0; i < 4; i++ {
		third_part += string(l[rand.Intn(len(l))])
	}
	var fourth_part string
	for i := 0; i < 4; i++ {
		fourth_part += string(l[rand.Intn(len(l))])
	}
	var last_part string
	for i := 0; i < 12; i++ {
		last_part += string(l[rand.Intn(len(l))])
	}

	gen_uuid := first_part + "-" + second_part + "-" + third_part + "-" + fourth_part + "-" + last_part

	return UUID(gen_uuid)
}

func RestRoute(w http.ResponseWriter, r *http.Request) {
	resp := RestResponse{
		Body: "API root!",
	}
	GetTableRecords()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(resp)
}

// Generate random UUID
func GetRandomUuid(w http.ResponseWriter, r *http.Request) {
	var random_uuid UUID = GenerateRandomUuid()
	fmt.Printf("Result: [ %s ]\n", random_uuid)
	if err := random_uuid.validate(); err != nil {
		fmt.Println(err)
	}

	resp := RestResponse{
		Body: string(random_uuid),
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(resp)
}

// func GetCompany(w http.ResponseWriter, r *http.Request) {

// }
