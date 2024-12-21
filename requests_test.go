package kalshigo

import (
	"fmt"
	"testing"
)

func Test_makeRequest(t *testing.T) {
	method := "GET"
	path := "/trade-api/v2/portfolio/balance"

	response, err := kg.makeRequest(method, path, nil)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", response.StatusCode)
	}
}
