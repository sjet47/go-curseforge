package curseforge

import "github.com/sjet47/go-curseforge/api"

type Client struct {
	*api.API
}

func NewClient(apiKey string) *Client {
	return &Client{api.New(apiKeyTransport(apiKey))}
}

func InitDefault(apiKey string) {
	api.InitDefault(apiKeyTransport(apiKey))
}
