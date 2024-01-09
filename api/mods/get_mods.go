package mods

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ASjet/go-curseforge/schema"
)

func NewGetModsAPI(t http.RoundTripper) GetMods {
	return func(modIDs []schema.ModID, o ...func(*GetModsRequest)) (*http.Response, error) {
		r := new(GetModsRequest)
		for _, f := range o {
			f(r)
		}
		r.ModIDs = modIDs
		return r.Do(r.ctx, t)
	}
}

type GetMods func(modIDs []schema.ModID, o ...func(*GetModsRequest)) (*http.Response, error)

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

func (GetMods) WithContext(ctx context.Context) func(*GetModsRequest) {
	return func(o *GetModsRequest) {
		o.ctx = ctx
	}
}

func (GetMods) WithFilterOnlyPC(onlyPC bool) func(*GetModsRequest) {
	return func(o *GetModsRequest) {
		o.OnlyPC = onlyPC
	}
}
