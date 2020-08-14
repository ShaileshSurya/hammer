
![alt text](https://github.com/ShaileshSurya/go-images/blob/master/go_pic.jpg?raw=true)

# Requesto  [![Build Status](https://travis-ci.org/ShaileshSurya/requesto.svg?branch=master)](https://travis-ci.org/ShaileshSurya/requesto)


Golang's Fluent HTTP Request Client 



## Recipes

```
client := &requesto.Requesto{}
request, err := requesto.RequestBuilder().
		Get().
		WithURL("http://localhost:8081/employee").
		WithHeaders("Accept", "application/json").
		WithHeaders("user-id", "10062").
		WithRequestParams("department", "HR").
		Build()

resp, err:= client.Execute(request)

```


```
client := &requesto.Requesto{}
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

client := &requesto.Requesto{}
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
