package minecraft

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ASjet/go-curseforge/schema"
)

type ModLoadersOption func(*ModLoadersRequest)

func NewModLoadersAPI(t http.RoundTripper) ModLoaders {
	return func(o ...ModLoadersOption) (*schema.ApiResponseOfListOfMinecraftModLoaderIndex, error) {
		r := new(ModLoadersRequest)
		for _, f := range o {
			f(r)
		}
		return schema.UnmarshalResponse[schema.ApiResponseOfListOfMinecraftModLoaderIndex](r.Do(r.ctx, t))
	}
}

type ModLoaders func(o ...ModLoadersOption) (*schema.ApiResponseOfListOfMinecraftModLoaderIndex, error)

// https://docs.curseforge.com/#get-minecraft-modloaders
type ModLoadersRequest struct {
	ctx context.Context

	GameVersion *schema.GameVersionStr
	IncludeAll  *bool
}

func (r *ModLoadersRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
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

func (ModLoaders) WithContext(ctx context.Context) ModLoadersOption {
	return func(o *ModLoadersRequest) {
		o.ctx = ctx
	}
}

func (ModLoaders) WithGameVersion(gameVersion schema.GameVersionStr) ModLoadersOption {
	return func(o *ModLoadersRequest) {
		o.GameVersion = &gameVersion
	}
}

func (ModLoaders) WithIncludeAll(includeAll bool) ModLoadersOption {
	return func(o *ModLoadersRequest) {
		o.IncludeAll = &includeAll
	}
}
