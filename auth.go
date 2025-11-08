package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserPassCredentinals struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type HandlerResponse struct {
	IsSuccess    bool   `json:"success"`
	Result       string `json:"token,omitempty"`
	ErrorMessage string `json:"error,omitempty"`
}

type JWTHeader struct {
	Algorithm string `json:"alg"`
	MediaType string `json:"typ"`
}

type JWTPayload struct {
	Issuer   string `json:"iss"`
	Subject  string `json:"sub"`
	Expire   int    `json:"exp"`
	Username string `json:"username"`
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		json.NewEncoder(w).Encode(HandlerResponse{
			IsSuccess:    false,
			ErrorMessage: "Method Not Allowed, user POST for login",
		})
	} else {
		var rbody UserPassCredentinals
		// .Decode(&rbody) will not validate fields of UserPassCredentials
		if err := json.NewDecoder(r.Body).Decode(&rbody); err != nil {
			json.NewEncoder(w).Encode(HandlerResponse{
				IsSuccess:    false,
				ErrorMessage: "Failed to parse POST request",
			})
		} else {
			fmt.Println("POST request, got: ")
			fmt.Println(rbody)

			json.NewEncoder(w).Encode(VerifyUser(rbody))
		}

	}
}

func VerifyUser(c UserPassCredentinals) HandlerResponse {
	fmt.Println(">>> Checking credentials ...")
	if c.Username == "hwakabh" && c.Password == "changeme" {
		return HandlerResponse{
			IsSuccess: true,
			Result:    IssueJsonWebToken(c.Username),
		}
	} else {
		return HandlerResponse{
			IsSuccess:    false,
			ErrorMessage: "Invalid credentials",
		}
	}
}

func IssueJsonWebToken(username string) string {
	fmt.Println(">>> Generating JWT string ...")

	header, _ := json.Marshal(JWTHeader{
		Algorithm: "HS256",
		MediaType: "JWT",
	})
	h_b64 := base64.URLEncoding.EncodeToString(header)

	payload, _ := json.Marshal(JWTPayload{
		Issuer:   "rainpole.app",
		Subject:  "your.example.com",
		Expire:   1234,
		Username: username,
	})
	p_b64 := base64.URLEncoding.EncodeToString(payload)

	fmt.Printf("Header: %s \nPayload: %s\n", h_b64, p_b64)
	return fmt.Sprintf("%s.%s", h_b64, p_b64)
}
