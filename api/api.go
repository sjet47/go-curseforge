package api

import (
	"net/http"

	"github.com/ASjet/go-curseforge/api/files"
)

type API struct {
	files.Files
	files.File
}

func New(t http.RoundTripper) *API {
	return &API{
		Files: files.NewFilesAPI(t),
		File:  files.NewFileAPI(t),
	}
}
