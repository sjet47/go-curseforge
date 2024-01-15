package schema

import (
	"time"

	"github.com/ASjet/go-curseforge/schema/enum"
)

// https://docs.curseforge.com/#tocS_SortableGameVersion
type SortableGameVersion struct {
	GameVersionName        string                 `json:"gameVersionName"`
	GameVersionPadded      string                 `json:"gameVersionPadded"`
	GameVersion            enum.GameVersion       `json:"gameVersion"`
	GameVersionReleaseDate time.Time              `json:"gameVersionReleaseDate"`
	GameVersionTypeID      enum.GameVersionTypeID `json:"gameVersionTypeId"`
}

// https://docs.curseforge.com/#tocS_Game
type Game struct {
	ID           enum.GameID        `json:"id"`
	Name         string             `json:"name"`
	Slug         string             `json:"slug"`
	DateModified time.Time          `json:"dateModified"`
	Assets       GameAssets         `json:"assets"`
	Status       enum.CoreStatus    `json:"status"`
	ApiStatus    enum.CoreApiStatus `json:"apiStatus"`
}

// https://docs.curseforge.com/#tocS_GameAssets
type GameAssets struct {
	IconURL  string `json:"iconUrl"`
	TitleURL string `json:"titleUrl"`
	CoverURL string `json:"coverUrl"`
}

// https://docs.curseforge.com/#tocS_GameVersion
type GameVersion struct {
	ID   GameVersionID `json:"id"`
	Slug string        `json:"slug"`
	Name string        `json:"name"`
}

// https://docs.curseforge.com/#tocS_GameVersionsByType
type GameVersionsByType struct {
	Type     enum.GameVersionTypeID `json:"type"`
	Versions []string               `json:"versions"`
}

// https://docs.curseforge.com/#tocS_GameVersionsByType2
type GameVersionsByType2 struct {
	Type     enum.GameVersionTypeID `json:"type"`
	Versions []GameVersion          `json:"versions"`
}

// https://docs.curseforge.com/#tocS_GameVersionType
type GameVersionType struct {
	ID         enum.GameVersionTypeID     `json:"id"`
	GameID     enum.GameID                `json:"gameId"`
	Name       string                     `json:"name"`
	Slug       string                     `json:"slug"`
	IsSyncable bool                       `json:"isSyncable"`
	Status     enum.GameVersionTypeStatus `json:"status"`
}
