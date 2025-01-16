package kalshigo

import (
	"testing"
	"time"

	"github.com/arvindh-manian/kalshigo/structs"
	"github.com/google/uuid"
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

func TestCreateOrder(t *testing.T) {
	params := structs.CreateOrderParams{
		Action:              structs.OrderActionBuy,
		MarketTicker:        "EVSHARE-30JAN-50",
		Count:               1,
		ClientOrderID:       uuid.New().String(),
		PostOnly:            true,
		Side:                structs.OrderSideYes,
		Type:                structs.OrderTypeLimit,
		YesPrice:            1,
		ExpirationTimestamp: &structs.Timestamp{time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)},
	}

	resp, err := kg.CreateOrder(&params)
	if err != nil {
		t.Errorf("Error creating order: %v", err)
	}

	if resp.Order.Status != structs.OrderStatusCanceled {
		t.Errorf("Order status is not canceled: %v", resp.Order.Status)
	}
}
