package schema

import (
	"time"

	"github.com/ASjet/go-curseforge/schema/enum"
)

// https://docs.curseforge.com/#tocS_SortableGameVersion
type SortableGameVersion struct {
	GameVersionName        string               `json:"gameVersionName"`
	GameVersionPadded      string               `json:"gameVersionPadded"`
	GameVersion            enum.GameVersion     `json:"gameVersion"`
	GameVersionReleaseDate time.Time            `json:"gameVersionReleaseDate"`
	GameVersionTypeID      enum.GameVersionType `json:"gameVersionTypeId"`
}
