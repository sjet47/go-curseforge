package api

import (
	"net/http"

	"github.com/ASjet/go-curseforge/api/categories"
	"github.com/ASjet/go-curseforge/api/files"
	"github.com/ASjet/go-curseforge/api/fingerprints"
	"github.com/ASjet/go-curseforge/api/games"
	"github.com/ASjet/go-curseforge/api/minecraft"
	"github.com/ASjet/go-curseforge/api/mods"
)

var (
	Games            games.Games
	Game             games.Game
	GameVersions     games.GameVersions
	GameVersionTypes games.GameVersionTypes
	GameVersionsV2   games.GameVersionsV2

	Categories categories.Categories

	SearchMod      mods.SearchMod
	Mod            mods.Mod
	Mods           mods.Mods
	FeaturedMods   mods.FeaturedMods
	ModDescription mods.ModDescription

	ModFile          files.ModFile
	ModFiles         files.ModFiles
	Files            files.Files
	ModFileChangelog files.ModFileChangelog
	ModFileUrl       files.ModFileUrl

	FingerprintMatchesByGame      fingerprints.FingerprintMatchesByGame
	FingerprintMatches            fingerprints.FingerprintMatches
	FingerprintFuzzyMatchesByGame fingerprints.FingerprintFuzzyMatchesByGame
	FingerprintFuzzyMatches       fingerprints.FingerprintFuzzyMatches

	MinecraftVersions minecraft.MinecraftVersions
	MinecraftVersion  minecraft.MinecraftVersion
	ModLoaders        minecraft.ModLoaders
	ModLoader         minecraft.ModLoader
)

func InitDefault(t http.RoundTripper) {
	d := New(t)

	Games = d.Games
	Game = d.Game
	GameVersions = d.GameVersions
	GameVersionTypes = d.GameVersionTypes
	GameVersionsV2 = d.GameVersionsV2

	Categories = d.Categories

	SearchMod = d.SearchMod
	Mod = d.Mod
	Mods = d.Mods
	FeaturedMods = d.FeaturedMods
	ModDescription = d.ModDescription

	ModFile = d.ModFile
	ModFiles = d.ModFiles
	Files = d.Files
	ModFileChangelog = d.ModFileChangelog
	ModFileUrl = d.ModFileUrl

	FingerprintMatchesByGame = d.FingerprintMatchesByGame
	FingerprintMatches = d.FingerprintMatches
	FingerprintFuzzyMatchesByGame = d.FingerprintFuzzyMatchesByGame
	FingerprintFuzzyMatches = d.FingerprintFuzzyMatches

	MinecraftVersions = d.MinecraftVersions
	MinecraftVersion = d.MinecraftVersion
	ModLoaders = d.ModLoaders
	ModLoader = d.ModLoader
}

type API struct {
	games.Games
	games.Game
	games.GameVersions
	games.GameVersionTypes
	games.GameVersionsV2

	categories.Categories

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

	fingerprints.FingerprintMatchesByGame
	fingerprints.FingerprintMatches
	fingerprints.FingerprintFuzzyMatchesByGame
	fingerprints.FingerprintFuzzyMatches

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

		Categories: categories.NewCategoriesAPI(t),

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

		FingerprintMatchesByGame:      fingerprints.NewFingerMatchesByGameAPI(t),
		FingerprintMatches:            fingerprints.NewFingerprintMatchesAPI(t),
		FingerprintFuzzyMatchesByGame: fingerprints.NewFingerprintFuzzyMatchesByGameAPI(t),
		FingerprintFuzzyMatches:       fingerprints.NewFingerprintFuzzyMatchesAPI(t),

		MinecraftVersions: minecraft.NewMinecraftVersionsAPI(t),
		MinecraftVersion:  minecraft.NewMinecraftVersionAPI(t),
		ModLoaders:        minecraft.NewModLoadersAPI(t),
		ModLoader:         minecraft.NewModLoaderAPI(t),
	}
}
