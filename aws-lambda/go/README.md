# Steps to test Sentry SDK integration with Go Lambda function:

[Demo Video](https://www.loom.com/share/e3625614df2f44f3bb96286ad7212f99)

## Create a development package
1. Clone the git repo on your development machine.

2. Build your executable by running

```
GOOS=linux go build main.go
```

3. Create a zip
```
zip function.zip main

// Include source code in zip to view source code in stacktrace.
// Or, setup a Sentry<>Github integration docs.sentry.io/product/integrations/github/
zip -r function.zip main main.go
```

## Upload Zip file to Lambda function on AWS.

1. Create a Go lambda function in AWS Lambda.
2. Function code > Actions > Upload zip
3. Click 'Test', the request payload is accessible as the second argument in the handler function.

## Usage
Common to any/all examples below, you'd need to __Configure test events__ for each Lambda function. 
 * Goto the Lambda function on your AWS Management console
 * Click *Test* on the far right.
 * Enter Event name such as MyTestEvent.
 * Leave the json as is or make it an empty json, your choice :-)
 * Click Create.
 * Follow specifics of the examples below.
 * __Click Test to send your example event to Sentry !!__

## Troubleshooting
Make sure your lambda Runtime Setting's handler name matches the name of your executable (i.e. main) or your function won't be found when you run the Test.

Setting GOOS to linux ensures that the compiled executable is compatible with the Go runtime, even if you compile it in a non-Linux environment.

## Additional Documentation
https://github.com/getsentry/examples/tree/master/aws-lambda
https://github.com/getsentry/examples/tree/master/gcp-cloud-functions

https://docs.sentry.io/platforms/go/serverless/

https://github.com/getsentry/sentry-go

https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html
