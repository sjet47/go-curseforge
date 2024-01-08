package api

import (
	"net/http"

	"github.com/ASjet/go-curseforge/api/files"
)

type API struct {
	files.ModFile
	files.ModFiles
	files.Files
	files.ModFileChangelog
	files.ModFileUrl
}

func New(t http.RoundTripper) *API {
	return &API{
		ModFile:          files.NewModFileAPI(t),
		ModFiles:         files.NewModFilesAPI(t),
		Files:            files.NewFilesAPI(t),
		ModFileChangelog: files.NewModFileChangelogAPI(t),
		ModFileUrl:       files.NewModFileUrlAPI(t),
	}
}
