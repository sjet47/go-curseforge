package schema

type GetModFilesRequestBody struct {
	FileIDs []FileID `json:"fileIds"`
}

type GetModsByIdsListRequestBody struct {
	ModIDs []ModID `json:"modIds"`
	OnlyPC bool    `json:"filterPcOnly"`
}
