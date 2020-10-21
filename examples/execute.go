package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ShaileshSurya/hammer"
)

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
	if err != nil {
		log.Fatal(err)
	}

	// Execute the request and manually handle response
	resp, err := client.Execute(request)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	type HelloWorld struct {
		Greeting string
		Place    string
	}
	var hw HelloWorld

	err = json.Unmarshal(body, &hw)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", hw)
}
