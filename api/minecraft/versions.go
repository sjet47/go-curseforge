package minecraft

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ASjet/go-curseforge/schema"
)

type MinecraftVersionsOption func(*MinecraftVersionsRequest)

func NewMinecraftVersionsAPI(t http.RoundTripper) MinecraftVersions {
	return func(o ...MinecraftVersionsOption) (*schema.ApiResponseOfListOfMinecraftGameVersion, error) {
		r := new(MinecraftVersionsRequest)
		for _, f := range o {
			f(r)
		}
		return schema.UnmarshalResponse[schema.ApiResponseOfListOfMinecraftGameVersion](r.Do(r.ctx, t))
	}
}

type MinecraftVersions func(o ...MinecraftVersionsOption) (*schema.ApiResponseOfListOfMinecraftGameVersion, error)

// https://docs.curseforge.com/#get-minecraft-versions
type MinecraftVersionsRequest struct {
	ctx context.Context

	SortDescending *bool
}

func (r *MinecraftVersionsRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodGet
		params = make(map[string]string)
		path   = fmt.Sprintf("%s/v1/minecraft/version", schema.BaseUrl)
	)

	if r.SortDescending != nil {
		params["sortDescending"] = strconv.FormatBool(*r.SortDescending)
	}

	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		return nil, err
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	if len(params) > 0 {
		query := req.URL.Query()
		for k, v := range params {
			query.Add(k, v)
		}
		req.URL.RawQuery = query.Encode()
	}

	return t.RoundTrip(req)
}

func (MinecraftVersions) WithContext(ctx context.Context) MinecraftVersionsOption {
	return func(o *MinecraftVersionsRequest) {
		o.ctx = ctx
	}
}

func (MinecraftVersions) WithSortDescending(isDescending bool) MinecraftVersionsOption {
	return func(o *MinecraftVersionsRequest) {
		o.SortDescending = &isDescending
	}
}
