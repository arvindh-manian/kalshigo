package kalshigo

import (
	"testing"
	"time"

	"github.com/arvindh-manian/kalshigo/structs"
)

func TestGetSeries(t *testing.T) {
	s, err := kg.GetSeries(&structs.GetSeriesParams{
		SeriesTicker: "KXPAYROLLS",
	})

	if err != nil {
		t.Errorf("Error getting series: %v", err)
	}

	if s.SeriesTicker != "KXPAYROLLS" {
		t.Errorf("Expected ticker to be KXPAYROLLS, got %v", s.SeriesTicker)
	}
}

func TestGetMarket(t *testing.T) {
	m, err := kg.GetMarket(&structs.GetMarketParams{
		MarketTicker: "KXSBADS-25-B",
	})

	if err != nil {
		t.Errorf("Error getting market: %v", err)
	}

	if m.MarketTicker != "KXSBADS-25-B" {
		t.Errorf("Expected ticker to be KXSBADS-25-B, got %v", m.MarketTicker)
	}
}

func TestGetMarkets(t *testing.T) {
	m, err := kg.GetMarkets(&structs.GetMarketsParams{
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
	e, err := kg.GetEvent(&structs.GetEventParams{
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

func TestGetEvents(t *testing.T) {
	e, err := kg.GetEvents(&structs.GetEventsParams{
		Limit: 10,
	})

	if err != nil {
		t.Errorf("Error getting events: %v", err)
	}

	if len(e.Events) != 10 {
		t.Errorf("Expected 10 events, got %v", len(e.Events))
	}

	if e.Cursor == "" {
		t.Errorf("Expected cursor to be non-empty, got %v", e.Cursor)
	}
}

func TestGetTrades(t *testing.T) {
	tt, err := kg.GetTrades(&structs.GetTradesParams{
		Limit: 10,
	})

	if err != nil {
		t.Errorf("Error getting trades: %v", err)
	}

	if len(tt.Trades) != 10 {
		t.Errorf("Expected 10 trades, got %v", len(tt.Trades))
	}

	if tt.Cursor == "" {
		t.Errorf("Expected cursor to be non-empty, got %v", tt.Cursor)
	}
}

func TestGetMarketOrderbook(t *testing.T) {
	_, err := kg.GetMarketOrderbook(&structs.GetMarketOrderbookParams{
		MarketTicker: "KXSBADS-25-B",
	})

	if err != nil {
		t.Errorf("Error getting market orderbook: %v", err)
	}
}

func TestGetMarketCandlesticks(t *testing.T) {
	startUnix := 1735187897
	endUnix := 1735619897

	c, err := kg.GetMarketCandlesticks(&structs.GetMarketCandlesticksParams{
		MarketTicker:   "KXUSTR-26DEC31-JG",
		SeriesTicker:   "KXUSTR",
		StartTimestamp: structs.Timestamp{time.Unix(int64(startUnix), 0)},
		EndTimestamp:   structs.Timestamp{time.Unix(int64(endUnix), 0)},
		PeriodInterval: structs.PeriodIntervalHour,
	})

	if err != nil {
		t.Errorf("Error getting market candlesticks: %v", err)
	}

	if len(c) == 0 {
		t.Errorf("Expected non-empty candlesticks, got %v", len(c))
	}
}
