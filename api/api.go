package api

import (
	"net/http"

	"github.com/ASjet/go-curseforge/api/files"
)

type API struct {
	files.ModFiles
	files.ModFile
	files.ModFileChangelog
	files.ModFileUrl
}

func New(t http.RoundTripper) *API {
	return &API{
		ModFiles:         files.NewModFilesAPI(t),
		ModFile:          files.NewModFileAPI(t),
		ModFileChangelog: files.NewModFileChangelogAPI(t),
		ModFileUrl:       files.NewModFileUrlAPI(t),
	}
}
