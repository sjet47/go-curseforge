package api

import (
	"net/http"

	"github.com/ASjet/go-curseforge/api/files"
	"github.com/ASjet/go-curseforge/api/mods"
)

type API struct {
	mods.SearchMod

	files.ModFile
	files.ModFiles
	files.Files
	files.ModFileChangelog
	files.ModFileUrl
}

func New(t http.RoundTripper) *API {
	return &API{
		SearchMod: mods.NewSearchModAPI(t),

		ModFile:          files.NewModFileAPI(t),
		ModFiles:         files.NewModFilesAPI(t),
		Files:            files.NewFilesAPI(t),
		ModFileChangelog: files.NewModFileChangelogAPI(t),
		ModFileUrl:       files.NewModFileUrlAPI(t),
	}
}
