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
	MinTimestamp Timestamp          `json:"min_ts,omitempty"`
	MaxTimestamp Timestamp          `json:"max_ts,omitempty"`
	Status       OrderRequestStatus `json:"status"`
	Cursor       string             `json:"cursor"`
	Limit        int                `json:"limit"` // 1-1000
}

type GetOrdersResponse struct {
	Orders []OrderResponse `json:"orders"`
	Cursor string          `json:"cursor,omitempty"`
}

type OrderResponse struct {
	// Action                    OrderAction   `json:"action"`
	// AmendCount                int           `json:"amend_count,omitempty"`
	// AmendTakerFillCount       int           `json:"amend_taker_fill_count,omitempty"`
	// ClientOrderID             string        `json:"client_order_id"`
	CloseCancelCount          int           `json:"close_cancel_count,omitempty"`
	CreatedTime               OptionalTime  `json:"created_time,omitempty"`
	DecreaseCount             int           `json:"decrease_count,omitempty"`
	ExpirationTime            OptionalTime  `json:"expiration_time,omitempty"`
	FccCancelCount            int           `json:"fcc_cancel_count,omitempty"`
	LastUpdateTime            OptionalTime  `json:"last_update_time,omitempty"`
	MakerFees                 int           `json:"maker_fees,omitempty"`
	MakerFillCost             int           `json:"maker_fill_cost,omitempty"`
	MakerFillCount            int           `json:"maker_fill_count,omitempty"`
	NoPrice                   int           `json:"no_price"`
	OrderID                   string        `json:"order_id"`
	PlaceCount                int           `json:"place_count,omitempty"`
	QueuePosition             int           `json:"queue_position,omitempty"`
	RemainingCount            int           `json:"remaining_count,omitempty"`
	Side                      TakerSideType `json:"side"`
	Status                    OrderStatus   `json:"status"`
	TakerFees                 int           `json:"taker_fees,omitempty"`
	TakerFillCost             int           `json:"taker_fill_cost,omitempty"`
	TakerFillCount            int           `json:"taker_fill_count,omitempty"`
	TakerSelfTradeCancelCount int           `json:"taker_self_trade_cancel_count,omitempty"`
	MarketTicker              string        `json:"ticker"`
	Type                      OrderType     `json:"type"`
	UserID                    string        `json:"user_id,omitempty"`
	YesPrice                  int           `json:"yes_price"`
}

type OrderSide string

const (
	OrderSideYes   OrderSide = "yes"
	OrderSideNo    OrderSide = "no"
	OrderSideUnset OrderSide = "" // this should only be used in response
)

type CreateOrderParams struct {
	Action              OrderAction `json:"action"`
	BuyMaxCost          int         `json:"buy_max_cost,omitempty"`
	ClientOrderID       string      `json:"client_order_id"`
	Count               int         `json:"count"`
	ExpirationTimestamp *Timestamp  `json:"expiration_ts,omitempty"`
	NoPrice             int         `json:"no_price,omitempty"`
	PostOnly            bool        `json:"post_only,omitempty"`
	SellPositionFloor   int         `json:"sell_position_floor,omitempty"`
	Side                OrderSide   `json:"side"`
	MarketTicker        string      `json:"ticker"`
	Type                OrderType   `json:"type"`
	YesPrice            int         `json:"yes_price,omitempty"`
}

type CreateOrderResponse struct {
	Order struct {
		Action         OrderAction  `json:"action"`
		ClientOrderID  string       `json:"client_order_id"`
		CreatedTime    OptionalTime `json:"created_time,omitempty"`
		ExpirationTime OptionalTime `json:"expiration_time,omitempty"`
		NoPrice        int          `json:"no_price"`
		OrderID        string       `json:"order_id"`
		Side           OrderSide    `json:"side"`
		Status         OrderStatus  `json:"status"`
		MarketTicker   string       `json:"ticker"`
		Type           OrderType    `json:"type"`
		UserID         string       `json:"user_id,omitempty"`
		YesPrice       int          `json:"yes_price"`
	} `json:"order"`
}
