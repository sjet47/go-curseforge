package api

import (
	"net/http"

	"github.com/ASjet/go-curseforge/api/files"
	"github.com/ASjet/go-curseforge/api/mods"
)

type API struct {
	mods.SearchMod
	mods.GetMod
	mods.GetMods

	files.ModFile
	files.ModFiles
	files.Files
	files.ModFileChangelog
	files.ModFileUrl
}

func New(t http.RoundTripper) *API {
	return &API{
		SearchMod: mods.NewSearchModAPI(t),
		GetMod:    mods.NewGetModAPI(t),
		GetMods:   mods.NewGetModsAPI(t),

		ModFile:          files.NewModFileAPI(t),
		ModFiles:         files.NewModFilesAPI(t),
		Files:            files.NewFilesAPI(t),
		ModFileChangelog: files.NewModFileChangelogAPI(t),
		ModFileUrl:       files.NewModFileUrlAPI(t),
	}
}
