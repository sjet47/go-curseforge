package api

import (
	"net/http"

	"github.com/ASjet/go-curseforge/api/files"
)

type API struct {
	files.Files
}

func New(t http.RoundTripper) *API {
	return &API{
		Files: files.NewFilesAPI(t),
	}
}
