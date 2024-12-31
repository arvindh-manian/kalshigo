package kalshigo

import (
	"testing"
)

func TestGetSeries(t *testing.T) {
	s, err := kg.GetSeries(&GetSeriesParams{
		SeriesTicker: "KXPAYROLLS",
	})

	if err != nil {
		t.Errorf("Error getting series: %v", err)
	}

	if s.Ticker != "KXPAYROLLS" {
		t.Errorf("Expected ticker to be KXPAYROLLS, got %v", s.Ticker)
	}
}

func TestGetMarket(t *testing.T) {
	m, err := kg.GetMarket(&GetMarketParams{
		MarketTicker: "KXSBADS-25-B",
	})

	if err != nil {
		t.Errorf("Error getting market: %v", err)
	}

	if m.Ticker != "KXSBADS-25-B" {
		t.Errorf("Expected ticker to be KXSBADS-25-B, got %v", m.Ticker)
	}
}

func TestGetMarkets(t *testing.T) {
	m, err := kg.GetMarkets(&GetMarketsParams{
		Limit: 10,
	})

	if err != nil {
		t.Errorf("Error getting markets: %v", err)
	}

	if len(m.Markets) != 10 {
		t.Errorf("Expected 10 markets, got %v", len(m.Markets))
	}

	if m.Cursor == "" {
		t.Errorf("Expected cursor to be non-empty, got %v", m.Cursor)
	}
}

func TestGetEvent(t *testing.T) {
	e, err := kg.GetEvent(&GetEventParams{
		EventTicker:       "KXSPEAKER",
		WithNestedMarkets: true,
	})

	if err != nil {
		t.Errorf("Error getting event: %v", err)
	}

	if e.EventTicker != "KXSPEAKER" {
		t.Errorf("Expected ticker to be KXSPEAKER, got %v", e.EventTicker)
	}
}
