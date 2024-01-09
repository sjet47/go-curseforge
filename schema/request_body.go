package schema

import "github.com/ASjet/go-curseforge/schema/enum"

type GetModFilesRequestBody struct {
	FileIDs []FileID `json:"fileIds"`
}

type GetModsByIdsListRequestBody struct {
	ModIDs []ModID `json:"modIds"`
	OnlyPC bool    `json:"filterPcOnly"`
}

type GetFeaturedModsRequestBody struct {
	GameID            enum.GameID          `json:"gameId"`
	ExcludedModIDs    []ModID              `json:"excludedModIds"`
	GameVersionTypeID enum.GameVersionType `json:"gameVersionTypeId"`
}
