package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"math/big"
)

func GenerateKeyPair() (string, string) {
	// Generate private key
	fmt.Println(">>> Generating private key ...")
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	//-> returns https://pkg.go.dev/crypto/rsa#PublicKey
	if err != nil {
		fmt.Println("Failed to generate Private Key")
	}
	// privateKeyPem := pem.EncodeToMemory(
	// 	&pem.Block{
	// 		Type:  "RSA PRIVATE KEY",
	// 		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	// 	},
	// )
	// if err := os.WriteFile("rsa.key", privateKeyPem, 0700); err != nil {
	// 	fmt.Println("Failed to export key file")
	// }

	// Generate public key
	fmt.Println(">>> Generating public key from private key ...")
	publicKey := privateKey.Public()
	//-> returns https://pkg.go.dev/crypto/rsa#PublicKey
	// publicKeyPem := pem.EncodeToMemory(
	// 	&pem.Block{
	// 		Type:  "RSA PUBLIC KEY",
	// 		Bytes: x509.MarshalPKCS1PublicKey(publicKey.(*rsa.PublicKey)),
	// 	},
	// )
	// if err := os.WriteFile("rsa.pub", publicKeyPem, 0700); err != nil {
	// 	fmt.Println("Failed to export key file")
	// }

	// // should be same
	// fmt.Printf("PrivateKey's Modulus: %s\n", privateKey.N)
	// fmt.Printf("PrivateKey's Exponent: %d\n", privateKey.E)
	// fmt.Printf("PublicKey's Modulus: %s\n", publicKey.(*rsa.PublicKey).N)
	// fmt.Printf("PublicKey's Exponent: %d\n", publicKey.(*rsa.PublicKey).E)
	n := publicKey.(*rsa.PublicKey).N
	e := publicKey.(*rsa.PublicKey).E

	return base64.URLEncoding.EncodeToString(n.Bytes()), base64.URLEncoding.EncodeToString(big.NewInt(int64(e)).Bytes())
}
