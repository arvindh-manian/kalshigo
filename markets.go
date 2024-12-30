package kalshigo

import (
	"encoding/json"
	"io"
	"net/url"
	"strings"
)

func (c *Client) GetSeries(params *GetSeriesParams) (Series, error) {
	SERIES_PATH := "/trade-api/v2/series"
	parsedUrl, err := url.Parse(SERIES_PATH)
	if err != nil {
		return Series{}, err
	}

	parsedUrl = parsedUrl.JoinPath(strings.ToUpper(params.SeriesTicker))

	resp, err := c.makeRequest("GET", parsedUrl.String(), nil)
	if err != nil {
		return Series{}, err
	}

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

func (c *Client) GetMarket(ticker string) (Market, error) {
	MARKET_PATH := "/trade-api/v2/markets"
	parsedUrl, err := url.Parse(MARKET_PATH)

	if err != nil {
		return Market{}, err
	}

	parsedUrl = parsedUrl.JoinPath(strings.ToUpper(ticker))

	resp, err := c.makeRequest("GET", parsedUrl.String(), nil)
	if err != nil {
		return Market{}, err
	}

	defer resp.Body.Close()
	s, err := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return Market{}, &APIError{
			StatusCode: resp.StatusCode,
			Body:       string(s),
		}
	}

	if err != nil {
		return Market{}, err
	}

	var returnMarket MarketResponse
	err = json.Unmarshal(s, &returnMarket)
	if err != nil {
		return Market{}, err
	}

	return returnMarket.Market, nil
}
