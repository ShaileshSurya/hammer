package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ShaileshSurya/hammer"
)

func main() {
	client := hammer.New()

	request, err := hammer.RequestBuilder().
		Get().
		WithURL("http://echo.jsontest.com/Greeting/hello/place/world").
		WithContext(context.Background()).
		WithHeaders("Accept", "application/json").
		Build()
	if err != nil {
		log.Fatal(err)
	}

	// ExecuteInto a map[string]interface{}
	responseMap := make(map[string]interface{})
	_ = client.ExecuteInto(request, &responseMap)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n\n", responseMap)

	type HelloWorld struct {
		Greeting string
		Place    string
	}

	// ExecuteInto a struct
	var hw HelloWorld
	err = client.ExecuteInto(request, &hw)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", hw)
}
