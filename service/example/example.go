package example

import (
	"fmt"

	"github.com/Rai-Sahil/go-sdk/client"
)

type ExampleService struct {
    Client *client.Client
}

type ExampleResponse struct {
    Message string `json:"message"`
}

func NewExampleService(client *client.Client) *ExampleService {
    return &ExampleService {
        Client: client,
    }
}

func (e *ExampleService) GetExampleData(endpoint string) (*ExampleResponse, error) {
    url := fmt.Sprintf("%s/%s", e.Client.BaseURL, endpoint)

    return &ExampleResponse{
        Message: url,
    }, nil
}

