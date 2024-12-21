package kalshigo

import (
	"testing"
)

func TestGetSeries(t *testing.T) {
	_, err := kg.GetSeries(&GetSeriesParams{
		EventTicker: "kxopenaidays",
	})

	if err != nil {
		t.Errorf("Error getting series: %v", err)
	}
}
