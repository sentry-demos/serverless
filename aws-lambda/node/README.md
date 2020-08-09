
# Steps to test Sentry SDK integration with Node Lambda function:

 - Create development package from the source code
 - Upload it into Lambda console
 - Refer usage for use case specific handling 

## Create a development package

1. Clone the git repo on your development machine.

2. Cloned repo contains multiple folders, each of these folder contains an example of how we can capture errors/exceptions on the Sentry dashboard.

3. For any of the examples (folders), edit below section of the js file: replace 'dsn' with your own DSN
```
Sentry.init({
  dsn: "<your DSN>",
}); 
```
4. Install the dependencies: 
```html
npm install
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
* Enter Function Name, perhaps a camelcase version of the folder you were in or the zip file and prefix with the runtime viz. NodeInitializationError or NodeHandledException etc.
* Choose Runtime as **Node.js 12x**
* Select Use an existing role:lamda-role (assuming you created one following the instructions from the Top Level folder's README)   
* Hit Create Function
  * For more information refer to [Node Lambda function](https://docs.aws.amazon.com/lambda/latest/dg/lambda-nodejs.htm).


## Usage

Common to any/all examples below, you'd need to __Configure test events__ for each Lambda function. 
 * Goto the Lambda function on your AWS Management console (steps above would have already left you there).
 * Click *Test* on the far right.
 * Enter Event name such as MyTestEvent.
 * Leave the json as is or make it an empty json, your choice :-)
 * Click Create.
 * Follow specifics of the examples below.
 * __Click Test to send your example event to Sentry !!__
 
 ### 1. initialization_error:

This function contains code that instruments lambda initialization error. This is instrumented by calling undefined function before invoking the Lambda handler.

Edit *Basic Settings* on the Lambda dashbaord:
```html
a) Handler : initialization_error.handler 
b) Timeout : 30 sec (perhaps anything more than 5 seconds should do)
```

#### 2. handled_exception
This function contains code that instruments handled the exception.

Edit *Basic Settings* on the Lambda dashbaord:
```html
a) Handler : handled_exception.handler
```

#### 3. network_error_wrong_address
This function contains code that instruments the network error with an incorrect ip address.

Edit *Basic settings* on the Lambda Dashboard:
```html
a) Handler : network_error_wrong_address.handler
b) Timeout : 2 min 30 sec
```

#### 4. network_error_wrong_port
This function contains code that instruments the network error with an incorrect port.

Edit *Basic settings* on the Lambda Dashboard:
```html
a) Handler : network_error_wrong_port.handler 
b) Timeout : 2 min 30 sec
```

#### 5. custom_tag_and_context:
This function contains code that instruments the handled exception and creates custom tag & context. 

Edit *Basic settings* on the Lambda Dashboard:
```html
a) Handler : custom_tag_and_context.handler
```

Edit  *Environment variables* 
```
Click _Add environment variable_
   Key : ENV_VAR
   Value : env_variable_value
Hit Save
```

## Side Note
The AWS Lambda console has a sort of chicken and egg problem with calling out that the __handler__ name does not match with the name of the file you've uploaded. So it keeps "warning" you under the Function Code section with something like: "Lambda can't find the file initialization_error.js. Make sure that your handler upholds the format: file-name.method." 

This happens because I have edited the Basic settings (correctly) with the handler name as out_of_memory. But I still havent uploaded the zip file that contains initialization_error.js 
If I upload the zip file first, it'd still flag it saying my hanlder name is not matching the file I have uploaded :-) 