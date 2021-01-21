package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/getsentry/sentry-go"
)

type Payload struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, payload Payload) (string, error) {
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("serverless function (lambda) AWS")

	return fmt.Sprintf("Program: %s!", payload.Name), nil
}

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:   "<your DSN>",
		Debug: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	lambda.Start(HandleRequest)
}
