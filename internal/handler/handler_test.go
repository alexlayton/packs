package handler

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler200Status(t *testing.T) {
	params := calculateRequest{
		Count:     251,
		PackSizes: []int{250, 500, 1000, 2000, 5000},
	}

	jsonBytes, err := json.Marshal(params)

	temp := string(jsonBytes)
	fmt.Println(temp)

	if err != nil {
		t.Error(err)
	}

	req := events.APIGatewayProxyRequest{
		Body: string(jsonBytes),
	}
	resp, err := CalculateHandler(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("expected status 200, actual - %d\n", resp.StatusCode)
	}

	calculateResp := calculateResponse{}
	err = json.Unmarshal([]byte(resp.Body), &calculateResp)
	if err != nil {
		t.Error(err)
	}
}
