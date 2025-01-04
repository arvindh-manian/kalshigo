package kalshigo

import (
	"testing"

	"github.com/arvindh-manian/kalshigo/structs"
)

func TestGetBalance(t *testing.T) {
	resp, err := kg.GetBalance()
	if err != nil {
		t.Errorf("Error getting balance: %v", err)
	}

	if resp < 0 {
		t.Errorf("Balance is negative: %v", resp)
	}
}

func TestGetFills(t *testing.T) {
	params := structs.GetFillsParams{
		Limit: 10,
	}

	resp, err := kg.GetFills(&params)
	if err != nil {
		t.Errorf("Error getting fills: %v", err)
	}

	if len(resp.Fills) == 0 {
		t.Errorf("No fills returned. This may be intended: this behavior expected with a fresh account.")
	}
}

func TestGetOrders(t *testing.T) {
	params := structs.GetOrdersParams{
		Limit: 10,
	}

	resp, err := kg.GetOrders(&params)
	if err != nil {
		t.Errorf("Error getting orders: %v", err)
	}

	if len(resp.Orders) == 0 {
		t.Errorf("No orders returned. This may be intended: this behavior expected with a fresh account.")
	}
}
