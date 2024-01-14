package mods

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ASjet/go-curseforge/schema"
)

type ModDescriptionOption func(*ModDescriptionRequest)

func NewModDescriptionAPI(t http.RoundTripper) ModDescription {
	return func(modID schema.ModID, o ...ModDescriptionOption) (*schema.StringResponse, error) {
		r := new(ModDescriptionRequest)
		for _, f := range o {
			f(r)
		}
		r.ModID = modID
		return schema.UnmarshalResponse[schema.StringResponse](r.Do(r.ctx, t))
	}
}

type ModDescription func(modID schema.ModID, o ...ModDescriptionOption) (*schema.StringResponse, error)

// https://docs.curseforge.com/#get-mod-description
type ModDescriptionRequest struct {
	ctx context.Context

	ModID    schema.ModID
	Raw      *bool
	Stripped *bool
	Markup   *bool
}

func (r *ModDescriptionRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodGet
		params = make(map[string]string)
		path   = fmt.Sprintf("%s/v1/mods/%d/description", schema.BaseUrl, r.ModID)
	)

	if r.Raw != nil {
		params["raw"] = strconv.FormatBool(*r.Raw)
	}
	if r.Stripped != nil {
		params["stripped"] = strconv.FormatBool(*r.Stripped)
	}
	if r.Markup != nil {
		params["markup"] = strconv.FormatBool(*r.Markup)
	}

	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		return nil, err
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return t.RoundTrip(req)
}

func (ModDescription) WithContext(ctx context.Context) ModDescriptionOption {
	return func(o *ModDescriptionRequest) {
		o.ctx = ctx
	}
}

func (ModDescription) WithRawContent(raw bool) ModDescriptionOption {
	return func(o *ModDescriptionRequest) {
		o.Raw = &raw
	}
}

func (ModDescription) WithStrippedContent(stripped bool) ModDescriptionOption {
	return func(o *ModDescriptionRequest) {
		o.Stripped = &stripped
	}
}

func (ModDescription) WithMarkupContent(markup bool) ModDescriptionOption {
	return func(o *ModDescriptionRequest) {
		o.Markup = &markup
	}
}
