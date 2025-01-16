package kalshigo

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/arvindh-manian/kalshigo/structs"
)

const BALANCE_ENDPOINT = "/trade-api/v2/portfolio/balance"
const FILLS_ENDPOINT = "/trade-api/v2/portfolio/fills"
const ORDERS_ENDPOINT = "/trade-api/v2/portfolio/orders"

func (c *Client) GetBalance() (int, error) {
	body, _, err := c.getRequest(BALANCE_ENDPOINT, nil)

	if err != nil {
		return 0, err
	}

	var balanceResponse struct {
		Balance int `json:"balance"`
	}

	err = json.Unmarshal(body, &balanceResponse)

	if err != nil {
		return 0, err
	}

	return balanceResponse.Balance, nil
}

func (c *Client) GetFills(params *structs.GetFillsParams) (*structs.GetFillsResponse, error) {
	q := url.Values{}

	if params.Limit != 0 {
		q.Add("limit", strconv.Itoa(params.Limit))
	}

	if params.Cursor != "" {
		q.Add("cursor", params.Cursor)
	}

	if !params.MinTimestamp.IsZero() {
		q.Add("min_ts", strconv.FormatInt(params.MinTimestamp.Unix(), 10))
	}

	if !params.MaxTimestamp.IsZero() {
		q.Add("max_ts", strconv.FormatInt(params.MaxTimestamp.Unix(), 10))
	}

	if params.MarketTicker != "" {
		q.Add("ticker", params.MarketTicker)
	}

	if params.OrderID != "" {
		q.Add("order_id", params.OrderID)
	}

	body, _, err := c.getRequest(FILLS_ENDPOINT, q)

	if err != nil {
		return &structs.GetFillsResponse{}, err
	}

	var returnFills structs.GetFillsResponse

	err = json.Unmarshal(body, &returnFills)

	if err != nil {
		return &structs.GetFillsResponse{}, err
	}

	return &returnFills, nil
}

func (c *Client) GetOrders(params *structs.GetOrdersParams) (*structs.GetOrdersResponse, error) {
	q := url.Values{}

	if params.Limit != 0 {
		q.Add("limit", strconv.Itoa(params.Limit))
	}

	if params.Cursor != "" {
		q.Add("cursor", params.Cursor)
	}

	if !params.MinTimestamp.IsZero() {
		q.Add("min_ts", strconv.FormatInt(params.MinTimestamp.Unix(), 10))
	}

	if !params.MaxTimestamp.IsZero() {
		q.Add("max_ts", strconv.FormatInt(params.MaxTimestamp.Unix(), 10))
	}

	if params.MarketTicker != "" {
		q.Add("ticker", params.MarketTicker)
	}

	if params.EventTicker != "" {
		q.Add("event_ticker", params.EventTicker)
	}

	body, _, err := c.getRequest(ORDERS_ENDPOINT, q)

	if err != nil {
		return &structs.GetOrdersResponse{}, err
	}

	var returnOrders structs.GetOrdersResponse

	err = json.Unmarshal(body, &returnOrders)

	if err != nil {
		return &structs.GetOrdersResponse{}, err
	}

	return &returnOrders, nil
}

func (c *Client) CreateOrder(params *structs.CreateOrderParams) (*structs.CreateOrderResponse, error) {
	body, _, err := c.postRequest(ORDERS_ENDPOINT, params)

	if err != nil {
		return &structs.CreateOrderResponse{}, err
	}

	var returnOrder structs.CreateOrderResponse

	err = json.Unmarshal(body, &returnOrder)

	if err != nil {
		return &structs.CreateOrderResponse{}, err
	}

	return &returnOrder, nil
}
