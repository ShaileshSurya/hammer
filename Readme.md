# Hammer [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) [![Build Status](https://travis-ci.org/ShaileshSurya/hammer.svg?branch=master)](https://travis-ci.org/ShaileshSurya/hammer) [![Coverage Status](https://coveralls.io/repos/github/ShaileshSurya/hammer/badge.svg?branch=master)](https://coveralls.io/github/ShaileshSurya/hammer?branch=master)

Golang's Fluent HTTP Request Client

![alt text](https://github.com/ShaileshSurya/go-images/blob/master/go_pic.jpg?raw=true)

## Recipes

```go
client := hammer.New()
request, err := hammer.RequestBuilder().
    <HttpVerb>().
    WithURL("http://localhost:8081/employee").
    WithContext(context.Background()).
    WithHeaders("Accept", "application/json").
    WithHeaders("user-id", "10062").
    WithRequestParams("department", "HR").
    Build()

resp, err:= client.Execute(request)

// or

responseMap := make(map[string]interface{})
err:= client.ExecuteInto(request, &m)

// or

responseModel := Employee{}
err:= client.ExecuteInto(request, &model)

```

### Supported HTTP Verbs

```go
Get()
Head()
Post()
Put()
Patch()
Delete()
Connect()
Options()
Trace()
```

### Hammer Client Api's

```go

// New intializes and returns new Hammer Client
New()

// WithHTTPClient returns Hammer client with custom HTTPClient
WithHTTPClient(*http.Client)

// Execute the Request
Execute(*Request)

// Execute the Request and unmarshal into map or struct provided with unmarshalInto. Please See recipes.
ExecuteInto(*Request,unmarshalInto interface{})

```

### RequestBuilder Api's

```go
// WithRequestBody struct or map can be sent
WithRequestBody(body interface{})

// WithContext ...
WithContext(ctx context.Context)

// WithHeaders ...
WithHeaders(key string, value string)

// WithRequestParams ...
WithRequestParams(key string, value string)

// WithRequestBodyParams ...
WithRequestBodyParams(key string, value interface{})

// WithURL ...
WithURL(value string)

// WithBasicAuth ...
WithBasicAuth(username, password string)

// WithTemplate will create a request with already created request. See example below.
WithTemplate(tempRequest *Request)

```

```go
client := hammer.New()
request, err := hammer.RequestBuilder().
    Get().
    WithURL("http://localhost:8081/employee").
    WithHeaders("Accept", "application/json").
    WithHeaders("user-id", "10062").
    WithRequestParams("department", "HR").
    Build()

employeeList := []Employee{}

err := client.ExecuteInto(request,&EmployeeList)
```

```go
client := hammer.New()
reqTemp, err := hammer.RequestBuilder().
    Get().
    WithURL("http://localhost:8081/employee").
    WithHeaders("Accept", "application/json").
    WithHeaders("user-id", "10062").
    WithRequestParams("department", "HR").
    Build()

request, err := hammer.RequestBuilder().
    WithTemplate(reqTemp).
    WithRequestParams("post","manager").
    WithRequestParams("centre","pune")
    Build()

employeeList := []Employee{}

err:= client.ExecuteInto(request,&EmployeeList)

```

## Contributing

1. Fork the repo and create your branch from master.
2. If you've added code that should be tested, add tests.
3. If you've changed APIs, update the documentation.
4. Ensure the test suite passes.
5. Make sure your code lints.
6. Issue that pull request!

View [CONTRIBUTING.md](CONTRIBUTING.md) to learn more about how to contribute.

## License

This project is open source and available under the [MIT License](LICENSE).
