package api

import (
	"net/http"

	"github.com/ASjet/go-curseforge/api/files"
	"github.com/ASjet/go-curseforge/api/minecraft"
	"github.com/ASjet/go-curseforge/api/mods"
)

type API struct {
	mods.SearchMod
	mods.GetMod
	mods.GetMods
	mods.GetFeaturedMods
	mods.GetModDescription

	files.ModFile
	files.ModFiles
	files.Files
	files.ModFileChangelog
	files.ModFileUrl

	minecraft.GetMinecraftVersions
	minecraft.GetMinecraftVersion
	minecraft.GetModLoaders
	minecraft.GetModLoader
}

func New(t http.RoundTripper) *API {
	return &API{
		SearchMod:         mods.NewSearchModAPI(t),
		GetMod:            mods.NewGetModAPI(t),
		GetMods:           mods.NewGetModsAPI(t),
		GetFeaturedMods:   mods.NewGetFeaturedModsAPI(t),
		GetModDescription: mods.NewGetModDescriptionAPI(t),

		ModFile:          files.NewModFileAPI(t),
		ModFiles:         files.NewModFilesAPI(t),
		Files:            files.NewFilesAPI(t),
		ModFileChangelog: files.NewModFileChangelogAPI(t),
		ModFileUrl:       files.NewModFileUrlAPI(t),

		GetMinecraftVersions: minecraft.NewGetMinecraftVersionsAPI(t),
		GetMinecraftVersion:  minecraft.NewGetMinecraftVersionAPI(t),
		GetModLoaders:        minecraft.NewGetModLoadersAPI(t),
		GetModLoader:         minecraft.NewGetModLoaderAPI(t),
	}
}
