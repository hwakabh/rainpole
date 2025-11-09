package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func GenerateKeyPair() {
	// Generate private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	//-> returns https://pkg.go.dev/crypto/rsa#PrivateKey
	if err != nil {
		fmt.Println("Failed to generate Private Key")
	}

	if _, err := os.Stat(PRIVATE_KEY_PATH); err != nil {
		privateKeyPem := pem.EncodeToMemory(
			&pem.Block{
				Type:  "RSA PRIVATE KEY",
				Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
			},
		)
		if err := os.WriteFile(PRIVATE_KEY_PATH, privateKeyPem, 0700); err != nil {
			fmt.Printf("Failed to export key file to %s\n", PRIVATE_KEY_PATH)
			os.Exit(1)
		}
	}

	// Generate public key
	fmt.Printf(">>> Generating public key from private key (%s) ...\n", PRIVATE_KEY_PATH)
	publicKey := privateKey.Public()
	//-> https://pkg.go.dev/crypto/rsa#PublicKey
	//-> https://pkg.go.dev/crypto#PublicKey

	if _, err := os.Stat(PUBLIC_KEY_PATH); err != nil {
		publicKeyPem := pem.EncodeToMemory(
			&pem.Block{
				Type:  "RSA PUBLIC KEY",
				Bytes: x509.MarshalPKCS1PublicKey(publicKey.(*rsa.PublicKey)),
			},
		)
		if err := os.WriteFile(PUBLIC_KEY_PATH, publicKeyPem, 0700); err != nil {
			fmt.Printf("Failed to export key file to %s\n", PUBLIC_KEY_PATH)
			os.Exit(1)
		}
	}
	// fmt.Printf("PrivateKey's Modulus: %s\n", privateKey.N)
	// fmt.Printf("PrivateKey's Exponent: %d\n", privateKey.E)
	// fmt.Printf("PublicKey's Modulus: %s\n", publicKey.(*rsa.PublicKey).N)
	// fmt.Printf("PublicKey's Exponent: %d\n", publicKey.(*rsa.PublicKey).E)
}
