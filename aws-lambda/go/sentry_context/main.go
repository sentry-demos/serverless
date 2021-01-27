package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/getsentry/sentry-go"
)

// type Payload struct {
// 	Name string `json:"name"`
// }

func HandleRequest(ctx context.Context, webhook Webhook) (string, error) {
	defer sentry.Flush(2 * time.Second)
	// 1. add payload body of unknown size+ctonent as extra
	// 2. send the FirehoseJSON object, see what happens

	prettyPrint(webhook)

	// sentry.ConfigureScope(func(scope *sentry.Scope) {
	// 	scope.SetExtra("character.name", "Mighty Fighter")
	// })

	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetContext("Webhook", webhook)
	})

	sentry.CaptureMessage("sentry_context Message")

	return fmt.Sprintf("Webhook: %s", webhook.Id), nil
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

type Webhook struct {
	Type                     string                   `json:"type"`
	Id                       string                   `json:"id"`
	Timestamp_ms             float64                  `json:"timestamp_ms"`
	Firehose_version         string                   `json:"firehose_version"`
	Account                  map[string]interface{}   `json:"account"`
	User_identities          []map[string]interface{} `json:"user_identities"`
	Runtime_environment      map[string]interface{}   `json:"runtime_environment"`
	Events                   []map[string]interface{} `json:"events"`
	Source_channel           string                   `json:"source_channel"`
	Device_application_stamp string                   `json:"device_application_stamp"`
	Consent_state            map[string]interface{}   `json:"consent_state"`
	Mpid                     string                   `json:"mpid"`
}

func prettyPrint(v interface{}) {
	pp, _ := json.MarshalIndent(v, "", "  ")
	fmt.Print(string(pp))
}
