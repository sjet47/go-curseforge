package mods

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ASjet/go-curseforge/schema"
	"github.com/ASjet/go-curseforge/schema/enum"
)

type FeaturedModsOption func(*FeaturedModsRequest)

func NewFeaturedModsAPI(t http.RoundTripper) FeaturedMods {
	return func(gameID enum.GameID, o ...FeaturedModsOption) (*schema.GetFeaturedModsRequestBody, error) {
		r := new(FeaturedModsRequest)
		for _, f := range o {
			f(r)
		}
		r.GameID = gameID
		return schema.UnmarshalResponse[schema.GetFeaturedModsRequestBody](r.Do(r.ctx, t))
	}
}

type FeaturedMods func(gameID enum.GameID, o ...FeaturedModsOption) (*schema.GetFeaturedModsRequestBody, error)

// https://docs.curseforge.com/#get-featured-mods
type FeaturedModsRequest struct {
	ctx context.Context

	schema.GetFeaturedModsRequestBody
}

func (r *FeaturedModsRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodPost
		path   = fmt.Sprintf("%s/v1/mods/featured", schema.BaseUrl)
	)

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(&r.GetFeaturedModsRequestBody); err != nil {
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

func (FeaturedMods) WithContext(ctx context.Context) FeaturedModsOption {
	return func(o *FeaturedModsRequest) {
		o.ctx = ctx
	}
}

func (FeaturedMods) WithExcludedModIDs(modID ...schema.ModID) FeaturedModsOption {
	return func(o *FeaturedModsRequest) {
		o.ExcludedModIDs = modID
	}
}

func (FeaturedMods) WithGameVersionTypeID(gameVersionTypeID schema.GameVersionTypeID) FeaturedModsOption {
	return func(o *FeaturedModsRequest) {
		o.GameVersionTypeID = gameVersionTypeID
	}
}
