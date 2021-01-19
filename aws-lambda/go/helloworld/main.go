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
	fmt.Println("this appears in the lambda log.....")
	sentry.CaptureMessage("serverless function (lambda) AWS")
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
	// TODO try putting it after...

	// EVAL prob put at end of 'main'
	defer sentry.Flush(2 * time.Second)
}
