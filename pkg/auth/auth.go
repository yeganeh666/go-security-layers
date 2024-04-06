package auth

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
)

type Algorithm string

const (
	HMAC  Algorithm = "HMAC"
	ECDSA Algorithm = "ECDSA"
)

var (
	CryptographyAlgorithm = ECDSA
	encryptionKey         = "verySecretEncryptionKey0"

	privateKeyECDSA *ecdsa.PrivateKey
)

func init() {
	var err error
	privateKeyECDSA, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
}
