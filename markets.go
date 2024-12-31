package kalshigo

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
)

const MARKET_PATH = "/trade-api/v2/markets"

const SERIES_PATH = "/trade-api/v2/series"

const EVENT_PATH = "/trade-api/v2/events"

func (c *Client) GetSeries(params *GetSeriesParams) (Series, error) {
	parsedUrl, err := url.Parse(SERIES_PATH)
	if err != nil {
		return Series{}, err
	}

	parsedUrl = parsedUrl.JoinPath(strings.ToUpper(params.SeriesTicker))

	body, _, err := c.getRequest(parsedUrl.String(), nil)

	if err != nil {
		return Series{}, err
	}

	var returnSeries GetSeriesResponse

	err = json.Unmarshal(body, &returnSeries)

	if err != nil {
		return Series{}, err
	}

	return returnSeries.Series, nil
}

func (c *Client) GetMarket(params *GetMarketParams) (Market, error) {
	parsedUrl, err := url.Parse(MARKET_PATH)
	if err != nil {
		return Market{}, err
	}

	parsedUrl = parsedUrl.JoinPath(strings.ToUpper(params.MarketTicker))

	body, _, err := c.getRequest(parsedUrl.String(), nil)

	if err != nil {
		return Market{}, err
	}

	var returnMarket GetMarketResponse

	err = json.Unmarshal(body, &returnMarket)

	if err != nil {
		return Market{}, err
	}

	return returnMarket.Market, nil
}

func (c *Client) GetMarkets(params *GetMarketsParams) (GetMarketsResponse, error) {
	q := url.Values{}

	if params != nil {
		// TODO: Consider using some kind of reflection to autogenerate the map from the JSON encoding rules
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

	body, _, err := c.getRequest(MARKET_PATH, q)

	if err != nil {
		return GetMarketsResponse{}, err
	}

	var returnMarkets GetMarketsResponse

	err = json.Unmarshal(body, &returnMarkets)

	if err != nil {
		return GetMarketsResponse{}, err
	}

	return returnMarkets, nil
}
func (c *Client) GetEvent(params *GetEventParams) (Event, error) {

	parsedUrl, err := url.Parse(EVENT_PATH)

	if err != nil {
		return Event{}, err
	}

	var q = url.Values{}

	if params != nil {
		if params.WithNestedMarkets {
			q.Set("with_nested_markets", "true")
		}
	}

	parsedUrl = parsedUrl.JoinPath(strings.ToUpper(params.EventTicker))

	body, _, err := c.getRequest(parsedUrl.String(), q)

	if err != nil {
		return Event{}, err
	}

	var returnEvent GetEventResponse

	err = json.Unmarshal(body, &returnEvent)

	if err != nil {
		return Event{}, err
	}

	return returnEvent.Event, nil
}

func (c *Client) GetEvents(params *GetEventsParams) (GetEventsResponse, error) {
	q := url.Values{}

	if params != nil {
		if params.Limit != 0 {
			q.Set("limit", strconv.FormatInt(params.Limit, 10))
		}

		if params.Cursor != "" {
			q.Set("cursor", params.Cursor)
		}

		if params.SeriesTicker != "" {
			q.Set("series_ticker", params.SeriesTicker)
		}

		if params.Status != "" {
			q.Set("status", string(params.Status))
		}

		if params.WithNestedMarkets {
			q.Set("with_nested_markets", "true")
		}
	}

	body, _, err := c.getRequest(EVENT_PATH, q)

	if err != nil {
		return GetEventsResponse{}, err
	}

	var returnEvents GetEventsResponse

	err = json.Unmarshal(body, &returnEvents)

	if err != nil {
		return GetEventsResponse{}, err
	}

	return returnEvents, nil
}
