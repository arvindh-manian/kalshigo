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
	m, err := kg.GetMarket("KXSBADS-25-B")

	if err != nil {
		t.Errorf("Error getting market: %v", err)
	}

	if m.Ticker != "KXSBADS-25-B" {
		t.Errorf("Expected ticker to be KXSBADS-25-B, got %v", m.Ticker)
	}
}
