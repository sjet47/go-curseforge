package api

import (
	"net/http"
)

type API struct {
}

func New(t http.RoundTripper) *API {
	return &API{}
}
