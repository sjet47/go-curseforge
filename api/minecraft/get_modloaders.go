package minecraft

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ASjet/go-curseforge/schema"
	"github.com/ASjet/go-curseforge/schema/enum"
)

type GetModLoadersOption func(*GetModLoadersRequest)

func NewGetModLoadersAPI(t http.RoundTripper) GetModLoaders {
	return func(o ...GetModLoadersOption) (*schema.ApiResponseOfListOfMinecraftModLoaderIndex, error) {
		r := new(GetModLoadersRequest)
		for _, f := range o {
			f(r)
		}
		return schema.UnmarshalResponse[schema.ApiResponseOfListOfMinecraftModLoaderIndex](r.Do(r.ctx, t))
	}
}

type GetModLoaders func(o ...GetModLoadersOption) (*schema.ApiResponseOfListOfMinecraftModLoaderIndex, error)

// https://docs.curseforge.com/#get-minecraft-modloaders
type GetModLoadersRequest struct {
	ctx context.Context

	GameVersion *enum.GameVersion
	IncludeAll  *bool
}

func (r *GetModLoadersRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodGet
		params = make(map[string]string)
		path   = fmt.Sprintf("%s/v1/minecraft/modloader", schema.BaseUrl)
	)

	if r.GameVersion != nil {
		params["version"] = r.GameVersion.Param()
	}
	if r.IncludeAll != nil {
		params["includeAll"] = strconv.FormatBool(*r.IncludeAll)
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

func (GetModLoaders) WithContext(ctx context.Context) GetModLoadersOption {
	return func(o *GetModLoadersRequest) {
		o.ctx = ctx
	}
}

func (GetModLoaders) WithGameVersion(gameVersion enum.GameVersion) GetModLoadersOption {
	return func(o *GetModLoadersRequest) {
		o.GameVersion = &gameVersion
	}
}

func (GetModLoaders) WithIncludeAll(includeAll bool) GetModLoadersOption {
	return func(o *GetModLoadersRequest) {
		o.IncludeAll = &includeAll
	}
}
