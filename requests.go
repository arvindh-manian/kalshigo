package kalshigo

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func (c *Client) makeRequest(method string, path string, payload interface{}) (*http.Response, error) {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	timestampStr := strconv.FormatInt(timestamp, 10)

	// get url without query params
	parsedUrl, err := url.Parse(path)

	if err != nil {
		return nil, err
	}

	noParamsPath := parsedUrl.Path

	msgString := timestampStr + method + noParamsPath

	signature, err := signPSS(c.PrivateKey, msgString)
	if err != nil {
		return nil, err
	}

	var req *http.Request

	if payload == nil {
		req, err = http.NewRequest(method, c.BaseURL+path, nil)
	} else {
		var payloadBytes []byte

		payloadBytes, err = json.Marshal(payload)

		if err != nil {
			return nil, err
		}

		req, err = http.NewRequest(method, c.BaseURL+path, bytes.NewBuffer(payloadBytes))
	}

	if err != nil {
		return nil, err
	}

	req.Header.Set("KALSHI-ACCESS-KEY", c.AccessKey)
	req.Header.Set("KALSHI-ACCESS-SIGNATURE", signature)
	req.Header.Set("KALSHI-ACCESS-TIMESTAMP", timestampStr)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	return client.Do(req)
}
