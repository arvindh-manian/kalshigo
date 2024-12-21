package kalshigo

import (
	"crypto/rsa"
)

type Client struct {
	PrivateKey *rsa.PrivateKey
	AccessKey  string
	BaseURL    string
}

func New(privateKey *rsa.PrivateKey, accessKey string, baseURL string) *Client {
	return &Client{
		PrivateKey: privateKey,
		AccessKey:  accessKey,
		BaseURL:    baseURL,
	}
}

func NewFromKeyPath(keyPath string, accessKey string, baseURL string) (*Client, error) {
	privateKey, err := loadPrivateKeyFromFile(keyPath)
	if err != nil {
		return nil, err
	}
	return New(privateKey, accessKey, baseURL), nil
}
