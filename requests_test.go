package kalshigo

import (
	"bytes"
	"fmt"
	"testing"
)

func TestMakeRequest(t *testing.T) {
	method := "GET"
	path := "/trade-api/v2/portfolio/balance"

	response, err := kg.makeRequest(method, path, nil)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	body := new(bytes.Buffer)
	body.ReadFrom(response.Body)
	fmt.Println("Response Body:", body.String())
}
