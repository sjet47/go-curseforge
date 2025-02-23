package schema

import "github.com/sjet47/go-curseforge/schema/enum"

// https://docs.curseforge.com/#tocS_GetModFilesRequestBody
type GetModFilesRequestBody struct {
	FileIDs []FileID `json:"fileIds"`
}

// https://docs.curseforge.com/#tocS_GetModsByIdsListRequestBody
type GetModsByIdsListRequestBody struct {
	ModIDs []ModID `json:"modIds"`
	OnlyPC bool    `json:"filterPcOnly"`
}

// https://docs.curseforge.com/#tocS_GetFeaturedModsRequestBody
type GetFeaturedModsRequestBody struct {
	GameID            enum.GameID       `json:"gameId"`
	ExcludedModIDs    []ModID           `json:"excludedModIds"`
	GameVersionTypeID GameVersionTypeID `json:"gameVersionTypeId"`
}

// https://docs.curseforge.com/#tocS_GetFingerprintMatchesRequestBody
type GetFingerprintMatchesRequestBody struct {
	Fingerprints []Fingerprint `json:"fingerprints"`
}

// https://docs.curseforge.com/#tocS_GetFuzzyMatchesRequestBody
type GetFuzzyMatchesRequestBody struct {
	GameID       enum.GameID         `json:"gameId"`
	Fingerprints []FolderFingerprint `json:"fingerprints"`
}
