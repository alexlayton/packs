package handler

import (
	"encoding/json"

	"github.com/alexlayton/packs/internal/packs"
	"github.com/aws/aws-lambda-go/events"
)

type calculateRequest struct {
	Count     int   `json:"count"`
	PackSizes []int `json:"pack_sizes"`
}

type calculateResponse struct {
	Packs packs.Packs `json:"packs"`
}

func CalculateHandler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	resp := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	params := calculateRequest{}
	err := json.Unmarshal([]byte(req.Body), &params)
	if err != nil {
		resp.StatusCode = 500
		return resp, nil
	}

	calculated := packs.Calculate(params.Count, params.PackSizes)

	calculateResp := calculateResponse{calculated}
	body, err := json.Marshal(calculateResp)
	if err != nil {
		resp.StatusCode = 500
	} else {
		resp.Body = string(body)
	}

	return resp, nil
}
