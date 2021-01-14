# Steps to test Sentry SDK integration with Node Lambda function:
TODO

## Create a development package
1. Clone the git repo on your development machine.

2. Build your executable by running `GOOS=linux go build main.go`

3. Make it into a zip by `zip function.zip main`. Make sure your Runtime Setting's handler name matches the name of your executable (i.e. main).


## Upload Zip file to Lambda function on AWS.
1. Create a Go lambda function in AWS Lambda
2. Function code > Actions > Upload zip
3. Click 'Test', the request payload is accessible as the second arg to the handler function:
```
// MyEvent is a struct type representing the request payload
func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
    ...
}
```