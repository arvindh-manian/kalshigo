package kalshigo

import (
	"encoding/json"
	"time"
)

type APIError struct {
	StatusCode int
	Body       string
}

func (e *APIError) Error() string {
	return e.Body
}

type GetSeriesParams struct {
	// SeriesTicker is the ticker of the series. This is a required field.
	SeriesTicker string `json:"series_ticker"`
}

type GetMarketParams struct {
	// MarketTicker is the ticker of the market. This is a required field.
	MarketTicker string `json:"market_ticker"`
}

type GetMarketsParams struct {
	Limit         int64        `json:"limit,omitempty"` // This should be within the range of 1-1000
	Cursor        string       `json:"cursor,omitempty"`
	EventTicker   string       `json:"event_ticker,omitempty"`
	SeriesTicker  string       `json:"series_ticker,omitempty"`
	MaxCloseTs    int64        `json:"max_close_ts,omitempty"`
	MinCloseTs    int64        `json:"min_close_ts,omitempty"`
	Status        MarketStatus `json:"status,omitempty"`
	MarketTickers []string     `json:"tickers,omitempty"`
}

type GetMarketsResponse struct {
	Markets []Market `json:"markets"`
	Cursor  string   `json:"cursor,omitempty"`
}

type MarketStatus string

const (
	MarketStatusOpen     MarketStatus = "open"
	MarketStatusClosed   MarketStatus = "closed"
	MarketStatusSettled  MarketStatus = "settled"
	MarketStatusUnopened MarketStatus = "unopened"
)

type MarketType string

const (
	MarketTypeBinary MarketType = "Binary"
	MarketTypeScalar MarketType = "Scalar"
)

type ResponsePriceUnit string

const (
	ResponsePriceUnitCent      ResponsePriceUnit = "usd_cent"
	ResponsePriceUnitCenticent ResponsePriceUnit = "usd_centi_cent"
)

type MarketResult string

const (
	MarketResultNoResult     MarketResult = ""
	MarketResultYes          MarketResult = "yes"
	MarketResultNo           MarketResult = "no"
	MarketResultVoid         MarketResult = "void"
	RangedMarketResultAllNo  MarketResult = "all_no"
	RangedMarketResultAllYes MarketResult = "all_yes"
)

type MarketStrikeType string

const (
	MarketStrikeTypeUnknown        MarketStrikeType = "unknown"
	MarketStrikeTypeGreater        MarketStrikeType = "greater"
	MarketStrikeTypeLess           MarketStrikeType = "less"
	MarketStrikeTypeGreaterOrEqual MarketStrikeType = "greater_or_equal"
	MarketStrikeTypeLessOrEqual    MarketStrikeType = "less_or_equal"
	MarketStrikeTypeBetween        MarketStrikeType = "between"
	MarketStrikeTypeFunctional     MarketStrikeType = "functional"
	MarketStrikeTypeCustom         MarketStrikeType = "custom"
)

type Market struct {
	CanCloseEarly           bool              `json:"can_close_early"`
	CapStrike               float64           `json:"cap_strike,omitempty"`
	Category                SeriesCategory    `json:"category"`
	CloseTime               time.Time         `json:"close_time"` // ISO 8601
	CustomStrike            interface{}       `json:"custom_strike,omitempty"`
	EventTicker             string            `json:"event_ticker"`
	ExpectedExpirationTime  time.Time         `json:"expected_expiration_time,omitempty"` // ISO 8601
	ExpirationTime          time.Time         `json:"expiration_time"`                    // ISO 8601
	ExpirationValue         string            `json:"expiration_value"`
	FeeWaiverExpirationTime time.Time         `json:"fee_waiver_expiration_time,omitempty"` // ISO 8601
	FloorStrike             float64           `json:"floor_strike,omitempty"`
	FunctionalStrike        string            `json:"functional_strike,omitempty"` // TODO: Parse this
	LastPrice               int64             `json:"last_price"`
	LatestExpirationTime    time.Time         `json:"latest_expiration_time"` // ISO 8601
	Liquidity               int64             `json:"liquidity"`
	MarketType              MarketType        `json:"market_type"`
	NoAsk                   int64             `json:"no_ask"`
	NoBid                   int64             `json:"no_bid"`
	NoSubTitle              string            `json:"no_sub_title"`
	NotionalValue           int64             `json:"notional_value"`
	OpenInterest            int64             `json:"open_interest"`
	OpenTime                time.Time         `json:"open_time"` // ISO 8601
	PreviousPrice           int64             `json:"previous_price"`
	PreviousYesAsk          int64             `json:"previous_yes_ask"`
	PreviousYesBid          int64             `json:"previous_yes_bid"`
	ResponsePriceUnits      ResponsePriceUnit `json:"response_price_units"`
	Result                  MarketResult      `json:"result"`
	RiskLimitCents          int64             `json:"risk_limit_cents"`
	RulesPrimary            string            `json:"rules_primary"`
	RulesSecondary          string            `json:"rules_secondary"`
	SettlementTimerSeconds  int32             `json:"settlement_timer_seconds"`
	SettlementValue         int64             `json:"settlement_value,omitempty"`
	Status                  MarketStatus      `json:"status"`
	StrikeType              MarketStrikeType  `json:"strike_type,omitempty"`
	SubTitle                string            `json:"subtitle"` // Deprecated
	TickSize                int64             `json:"tick_size"`
	MarketTicker            string            `json:"ticker"`
	Title                   string            `json:"title"`
	Volume                  int64             `json:"volume"`
	Volume24h               int64             `json:"volume_24h"`
	YesAsk                  int64             `json:"yes_ask"`
	YesBid                  int64             `json:"yes_bid"`
	YesSubTitle             string            `json:"yes_sub_title"`
}

// JSON marshalling for Market
type GetMarketResponse struct {
	Market Market `json:"market"`
}

type SeriesCategory string

const (
	SeriesCategoryPolitics   SeriesCategory = "Politics"
	SeriesCategoryEconomics  SeriesCategory = "Economics"
	SeriesCategoryCulture    SeriesCategory = "Entertainment"
	SeriesCategoryClimate    SeriesCategory = "Climate and Weather"
	SeriesCategoryFinancials SeriesCategory = "Financials"
	SeriesCategoryCrypto     SeriesCategory = "Crypto"
)

type SettlementSource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type GetSeriesResponse struct {
	Series Series `json:"series"`
}

type Series struct {
	Category          SeriesCategory     `json:"category"`
	ContractUrl       string             `json:"contract_url"`
	Frequency         string             `json:"frequency"`
	SettlementSources []SettlementSource `json:"settlement_sources"`
	Tags              []string           `json:"tags"`
	SeriesTicker      string             `json:"ticker"`
	Title             string             `json:"title"`
}

type Event struct {
	// deprecated
	Category             string    `json:"category"`
	CollateralReturnType string    `json:"collateral_return_type"`
	EventTicker          string    `json:"event_ticker"`
	Markets              []Market  `json:"markets,omitempty"`
	MutuallyExclusive    bool      `json:"mutually_exclusive"`
	SeriesTicker         string    `json:"series_ticker"`
	StrikeDate           time.Time `json:"strike_date,omitempty"`
	StrikePeriod         string    `json:"strike_period,omitempty"`
	SubTitle             string    `json:"sub_title"`
	Title                string    `json:"title"`
}

type GetEventResponse struct {
	Event Event `json:"event"`
	// deprecated
	Markets []Market `json:"markets,omitempty"`
}

type GetEventParams struct {
	EventTicker       string `json:"event_ticker"`
	WithNestedMarkets bool   `json:"with_nested_markets"`
}

type GetEventsParams struct {
	// This should be within the range of 1-200
	Limit int64 `json:"limit,omitempty"`
	// omit to get the first page
	Cursor            string       `json:"cursor,omitempty"`
	Status            MarketStatus `json:"status,omitempty"`
	SeriesTicker      string       `json:"series_ticker,omitempty"`
	WithNestedMarkets bool         `json:"with_nested_markets"`
}

type GetEventsResponse struct {
	Cursor string  `json:"cursor,omitempty"`
	Events []Event `json:"events"`
}

type GetTradesParams struct {
	Cursor string `json:"cursor,omitempty"`
	// This should be within the range of 1-1000
	Limit        int32  `json:"limit,omitempty"`
	MarketTicker string `json:"ticker,omitempty"`
	MinTimestamp int64  `json:"min_ts,omitempty"`
	MaxTimestamp int64  `json:"max_ts,omitempty"`
}

type TakerSideType string

const (
	TakerSideYes   TakerSideType = "yes"
	TakerSideNo    TakerSideType = "no"
	TakerSideUnset TakerSideType = ""
)

type Trade struct {
	Count        int32         `json:"count"`
	CreatedTime  time.Time     `json:"created_time"`
	NoPrice      int64         `json:"no_price"`
	TakerSide    TakerSideType `json:"taker_side"`
	MarketTicker string        `json:"ticker"`
	TradeID      string        `json:"trade_id"`
	YesPrice     int64         `json:"yes_price"`
}

type GetTradesResponse struct {
	Cursor string  `json:"cursor,omitempty"`
	Trades []Trade `json:"trades"`
}

type GetMarketOrderbookParams struct {
	MarketTicker string `json:"ticker"`
	// maximum number of orderbook price levels you want to see for either side
	Depth int32 `json:"depth,omitempty"`
}

type GetMarketOrderbookResponse struct {
	Orderbook MarketOrderbook `json:"orderbook"`
}

type MarketOrderbook struct {
	No  PriceToFreq `json:"no"`
	Yes PriceToFreq `json:"yes"`
}

// cent to number of pending resting orders
type PriceToFreq map[int]int

func (p *PriceToFreq) UnmarshalJSON(b []byte) error {
	var arr = make([][]int, 0)
	err := json.Unmarshal(b, &arr)

	if err != nil {
		return err
	}

	*p = make(map[int]int)

	for _, pair := range arr {
		(*p)[pair[0]] = pair[1]
	}

	return nil
}

func (p PriceToFreq) MarshalJSON() ([]byte, error) {
	arr := make([][]int, 0)

	for k, v := range p {
		arr = append(arr, []int{k, v})
	}

	return json.Marshal(arr)
}
