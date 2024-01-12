package mods

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ASjet/go-curseforge/schema"
)

type GetModsOption func(*GetModsRequest)

func NewGetModsAPI(t http.RoundTripper) GetMods {
	return func(modIDs []schema.ModID, o ...GetModsOption) (*schema.GetModsResponse, error) {
		r := new(GetModsRequest)
		for _, f := range o {
			f(r)
		}
		r.ModIDs = modIDs
		return schema.UnmarshalResponse[schema.GetModsResponse](r.Do(r.ctx, t))
	}
}

type GetMods func(modIDs []schema.ModID, o ...GetModsOption) (*schema.GetModsResponse, error)

// https://docs.curseforge.com/#get-mods
type GetModsRequest struct {
	ctx context.Context

	schema.GetModsByIdsListRequestBody
}

func (r *GetModsRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodPost
		path   = fmt.Sprintf("%s/v1/mods", schema.BaseUrl)
	)

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(&r.GetModsByIdsListRequestBody); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, path, buf)
	if err != nil {
		return nil, err
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	req.Header.Add("Content-Type", "application/json")

	return t.RoundTrip(req)
}

func (GetMods) WithContext(ctx context.Context) GetModsOption {
	return func(o *GetModsRequest) {
		o.ctx = ctx
	}
}

func (GetMods) WithFilterOnlyPC(onlyPC bool) GetModsOption {
	return func(o *GetModsRequest) {
		o.OnlyPC = onlyPC
	}
}
