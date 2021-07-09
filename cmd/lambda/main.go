package main

import (
	"github.com/alexlayton/packs/internal/handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.CalculateHandler)
}
