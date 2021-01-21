package main

import (
	"context"
	"errors"
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
	fmt.Println("handled_exception")

	sentry.CaptureException(errors.New("handled, capture exception"))

	return fmt.Sprintf("Program: %s!", payload.Name), nil
}

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:   "https://879a3ddfdd5241b0b4f6fcf9011896ad@o87286.ingest.sentry.io/5426957",
		Debug: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	lambda.Start(HandleRequest)
}
