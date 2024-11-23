package main

import (
	"fmt"

	"github.com/Rai-Sahil/go-sdk/client"
	"github.com/Rai-Sahil/go-sdk/service/example"
)

func main() {
    clientID := "your-client-id"
    clientSecret := "your-client-secret-key"
    baseURL := "http://localhost:3000"

    sdkClient, err := client.NewClient(clientID, clientSecret, baseURL)
    if err != nil {
        fmt.Println("An error occured", err)
        return
    }

    exampleService := example.NewExampleService(sdkClient)

    endpoint := "example-data"
    data, err := exampleService.GetExampleData(endpoint)
    if err != nil {
        fmt.Println("An error occured", err)
        return
    }

    fmt.Println("Example data: ", data.Message)
}

