package curseforge

import (
	"errors"
	"net/http"
)

func apiKeyTransport(apiKey string) http.RoundTripper {
	return &transport{apiKey, http.DefaultTransport}
}

type transport struct {
	apiKey              string
	underlyingTransport http.RoundTripper
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("x-api-key", t.apiKey)
	rsp, err := t.underlyingTransport.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode != http.StatusOK {
		return nil, errors.New(rsp.Status)
	}
	return rsp, nil
}
