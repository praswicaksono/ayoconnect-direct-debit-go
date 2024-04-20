package directdebit

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

var (
	ErrParsePEMBlock      = errors.New("failed to parse PEM block containing the private key")
	GenerateRSASignature  = generateRSASignature
	GenerateHmacSignature = generateHmacSignature
)

func generateRSASignature(timestamp string, privkey string, clientID string) (string, error) {
	privKeyBlock, _ := pem.Decode([]byte(privkey))
	if privKeyBlock == nil {
		return "", ErrParsePEMBlock
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(privKeyBlock.Bytes)
	if err != nil {
		return "", err
	}

	pkey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return "", err
	}

	hash := sha256.Sum256([]byte(clientID + "|" + timestamp))
	signature, err := rsa.SignPKCS1v15(rand.Reader, pkey, crypto.SHA256, hash[:])
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", signature), nil
}

func generateHmacSignature(httpMethod string, path string, accessToken string, jsonString string, timestamp string, clientSecret string) string {
	stringToSign := httpMethod + ":" + path + ":" + accessToken + ":" + fmt.Sprintf("%x", sha256.Sum256([]byte(jsonString))) + ":" + timestamp
	hmac := hmac.New(sha512.New, []byte(clientSecret))

	hmac.Write([]byte(stringToSign))
	return fmt.Sprintf("%x", hmac.Sum(nil))
}
