package kalshigo

type APIError struct {
	StatusCode int
	Body       string
}

func (e *APIError) Error() string {
	return e.Body
}

type GetSeriesParams struct {
	// EventTicker is the ticker of the event. This is a required field.
	EventTicker string `json:"event_ticker"`

	// WithNestedMarkets is a boolean that determines whether to include the markets in the response. This is an optional field.
	WithNestedMarkets bool `json:"with_nested_markets,omitempty"`
}

type SeriesCategory string

const (
	SeriesCategoryPolitics   SeriesCategory = "Politics"
	SeriesCategoryEconomics  SeriesCategory = "Economics"
	SeriesCategoryCulture    SeriesCategory = "Culture"
	SeriesCategoryClimate    SeriesCategory = "Climate"
	SeriesCategoryFinancials SeriesCategory = "Financials"
	SeriesCategoryCrypto     SeriesCategory = "Crypto"
)

type SettlementSource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type SeriesResponse struct {
	Series Series `json:"series"`
}

type Series struct {
	Category          SeriesCategory     `json:"category"`
	ContractUrl       string             `json:"contract_url"`
	Frequency         string             `json:"frequency"`
	SettlementSources []SettlementSource `json:"settlement_sources"`
	Tags              []string           `json:"tags"`
	Ticker            string             `json:"ticker"`
	Title             string             `json:"title"`
}
