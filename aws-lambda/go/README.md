# Steps to test Sentry SDK integration with Node Lambda function:
TODO

## Create a development package
1. Clone the git repo on your development machine.

2. Build your executable by running `GOOS=linux go build main.go`

TODO
```
// https://docs.sentry.io/platforms/go/serverless/
GOOS=linux go build -o bin/upload-image functions/upload-image/main.go && zip -r handler.zip bin/upload-image functions/upload-image/ helper/ util/
```
how about:  
```
GOOS=linux go build main.go && zip function.zip main
GOOS=linux go build main.go && zip -r function.zip main main.go

```
3. Make it into a zip by `zip function.zip main`. Make sure your Runtime Setting's handler name matches the name of your executable (i.e. main).


## Upload Zip file to Lambda function on AWS.
1. Create a Go lambda function in AWS Lambda
2. Function code > Actions > Upload zip
3. Click 'Test', the request payload is accessible as the second arg to the handler function:
```
func HandleRequest(ctx context.Context, payload Payload) (string, error) {
	return fmt.Sprintf("Hello: %s!", payload.Name), nil
}
```


## Additional Documentation
https://github.com/getsentry/examples/tree/master/aws-lambda
https://github.com/getsentry/examples/tree/master/gcp-cloud-functions

https://docs.sentry.io/platforms/go/serverless/

https://github.com/getsentry/sentry-go