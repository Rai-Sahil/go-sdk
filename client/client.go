package client

import (
	"fmt"
	"net/http"
)

type Client struct {
    ClientID     string
    ClientSecret string
    BaseURL      string
    HTTPClient   *http.Client
}

func NewClient(clientID, clientSecret, baseURL string) (*Client, error) {
    if clientID == "" || clientSecret == "" {
        return nil, fmt.Errorf("Client ID and Secret are required")
    }

    return &Client{
        ClientID:     clientID,
        ClientSecret: clientSecret,
        BaseURL:      baseURL,
        HTTPClient:   &http.Client{},
    }, nil
}

func (c *Client) DoRequest(req *http.Request) (*http.Response, error) {
    req.Header.Set("Client-ID", c.ClientID)
    req.Header.Set("Client-Secret", c.ClientSecret)
    req.Header.Set("Content-Type", "application/json")

    resp, err := c.HTTPClient.Do(req)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 400 {
        return nil, fmt.Errorf("Request failed with status code %d", resp.StatusCode)
    }

    return resp, nil
}

