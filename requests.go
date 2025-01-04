package kalshigo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/arvindh-manian/kalshigo/structs"
)

func (c *Client) getRequest(path string, query url.Values) (body []byte, statusCode int, err error) {
	fullUrl := c.BaseURL.JoinPath(path)

	if query != nil {
		fullUrl.RawQuery = query.Encode()
	}

	resp, err := c.makeRequest("GET", fullUrl, nil)

	if err != nil {
		return nil, resp.StatusCode, err
	}

	if resp.StatusCode != 200 {
		return nil, resp.StatusCode, &structs.APIError{
			StatusCode: resp.StatusCode,
			Body:       fmt.Sprintf("Error getting %s: %s", path, resp.Status),
		}
	}

	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)

	if err != nil {
		return nil, resp.StatusCode, err
	}

	return body, resp.StatusCode, nil
}

// this should only be used by internal functions
// the base URL is added in other functions
func (c *Client) makeRequest(method string, requestUrl *url.URL, payload interface{}) (*http.Response, error) {
	timestampString := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)

	msgString := timestampString + method + "/" + requestUrl.Path

	signature, err := signPSS(c.PrivateKey, msgString)

	if err != nil {
		return nil, err
	}

	var req *http.Request

	if payload != nil {
		var payloadBytes []byte
		payloadBytes, err = json.Marshal(payload)

		if err != nil {
			return nil, err
		}

		req, err = http.NewRequest(method, requestUrl.String(), bytes.NewBuffer(payloadBytes))
	} else {
		req, err = http.NewRequest(method, requestUrl.String(), nil)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Set("KALSHI-ACCESS-KEY", c.AccessKey)
	req.Header.Set("KALSHI-ACCESS-SIGNATURE", signature)
	req.Header.Set("KALSHI-ACCESS-TIMESTAMP", timestampString)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return c.httpClient.Do(req)

}
