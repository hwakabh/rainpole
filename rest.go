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

// HTTP client with types
type HttpbinGetSchema struct {
	Args      map[string]string `json:"args"`
	HeaderMap map[string]string `json:"headers"`
	OriginIp  string            `json:"origin"`
	UrlString string            `json:"url"`
}

type SourceIpResponseBody struct {
	Body   string `json:"content"`
	IpAddr string `json:"address"`
}

func GetSourceIpAddress(w http.ResponseWriter, r *http.Request) {
	url := "https://httpbin.org/get"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to invoke URL [ %s ]\n", url)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to get responses. (Code: %v)\n", resp.StatusCode)
	}

	raw := &HttpbinGetSchema{}

	d := json.NewDecoder(resp.Body)
	err = d.Decode(raw)
	if err != nil {
		fmt.Println("Failed to decode JSON")
	}
	fmt.Printf("The IP address of you is [ %s ]\n", raw.OriginIp)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	body := SourceIpResponseBody{
		Body:   "OK",
		IpAddr: raw.OriginIp,
	}
	json.NewEncoder(w).Encode(body)
	//-> results:
	// {
	// 	"content": "OK",
	// 	"address": "125.14.177.109"
	// }
}

// // HTTP Client without types
// func GetHttpBinResponse() {
// 	url := "https://httpbin.org/get"
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Printf("Failed to invoke URL [ %s ]\n", url)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("Failed to get responses. (Code: %v)\n", resp.StatusCode)
// 	}

// 	r, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Failed to parse response with JSON")
// 	}
// 	fmt.Println(r)
// 	fmt.Println(string(r))
// }
