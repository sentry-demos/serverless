# Sentry for Serverless - AWS Lambda

Here you find examples of how to use Sentry SDKs to capture errors from code running on AWS Lambda.

- [Node](node)
- [Python](python)

# Prerequisites

## AWS Account
  * [AWS Signup](https://portal.aws.amazon.com/billing/signup#/start)
  * NOTE: The free tier should be able to cover these test runs (**YMMV**)

### Execution role for Lambda
 * Lambda functions use an execution role to get permission to write logs to Amazon CloudWatch Logs, and to access other services and resources. If you don't already have an execution role for function development, create one.

### To create an execution role
 * Open the roles page in the IAM console.
 * Choose Create role.
 * Create a role with the following properties:
   * Trusted entity – Lambda.
   * Permissions – AWSLambdaBasicExecutionRole.
   * Role name – lambda-role.
* The AWSLambdaBasicExecutionRole policy has the permissions that the function needs to write logs to CloudWatch Logs.
* [Reference](https://docs.aws.amazon.com/lambda/latest/dg/lambda-nodejs.html)

## (Optional) Install AWS CLI
 * Under node's and python's README, we'll cover the steps of uploading our Lambda code via the Management console. You'd probably want to be using the AWS CLI if you'd be doing this often.
   * [Installation link (MacOS)](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2-mac.html)
   * [Lambda specific setup](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-awscli.html)

