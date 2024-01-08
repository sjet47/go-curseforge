package curseforge

import "net/http"

func apiKeyTransport(apiKey string) http.RoundTripper {
	return &transport{apiKey, http.DefaultTransport}
}

type transport struct {
	apiKey              string
	underlyingTransport http.RoundTripper
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("x-api-key", t.apiKey)
	return t.underlyingTransport.RoundTrip(req)
}
