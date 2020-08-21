# Steps to test Sentry SDK integration with Python Lambda function:

 - Create development package from the source code
 - Upload it into Lambda console
 - Refer usage for use case specific handling 

## Create a development package

1. Clone the git repo on your development machine.

2. Cloned repo contains multiple folders, each of these folder contains an example of how we can capture errors/exceptions on the Sentry dashboard.

3. For any of the example (folders), edit below section of the .py file: replace 'dsn' with your own DSN
```python
sentry_sdk.init(
    dsn="<your DSN>",
    integrations=[AwsLambdaIntegration()]
)
```
4. Install the dependencies: 
This of course depends on what version of python is in your path or how you may have symlinked pip TO pip2|pip3 
```html
pip install -r requirements.txt -t .
OR
pip3 install -r requirements.txt -t .
```

5. Zip the folder contents so it can be uploaded to Lambda 
```html
$zip -r9 <folder_name>.zip .
```

## Upload Zip file to Lambda function on AWS.

* Login to [AWS console](https://aws.amazon.com/)
* Click *Lambda* under *Compute*
* Click *Create Function*
* (Author from scratch tile should already be selected)
* Enter Function Name, perhaps a CamelCase version of the folder you were in or the zip file and prefix with the runtime viz. PythonInvalidHandler or PythonHandledException etc.
* Choose Runtime as **Python 3.7*
* __NOTE__: The Latest available Python Runtime 3.8 in Lambda __DOES NOT__ work with Sentry SDK
* Select Use an existing role:**lamda-role** (assuming you created one following the instructions from the Top Level folder's README)   
* Hit Create Function
 * For more information refer to [Python Lambda function](https://docs.aws.amazon.com/lambda/latest/dg/lambda-python.htm).


## Usage

Common to any/all examples below, you'd need to __Configure test events__ for each Lambda function. 
 * Goto the Lambda function on your AWS Management console (steps above would have already left you there).
 * Click *Test* on the far right.
 * Enter Event name such as MyTestEvent.
 * Leave the json as is or make it an empty json, your choice :-)
 * Click Create.
 * Follow specifics of the examples below.
 * __Click Test to send your example event to Sentry !!__

#### 1. out_of_memory:

This AWS Lambda function contains code that consumes memory limit equal to set memory limit in the configuration

Edit *Basic Settings* on the Lambda dashbaord:
```html
a) Handler : out_of_memory.lambda_handler
b) Timeout : 1 min
```

#### 2. invalid_handler:
This Lambda function contains code that instruments Lambda invalid handler error.

Edit *Basic Settings* on the Lambda dashbaord:
```html
a) Handler : invalid_handler.lambda_handler_changed
```

#### 3. handled_exception:
This function contains code that instruments handled exception.

Edit *Basic Settings* on the Lambda dashbaord:
```html
a) Handler : handled_exception.lambda_handler
```

#### 4. unhandled_exception:
This function contains code that instruments unhandled exception.

Edit *Basic Settings* on the Lambda dashbaord:
```html
a) Handler : unhandled_exception.lambda_handler
```

#### 5. network_error_wrong_address:
This function contains code that instruments the network error with an incorrect IP address.

Edit *Basic Settings* on the Lambda dashbaord:
```html
a) Handler : network_error_wrong_address.handler 
b) Timeout : 2 min 30 sec
```

#### 6. network_error_wrong_port:
This function contains code that instruments the network error with an incorrect PORT number.

Edit *Basic Settings* on the Lambda dashbaord:
```html
a) Handler : network_error_wrong_port.handler 
b) Timeout : 2 min 30 sec
```

#### 7. Custom tags & Context:
This function contains code that instruments the handled exception and creates custom tag & context.

Edit *Basic Settings* on the Lambda dashbaord:
```html
a) Handler : custom_tag_and_context.lambda_handler
```
Edit  *Environment variables* 
```html
Click _Add environment variable_
   Key : ENV_VAR
   Value : Test Value
Hit Save   
```

## Side Note
The AWS Lambda console has a sort of chicken and egg problem with calling out that the __handler__ name does not match with the name of the file you've uploaded. So it keeps "warning" you under the Function Code section with something like: "Lambda can't find the file out_of_memory.py. Make sure that your handler upholds the format: file-name.method." 

This happens because I have edited the Basic settings (correctly) with the handler name as out_of_memory. But I still havent uploaded the zip file that contains out_of_memory.py. 
If I upload the zip file first, it'd still flag it saying my hanlder name is not matching the file I have uploaded :-) 