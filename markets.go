package kalshigo

import (
	"encoding/json"
	"io"
	"net/url"
	"strings"
)

func (c *Client) GetSeries(params *GetSeriesParams) (Series, error) {
	originalPath := "/trade-api/v2/series"
	parsedUrl, err := url.Parse(originalPath)
	if err != nil {
		return Series{}, err
	}

	parsedUrl = parsedUrl.JoinPath(strings.ToUpper(params.EventTicker))

	resp, err := c.makeRequest("GET", parsedUrl.String(), nil)
	if err != nil {
		return Series{}, err
	}

	// print esp
	defer resp.Body.Close()
	s, err := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return Series{}, &APIError{
			StatusCode: resp.StatusCode,
			Body:       string(s),
		}
	}

	if err != nil {
		return Series{}, err
	}

	var returnSeries SeriesResponse
	err = json.Unmarshal(s, &returnSeries)
	if err != nil {
		return Series{}, err
	}

	return returnSeries.Series, nil
}
