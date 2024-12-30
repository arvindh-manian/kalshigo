package kalshigo

import (
	"encoding/json"
	"io"
	"net/url"
	"strconv"
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

func (c *Client) GetMarket(params *GetMarketParams) (Market, error) {
	MARKET_PATH := "/trade-api/v2/markets"
	parsedUrl, err := url.Parse(MARKET_PATH)

	if err != nil {
		return Market{}, err
	}

	parsedUrl = parsedUrl.JoinPath(strings.ToUpper(params.MarketTicker))

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

func (c *Client) GetMarkets(params *GetMarketsParams) (MarketsResponse, error) {
	MARKETS_PATH := "/trade-api/v2/markets"
	parsedUrl, err := url.Parse(MARKETS_PATH)

	if err != nil {
		return MarketsResponse{}, err
	}

	// add params as query parameters

	q := parsedUrl.Query()

	if params != nil {
		if params.Limit != 0 {
			q.Set("limit", strconv.FormatInt(params.Limit, 10))
		}

		if params.Cursor != "" {
			q.Set("cursor", params.Cursor)
		}

		if params.EventTicker != "" {
			q.Set("event_ticker", params.EventTicker)
		}

		if params.SeriesTicker != "" {
			q.Set("series_ticker", params.SeriesTicker)
		}

		if params.MaxCloseTs != 0 {
			q.Set("max_close_ts", strconv.FormatInt(params.MaxCloseTs, 10))
		}

		if params.MinCloseTs != 0 {
			q.Set("min_close_ts", strconv.FormatInt(params.MinCloseTs, 10))
		}

		if params.Status != "" {
			q.Set("status", string(params.Status))
		}

		if len(params.MarketTickers) > 0 {
			q.Set("tickers", strings.Join(params.MarketTickers, ","))
		}
	}

	parsedUrl.RawQuery = q.Encode()

	resp, err := c.makeRequest("GET", parsedUrl.String(), nil)
	if err != nil {
		return MarketsResponse{}, &APIError{
			StatusCode: resp.StatusCode,
			Body:       err.Error(),
		}
	}

	defer resp.Body.Close()
	s, err := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return MarketsResponse{}, &APIError{
			StatusCode: resp.StatusCode,
			Body:       string(s),
		}
	}

	if err != nil {
		return MarketsResponse{}, err
	}

	var returnMarkets MarketsResponse
	err = json.Unmarshal(s, &returnMarkets)

	if err != nil {
		return MarketsResponse{}, err
	}

	return returnMarkets, nil
}
