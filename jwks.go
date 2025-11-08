package main

import (
	"encoding/json"
	"net/http"
)

type JwksEndpoint struct {
	Keys []JWKS `json:"keys"`
}

type JWKS struct {
	SingArgo    string `json:"kty"`
	UseFor      string `json:"use"`
	RsaModulus  string `json:"n"`
	RsaExponent string `json:"e"`
	KeyId       string `json:"kid"`
	KeyArgo     string `json:"alg"`
}

func GetJsonWebKeySet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		json.NewEncoder(w).Encode(nil)
	}

	modulus, exponent := GenerateKeyPair()
	keys := []JWKS{
		{
			SingArgo:    "RSA",
			UseFor:      "sig",
			RsaModulus:  modulus,
			RsaExponent: exponent,
			KeyId:       "1",
			KeyArgo:     "RS256",
		},
	}
	json.NewEncoder(w).Encode((JwksEndpoint{
		Keys: keys,
	}))
}
