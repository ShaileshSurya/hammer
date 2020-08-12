# Requesto

Golang's Fluent HTTP Request Client 


## Recipes
```
    // This pattern 1 
	client := &requesto.Requesto{}
	resp, err := client.Get("http://localhost:8081/employee").
		WithHeaders("Accept", "application/json").
		WithHeaders("user-id", "10000").
		Execute()


    // Fetch the response into the employee struct. 
    client := &requesto.Requesto{}
	resp, err := client.Get("http://localhost:8081/employee").
		WithHeaders("Accept", "application/json").
        	WithHeaders("user-id", "10000").
        	WithParams("department","HR")
		Into(&employee{}).
		Execute()
```

```
// This is a pattern 2 >> Preffered. 

client := &requesto.Requesto{}
request, err := requesto.RequestBuilder().
		Get().
		WithURL("http://localhost:8081/employee").
		WithHeaders("Accept", "application/json").
		WithHeaders("user-id", "10062").
		WithParams("department", "HR").
		Build()

resp, err:= client.Execute(request)

```


```
// This is a pattern 2 >> Preffered. 

client := &requesto.Requesto{}
request, err := requesto.RequestBuilder().
		Get().
		WithURL("http://localhost:8081/employee").
		WithHeaders("Accept", "application/json").
		WithHeaders("user-id", "10062").
		WithParams("department", "HR").
		Build()

employeeList := []Employee{}

resp, err:= client.FetchInto(&employeeList).Execute(request)

or 

resp, err:= client.ExecuteInto(request,&EmployeeList)
```


```

client := &requesto.Requesto{}
reqTemp, err := requesto.RequestBuilder().
		Get().
		WithURL("http://localhost:8081/employee").
		WithHeaders("Accept", "application/json").
		WithHeaders("user-id", "10062").
		WithParams("department", "HR").
		Build()

request, err := requesto.RequestBuilder().
		WithTemplate(reqTemp).
        	WithParams("post","manager").
        	WithParams("centre","pune")
		Build()

employeeList := []Employee{}

resp, err:= client.FetchInto(&employeeList).Execute(request)

```
