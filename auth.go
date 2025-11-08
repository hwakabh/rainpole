package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"
	"time"
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
	KeyId     string `json:"kid"`
	JwksUrl   string `json:"jku"`
}

type JWTPayload struct {
	Issuer    string `json:"iss"`
	Subject   string `json:"sub"`
	Audience  string `json:"aud"`
	IssueAt   int64  `json:"iat"`
	Expire    int64  `json:"exp"`
	NotBefore int64  `json:"nbf"`
	Username  string `json:"username"`
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
		// Get JWT (header + payload)
		jwt := IssueJsonWebToken(c.Username)

		return HandlerResponse{
			IsSuccess: true,
			// Sign to JWT
			Result: AppendSignature(jwt),
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
		Algorithm: "RS256",
		MediaType: "JWT",
		KeyId:     "1",
		JwksUrl:   "http://localhost:8080/api/v1/.well-known/jwks.json",
	})
	h_b64 := base64.RawURLEncoding.EncodeToString(header)

	ut := time.Now().Unix()
	payload, _ := json.Marshal(JWTPayload{
		Issuer:    "rainpole.app",
		Subject:   "your.example.com",
		IssueAt:   ut,
		NotBefore: ut,
		Expire:    ut + 600, // TTL is 10m with "exp" claim
		Audience:  "vault.example.com",
		Username:  username,
	})
	p_b64 := base64.RawURLEncoding.EncodeToString(payload)

	fmt.Printf("Header: %s \nPayload: %s\n", h_b64, p_b64)
	return fmt.Sprintf("%s.%s", h_b64, p_b64)
}

func LoadPrivateKey() *rsa.PrivateKey {
	if _, err := os.Stat("rsa.key"); err == nil {
		f, _ := os.ReadFile("rsa.key")
		privKeyBlock, _ := pem.Decode(f)
		privKey, err := x509.ParsePKCS1PrivateKey(privKeyBlock.Bytes)
		if err != nil {
			fmt.Println(err)
		}
		return privKey
	}
	fmt.Println("rsa.key file does not exist.")
	return nil
}

func AppendSignature(jwt string) string {
	fmt.Println(">>> Loading private key file")
	k := LoadPrivateKey()

	// sign to jwt with RS256
	hasher := sha256.New()
	hasher.Write([]byte(jwt))
	digest := hasher.Sum(nil)

	signature, err := k.Sign(rand.Reader, digest, crypto.SHA256)
	if err != nil {
		fmt.Println("Failed to sign")
		fmt.Println(err)
		return fmt.Sprintf("%s.%s", jwt, "Failed")
	}

	sig_b64 := base64.RawURLEncoding.EncodeToString(signature)

	fmt.Printf("Signature: %s\n", sig_b64)
	return fmt.Sprintf("%s.%s", jwt, sig_b64)
}
