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
	//-> returns https://pkg.go.dev/crypto/rsa#PublicKey
	if err != nil {
		fmt.Println("Failed to generate Private Key")
	}

	if _, err := os.Stat("rsa.key"); err != nil {
		privateKeyPem := pem.EncodeToMemory(
			&pem.Block{
				Type:  "RSA PRIVATE KEY",
				Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
			},
		)
		if err := os.WriteFile("rsa.key", privateKeyPem, 0700); err != nil {
			fmt.Println("Failed to export key file")
			os.Exit(1)
		}
	}

	// Generate public key
	fmt.Println(">>> Generating public key from private key ...")
	publicKey := privateKey.Public()
	//-> returns https://pkg.go.dev/crypto/rsa#PublicKey

	if _, err := os.Stat("rsa.pub"); err != nil {
		publicKeyPem := pem.EncodeToMemory(
			&pem.Block{
				Type:  "RSA PUBLIC KEY",
				Bytes: x509.MarshalPKCS1PublicKey(publicKey.(*rsa.PublicKey)),
			},
		)
		if err := os.WriteFile("rsa.pub", publicKeyPem, 0700); err != nil {
			fmt.Println("Failed to export key file")
			os.Exit(1)
		}
	}

	// return privateKey, publicKey

	// // should be same
	// fmt.Printf("PrivateKey's Modulus: %s\n", privateKey.N)
	// fmt.Printf("PrivateKey's Exponent: %d\n", privateKey.E)
	// fmt.Printf("PublicKey's Modulus: %s\n", publicKey.(*rsa.PublicKey).N)
	// fmt.Printf("PublicKey's Exponent: %d\n", publicKey.(*rsa.PublicKey).E)
}
