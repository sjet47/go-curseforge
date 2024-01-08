package curseforge

import "github.com/ASjet/go-curseforge/api"

type Client struct {
	*api.API
}

func NewClient(apiKey string) *Client {
	return &Client{api.New(apiKeyTransport(apiKey))}
}
