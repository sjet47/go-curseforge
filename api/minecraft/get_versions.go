package minecraft

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ASjet/go-curseforge/schema"
)

type GetMinecraftVersionsOption func(*GetMinecraftVersionsRequest)

func NewGetMinecraftVersionsAPI(t http.RoundTripper) GetMinecraftVersions {
	return func(o ...GetMinecraftVersionsOption) (*schema.ApiResponseOfListOfMinecraftGameVersion, error) {
		r := new(GetMinecraftVersionsRequest)
		for _, f := range o {
			f(r)
		}
		return schema.UnmarshalResponse[schema.ApiResponseOfListOfMinecraftGameVersion](r.Do(r.ctx, t))
	}
}

type GetMinecraftVersions func(o ...GetMinecraftVersionsOption) (*schema.ApiResponseOfListOfMinecraftGameVersion, error)

// https://docs.curseforge.com/#get-minecraft-versions
type GetMinecraftVersionsRequest struct {
	ctx context.Context

	SortDescending *bool
}

func (r *GetMinecraftVersionsRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
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

func (GetMinecraftVersions) WithContext(ctx context.Context) GetMinecraftVersionsOption {
	return func(o *GetMinecraftVersionsRequest) {
		o.ctx = ctx
	}
}

func (GetMinecraftVersions) WithSortDescending(isDescending bool) GetMinecraftVersionsOption {
	return func(o *GetMinecraftVersionsRequest) {
		o.SortDescending = &isDescending
	}
}
