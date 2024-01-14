package mods

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ASjet/go-curseforge/schema"
)

type ModsOption func(*ModsRequest)

func NewModsAPI(t http.RoundTripper) Mods {
	return func(modIDs []schema.ModID, o ...ModsOption) (*schema.GetModsResponse, error) {
		r := new(ModsRequest)
		for _, f := range o {
			f(r)
		}
		r.ModIDs = modIDs
		return schema.UnmarshalResponse[schema.GetModsResponse](r.Do(r.ctx, t))
	}
}

type Mods func(modIDs []schema.ModID, o ...ModsOption) (*schema.GetModsResponse, error)

// https://docs.curseforge.com/#get-mods
type ModsRequest struct {
	ctx context.Context

	schema.GetModsByIdsListRequestBody
}

func (r *ModsRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
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

func (Mods) WithContext(ctx context.Context) ModsOption {
	return func(o *ModsRequest) {
		o.ctx = ctx
	}
}

func (Mods) WithFilterOnlyPC(onlyPC bool) ModsOption {
	return func(o *ModsRequest) {
		o.OnlyPC = onlyPC
	}
}
