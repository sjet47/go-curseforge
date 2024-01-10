package schema

import (
	"encoding/json"
	"net/http"
)

func UnmarshalResponse[T any](rsp *http.Response, err error) (*T, error) {
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	result := new(T)
	if err := json.NewDecoder(rsp.Body).Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

// https://docs.curseforge.com/#tocS_String%20Response
type StringResponse struct {
	Data string `json:"data"`
}

// https://docs.curseforge.com/#tocS_Get%20Mod%20File%20Response
type GetModFileResponse struct {
	Data File `json:"data"`
}

// https://docs.curseforge.com/#tocS_Get%20Mod%20Files%20Response
type GetModFilesResponse struct {
	Data       []File     `json:"data"`
	Pagination Pagination `json:"pagination"`
}

// https://docs.curseforge.com/#tocS_Get%20Featured%20Mods%20Response
type GetFilesResponse struct {
	Data []File `json:"data"`
}
