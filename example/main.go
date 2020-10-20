package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ShaileshSurya/hammer"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// init client
	client := hammer.New()

	// build request
	request, err := hammer.RequestBuilder().
		Get().
		WithURL("http://echo.jsontest.com/Greeting/hello/place/world").
		WithContext(context.Background()).
		WithHeaders("Accept", "application/json").
		Build()
	handleErr(err)

	// Execute the request and manually handle response
	resp, err := client.Execute(request)
	handleErr(err)

	body, err := ioutil.ReadAll(resp.Body)
	handleErr(err)

	type HelloWorld struct {
		Greeting string
		Place    string
	}
	var hw HelloWorld

	err = json.Unmarshal(body, &hw)
	handleErr(err)

	fmt.Printf("%+v\n\n", hw)

	// Execute the request and feed response body into an empty interface
	m := make(map[string]interface{})
	_ = client.ExecuteInto(request, &m)
	handleErr(err)

	fmt.Printf("%+v\n\n", m)

	// Execute the request and feed response body into a struct
	// (because a every struct implements interface{})
	hw2 := HelloWorld{}
	err = client.ExecuteInto(request, &hw2)
	handleErr(err)

	fmt.Printf("%+v\n\n", hw2)
}
