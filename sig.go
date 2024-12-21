package kalshigo

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"os"
)

var ErrPEMParse = errors.New("failed to parse PEM block containing the private key")

func loadPrivateKeyFromFile(filePath string) (*rsa.PrivateKey, error) {
	// Read the private key file
	privKeyBytes, err := os.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(privKeyBytes)
	if block == nil {
		return nil, ErrPEMParse
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func signPSS(privateKey *rsa.PrivateKey, message string) (string, error) {
	hash := sha256.New()
	hash.Write([]byte(message))
	hashedMessage := hash.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, hashedMessage, nil)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}
