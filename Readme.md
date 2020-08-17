
![alt text](https://github.com/ShaileshSurya/go-images/blob/master/go_pic.jpg?raw=true)

# Requesto  [![Build Status](https://travis-ci.org/ShaileshSurya/requesto.svg?branch=master)](https://travis-ci.org/ShaileshSurya/requesto) [![Coverage Status](https://coveralls.io/repos/github/ShaileshSurya/requesto/badge.svg?branch=master)](https://coveralls.io/github/ShaileshSurya/requesto?branch=master)


Golang's Fluent HTTP Request Client 



## Recipes

```
client := requesto.New()
request, err := requesto.RequestBuilder().
		<HttpVerb>().
		WithURL("http://localhost:8081/employee").
		WithHeaders("Accept", "application/json").
		WithHeaders("user-id", "10062").
		WithRequestParams("department", "HR").
		Build()

resp, err:= client.Execute(request)

or 

map := make(map[string]interface{})
resp.err:= client.ExecuteInto(request, &map)

or 

model := Employee{}
resp.err:= client.ExecuteInto(request, &model)

```

### Supported HTTP Verbs
```
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
### Requesto Client Api's
```
// New intializes and returns new Requesto Client
New()

// WithHTTPClient returns Requesto client with custom HTTPClient
WithHTTPClient(*http.Client)
```

### RequestBuilder Api's
```
// WithRequestBody ....
WithRequestBody(body interface{}) 

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


```
client := requesto.New()
request, err := requesto.RequestBuilder().
		Get().
		WithURL("http://localhost:8081/employee").
		WithHeaders("Accept", "application/json").
		WithHeaders("user-id", "10062").
		WithRequestParams("department", "HR").
		Build()

employeeList := []Employee{}

resp, err:= client.ExecuteInto(request,&EmployeeList)
```


```

client := requesto.New()
reqTemp, err := requesto.RequestBuilder().
		Get().
		WithURL("http://localhost:8081/employee").
		WithHeaders("Accept", "application/json").
		WithHeaders("user-id", "10062").
		WithRequestParams("department", "HR").
		Build()

request, err := requesto.RequestBuilder().
		WithTemplate(reqTemp).
        	WithRequestParams("post","manager").
        	WithRequestParams("centre","pune")
		Build()

employeeList := []Employee{}

resp, err:= client.ExecuteInto(request,&EmployeeList)
```
