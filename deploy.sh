#! /bin/bash
GOOS=linux GOARCH=amd64 go build -o bin/calculate cmd/lambda/main.go
serverless deploy