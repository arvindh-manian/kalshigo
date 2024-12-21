package kalshigo

import (
	"os"
	"testing"
)

var (
	kg        *Client
	keyPath   = os.Getenv("TEST_KALSHI_KEY_PATH")
	accessKey = os.Getenv("TEST_KALSHI_ACCESS_KEY_ID")
	baseURL   = "https://api.elections.kalshi.com" // this should be changed to the demo URL once the demo is working again
)

func TestMain(m *testing.M) {
	if keyPath == "" || accessKey == "" {
		panic("TEST_KALSHI_KEY_PATH and TEST_KALSHI_ACCESS_KEY_ID must be set")
	}

	var err error
	kg, err = NewFromKeyPath(keyPath, accessKey, baseURL)

	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}
