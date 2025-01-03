package kalshigo

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	"github.com/arvindh-manian/kalshigo/structs"
)

const MARKET_PATH = "/trade-api/v2/markets"

const SERIES_PATH = "/trade-api/v2/series"

const EVENT_PATH = "/trade-api/v2/events"

const TRADE_PATH = "trade-api/v2/markets/trades"

func (c *Client) GetSeries(params *structs.GetSeriesParams) (*structs.Series, error) {
	parsedUrl, err := url.Parse(SERIES_PATH)
	if err != nil {
		return &structs.Series{}, err
	}

	parsedUrl = parsedUrl.JoinPath(strings.ToUpper(params.SeriesTicker))

	body, _, err := c.getRequest(parsedUrl.String(), nil)

	if err != nil {
		return &structs.Series{}, err
	}

	var returnSeries structs.GetSeriesResponse

	err = json.Unmarshal(body, &returnSeries)

	if err != nil {
		return &structs.Series{}, err
	}

	return &returnSeries.Series, nil
}

func (c *Client) GetMarket(params *structs.GetMarketParams) (*structs.Market, error) {
	parsedUrl, err := url.Parse(MARKET_PATH)
	if err != nil {
		return &structs.Market{}, err
	}

	parsedUrl = parsedUrl.JoinPath(strings.ToUpper(params.MarketTicker))

	body, _, err := c.getRequest(parsedUrl.String(), nil)

	if err != nil {
		return &structs.Market{}, err
	}

	var returnMarket structs.GetMarketResponse

	err = json.Unmarshal(body, &returnMarket)

	if err != nil {
		return &structs.Market{}, err
	}

	return &returnMarket.Market, nil
}

func (c *Client) GetMarkets(params *structs.GetMarketsParams) (*structs.GetMarketsResponse, error) {
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

		if !params.MaxCloseTimestamp.IsZero() {
			q.Set("max_close_ts", strconv.FormatInt(params.MaxCloseTimestamp.Unix(), 10))
		}

		if !params.MinCloseTimestamp.IsZero() {
			q.Set("min_close_ts", strconv.FormatInt(params.MinCloseTimestamp.Unix(), 10))
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
		return &structs.GetMarketsResponse{}, err
	}

	var returnMarkets structs.GetMarketsResponse

	err = json.Unmarshal(body, &returnMarkets)

	if err != nil {
		return &structs.GetMarketsResponse{}, err
	}

	return &returnMarkets, nil
}
func (c *Client) GetEvent(params *structs.GetEventParams) (*structs.Event, error) {

	parsedUrl, err := url.Parse(EVENT_PATH)

	if err != nil {
		return &structs.Event{}, err
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
		return &structs.Event{}, err
	}

	var returnEvent structs.GetEventResponse

	err = json.Unmarshal(body, &returnEvent)

	if err != nil {
		return &structs.Event{}, err
	}

	return &returnEvent.Event, nil
}

func (c *Client) GetEvents(params *structs.GetEventsParams) (*structs.GetEventsResponse, error) {
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
		return &structs.GetEventsResponse{}, err
	}

	var returnEvents structs.GetEventsResponse

	err = json.Unmarshal(body, &returnEvents)

	if err != nil {
		return &structs.GetEventsResponse{}, err
	}

	return &returnEvents, nil
}

func (c *Client) GetTrades(params *structs.GetTradesParams) (*structs.GetTradesResponse, error) {
	q := url.Values{}

	if params != nil {
		if params.Limit != 0 {
			q.Set("limit", strconv.FormatInt(int64(params.Limit), 10))
		}

		if params.Cursor != "" {
			q.Set("cursor", params.Cursor)
		}

		if params.MarketTicker != "" {
			q.Set("ticker", params.MarketTicker)
		}

		if !params.MinTimestamp.IsZero() {
			q.Set("min_ts", strconv.FormatInt(params.MinTimestamp.Unix(), 10))
		}

		if !params.MaxTimestamp.IsZero() {
			q.Set("max_ts", strconv.FormatInt(params.MaxTimestamp.Unix(), 10))
		}
	}

	body, _, err := c.getRequest(TRADE_PATH, q)

	if err != nil {
		return &structs.GetTradesResponse{}, err
	}

	var returnTrades structs.GetTradesResponse

	err = json.Unmarshal(body, &returnTrades)

	if err != nil {
		return &structs.GetTradesResponse{}, err
	}

	return &returnTrades, nil
}

func (c *Client) GetMarketOrderbook(params *structs.GetMarketOrderbookParams) (*structs.MarketOrderbook, error) {
	parsedUrl, err := url.Parse(MARKET_PATH)

	if err != nil {
		return &structs.MarketOrderbook{}, err
	}

	parsedUrl = parsedUrl.JoinPath(strings.ToUpper(params.MarketTicker), "orderbook")

	q := url.Values{}

	if params.Depth != 0 {
		q.Set("depth", strconv.FormatInt(int64(params.Depth), 10))
	}

	body, _, err := c.getRequest(parsedUrl.String(), q)

	if err != nil {
		return &structs.MarketOrderbook{}, err
	}

	var returnOrderbook structs.GetMarketOrderbookResponse

	err = json.Unmarshal(body, &returnOrderbook)

	if err != nil {
		return &structs.MarketOrderbook{}, err
	}

	return &returnOrderbook.Orderbook, nil
}

func (c *Client) GetMarketCandlesticks(params *structs.GetMarketCandlesticksParams) ([]structs.MarketCandlestick, error) {
	parsedUrl, err := url.Parse(SERIES_PATH)

	if err != nil {
		return nil, err
	}

	parsedUrl = parsedUrl.JoinPath(strings.ToUpper(params.SeriesTicker), "markets", strings.ToUpper(params.MarketTicker), "candlesticks")

	q := url.Values{}

	if !params.StartTimestamp.IsZero() {
		q.Set("start_ts", strconv.FormatInt(params.StartTimestamp.Unix(), 10))
	}

	if !params.EndTimestamp.IsZero() {
		q.Set("end_ts", strconv.FormatInt(params.EndTimestamp.Unix(), 10))
	}

	if params.PeriodInterval != 0 {
		q.Set("period_interval", strconv.FormatInt(int64(params.PeriodInterval), 10))
	}

	body, _, err := c.getRequest(parsedUrl.String(), q)

	if err != nil {
		return nil, err
	}

	var returnCandlesticks structs.GetMarketCandlesticksResponse

	err = json.Unmarshal(body, &returnCandlesticks)

	if err != nil {
		return nil, err
	}

	return returnCandlesticks.Candlesticks, nil
}
