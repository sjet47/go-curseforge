package mods

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ASjet/go-curseforge/schema"
)

type ModOption func(*ModRequest)

func NewModAPI(t http.RoundTripper) Mod {
	return func(modID schema.ModID, o ...ModOption) (*schema.GetModResponse, error) {
		r := new(ModRequest)
		for _, f := range o {
			f(r)
		}
		r.ModID = modID
		return schema.UnmarshalResponse[schema.GetModResponse](r.Do(r.ctx, t))
	}
}

type Mod func(modID schema.ModID, o ...ModOption) (*schema.GetModResponse, error)

// https://docs.curseforge.com/#get-mod
type ModRequest struct {
	ctx context.Context

	ModID schema.ModID
}

func (r *ModRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodGet
		path   = fmt.Sprintf("%s/v1/mods/%d", schema.BaseUrl, r.ModID)
	)

	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		return nil, err
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return t.RoundTrip(req)
}

func (Mod) WithContext(ctx context.Context) ModOption {
	return func(o *ModRequest) {
		o.ctx = ctx
	}
}
