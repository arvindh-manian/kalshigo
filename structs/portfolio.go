package structs

import "time"

type GetFillsParams struct {
	MarketTicker string    `json:"ticker"`
	OrderID      string    `json:"order_id"`
	MinTimestamp Timestamp `json:"min_ts"`
	MaxTimestamp Timestamp `json:"max_ts"`
	// limit should be 1-1000
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor"`
}

type OrderAction string

const (
	OrderActionBuy     OrderAction = "buy"
	OrderActionSell    OrderAction = "sell"
	OrderActionUnknown OrderAction = ""
)

type Fill struct {
	Action      OrderAction `json:"action"`
	Count       int         `json:"count"`
	CreatedTime time.Time   `json:"created_time"`
	IsTaker     bool        `json:"is_taker"`
	// No price in cents
	NoPrice int    `json:"no_price"`
	OrderID string `json:"order_id"`
	// Same choices as TakerSideType
	Side         TakerSideType `json:"side"`
	MarketTicker string        `json:"ticker"`
	TradeID      string        `json:"trade_id"`
	// Yes price in cents
	YesPrice int `json:"yes_price"`
}

type GetFillsResponse struct {
	Fills  []Fill `json:"fills"`
	Cursor string `json:"cursor"`
}

type OrderRequestStatus string

const (
	OrderRequestStatusResting  OrderRequestStatus = "resting"
	OrderRequestStatusCanceled OrderRequestStatus = "canceled"
	OrderRequestStatusExecuted OrderRequestStatus = "executed"
)

type OrderStatus string

const (
	OrderStatusResting  OrderStatus = "resting"
	OrderStatusCanceled OrderStatus = "canceled"
	OrderStatusExecuted OrderStatus = "executed"
	OrderStatusPending  OrderStatus = "pending"
)

type OrderType string

const (
	OrderTypeMarket  OrderType = "market"
	OrderTypeLimit   OrderType = "limit"
	OrderTypeUnknown OrderType = ""
)

type GetOrdersParams struct {
	MarketTicker string             `json:"ticker"`
	EventTicker  string             `json:"event_ticker"`
	MinTimestamp Timestamp          `json:"min_ts"`
	MaxTimestamp Timestamp          `json:"max_ts"`
	Status       OrderRequestStatus `json:"status"`
	Cursor       string             `json:"cursor"`
	Limit        int                `json:"limit"` // 1-1000
}

type GetOrdersResponse struct {
	Orders []OrderResponse `json:"orders"`
	Cursor string          `json:"cursor"`
}

type OrderResponse struct {
	Action                    OrderAction   `json:"action"`
	AmendCount                int           `json:"amend_count"`
	AmendTakerFillCount       int           `json:"amend_taker_fill_count"`
	ClientOrderID             string        `json:"client_order_id"`
	CloseCancelCount          int           `json:"close_cancel_count"`
	CreatedTime               time.Time     `json:"created_time"`
	DecreaseCount             int           `json:"decrease_count"`
	ExpirationTime            time.Time     `json:"expiration_time"`
	FccCancelCount            int           `json:"fcc_cancel_count"`
	LastUpdateTime            time.Time     `json:"last_update_time"`
	MakerFees                 int           `json:"maker_fees"`
	MakerFillCost             int           `json:"maker_fill_cost"`
	MakerFillCount            int           `json:"maker_fill_count"`
	NoPrice                   int           `json:"no_price"`
	OrderID                   string        `json:"order_id"`
	PlaceCount                int           `json:"place_count"`
	QueuePosition             int           `json:"queue_position"`
	RemainingCount            int           `json:"remaining_count"`
	Side                      TakerSideType `json:"side"`
	Status                    OrderStatus   `json:"status"`
	TakerFees                 int           `json:"taker_fees"`
	TakerFillCost             int           `json:"taker_fill_cost"`
	TakerFillCount            int           `json:"taker_fill_count"`
	TakerSelfTradeCancelCount int           `json:"taker_self_trade_cancel_count"`
	MarketTicker              string        `json:"ticker"`
	Type                      OrderType     `json:"type"`
	UserID                    string        `json:"user_id"`
	YesPrice                  int           `json:"yes_price"`
}
