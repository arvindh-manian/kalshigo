package kalshigo

import (
	"crypto/rsa"
	"net/http"
	"net/url"
)

type Client struct {
	PrivateKey *rsa.PrivateKey
	AccessKey  string
	BaseURL    url.URL
	httpClient *http.Client
}

func New(privateKey *rsa.PrivateKey, accessKey string, baseURL string) *Client {
	parsedUrl, err := url.Parse(baseURL)

	if err != nil {
		panic(err)
	}

	return &Client{
		PrivateKey: privateKey,
		AccessKey:  accessKey,
		BaseURL:    *parsedUrl,
		httpClient: &http.Client{},
	}
}

func NewFromKeyPath(keyPath string, accessKey string, baseURL string) (*Client, error) {
	privateKey, err := loadPrivateKeyFromFile(keyPath)
	if err != nil {
		return nil, err
	}
	return New(privateKey, accessKey, baseURL), nil
}
