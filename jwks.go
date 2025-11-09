package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"math/big"
	"net/http"
	"os"
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

func LoadPublicKey() *rsa.PublicKey {
	if _, err := os.Stat("rsa.pub"); err == nil {
		f, _ := os.ReadFile("rsa.pub")
		pubKeyBlock, _ := pem.Decode(f)
		pubKey, err := x509.ParsePKCS1PublicKey(pubKeyBlock.Bytes)
		if err != nil {
			fmt.Println(err)
		}
		return pubKey
	}
	return nil
}

func GetJsonWebKeySet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		json.NewEncoder(w).Encode(nil)
	}

	fmt.Println(">>> Loading Public Key for JWKS responses ...")
	pub := LoadPublicKey()
	// modulus
	n := pub.N
	modulus := base64.RawURLEncoding.EncodeToString(n.Bytes())
	// exponent
	e := pub.E
	exponent := base64.RawURLEncoding.EncodeToString(big.NewInt(int64(e)).Bytes())

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
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode((JwksEndpoint{
		Keys: keys,
	}))
}
