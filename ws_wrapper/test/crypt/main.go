package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func main() {
	message := []byte("hello world")
	h := sha1.New()
	h.Write(message)
	hashed := h.Sum(nil)

	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		panic("failed to decode PEM block containing public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		panic("failed to get RSA public key")
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPub, crypto.SHA1, hashed)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Message: %s\n", message)
	fmt.Printf("Signature: %x\n", signature)
}

const publicKey = `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCYGOGFYlRkJ8azMJ6JuiovyQ8zjlrLjWbmXSsemlJ4ptQ02+GrFp4ozPFZdhs8tCxU0WO6M/DQFzeyxTI0I9rvH2QHSuCsvkZ3wN0SU6yK6jKtkbtCRE2TJQX7HFblsETJJ/X63MfDI8PqMXdkJuXDLlwDIuJWXaselZiFgQiCEwIDAQAB`
