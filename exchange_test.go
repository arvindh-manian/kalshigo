package kalshigo

import (
	"testing"
)

func TestGetExchangeAnnouncements(t *testing.T) {
	_, err := kg.GetExchangeAnnouncements()

	if err != nil {
		t.Errorf("Error getting exchange announcements: %v", err)
	}
}

func TestGetExchangeSchedule(t *testing.T) {
	_, err := kg.GetExchangeSchedule()

	if err != nil {
		t.Errorf("Error getting exchange schedule: %v", err)
	}
}

func TestGetExchangeStatus(t *testing.T) {
	_, err := kg.GetExchangeStatus()

	if err != nil {
		t.Errorf("Error getting exchange status: %v", err)
	}
}
