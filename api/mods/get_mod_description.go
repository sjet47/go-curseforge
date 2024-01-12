package mods

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ASjet/go-curseforge/schema"
)

type GetModDescriptionOption func(*GetModDescriptionRequest)

func NewGetModDescriptionAPI(t http.RoundTripper) GetModDescription {
	return func(modID schema.ModID, o ...GetModDescriptionOption) (*schema.StringResponse, error) {
		r := new(GetModDescriptionRequest)
		for _, f := range o {
			f(r)
		}
		r.ModID = modID
		return schema.UnmarshalResponse[schema.StringResponse](r.Do(r.ctx, t))
	}
}

type GetModDescription func(modID schema.ModID, o ...GetModDescriptionOption) (*schema.StringResponse, error)

// https://docs.curseforge.com/#get-mod-description
type GetModDescriptionRequest struct {
	ctx context.Context

	ModID    schema.ModID
	Raw      *bool
	Stripped *bool
	Markup   *bool
}

func (r *GetModDescriptionRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
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

func (GetModDescription) WithContext(ctx context.Context) GetModDescriptionOption {
	return func(o *GetModDescriptionRequest) {
		o.ctx = ctx
	}
}

func (GetModDescription) WithRawContent(raw bool) GetModDescriptionOption {
	return func(o *GetModDescriptionRequest) {
		o.Raw = &raw
	}
}

func (GetModDescription) WithStrippedContent(stripped bool) GetModDescriptionOption {
	return func(o *GetModDescriptionRequest) {
		o.Stripped = &stripped
	}
}

func (GetModDescription) WithMarkupContent(markup bool) GetModDescriptionOption {
	return func(o *GetModDescriptionRequest) {
		o.Markup = &markup
	}
}
