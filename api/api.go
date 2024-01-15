package api

import (
	"net/http"

	"github.com/ASjet/go-curseforge/api/files"
	"github.com/ASjet/go-curseforge/api/games"
	"github.com/ASjet/go-curseforge/api/minecraft"
	"github.com/ASjet/go-curseforge/api/mods"
)

type API struct {
	games.Games
	games.Game
	games.GameVersions
	games.GameVersionTypes
	games.GameVersionsV2

	mods.SearchMod
	mods.Mod
	mods.Mods
	mods.FeaturedMods
	mods.ModDescription

	files.ModFile
	files.ModFiles
	files.Files
	files.ModFileChangelog
	files.ModFileUrl

	minecraft.MinecraftVersions
	minecraft.MinecraftVersion
	minecraft.ModLoaders
	minecraft.ModLoader
}

func New(t http.RoundTripper) *API {
	return &API{
		Games:            games.NewGamesAPI(t),
		Game:             games.NewGameAPI(t),
		GameVersions:     games.NewGameVersionsAPI(t),
		GameVersionTypes: games.NewGameVersionTypesAPI(t),
		GameVersionsV2:   games.NewGameVersionsV2API(t),

		SearchMod:      mods.NewSearchModAPI(t),
		Mod:            mods.NewModAPI(t),
		Mods:           mods.NewModsAPI(t),
		FeaturedMods:   mods.NewFeaturedModsAPI(t),
		ModDescription: mods.NewModDescriptionAPI(t),

		ModFile:          files.NewModFileAPI(t),
		ModFiles:         files.NewModFilesAPI(t),
		Files:            files.NewFilesAPI(t),
		ModFileChangelog: files.NewModFileChangelogAPI(t),
		ModFileUrl:       files.NewModFileUrlAPI(t),

		MinecraftVersions: minecraft.NewMinecraftVersionsAPI(t),
		MinecraftVersion:  minecraft.NewMinecraftVersionAPI(t),
		ModLoaders:        minecraft.NewModLoadersAPI(t),
		ModLoader:         minecraft.NewModLoaderAPI(t),
	}
}
