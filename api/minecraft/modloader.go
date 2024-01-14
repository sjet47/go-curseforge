package minecraft

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ASjet/go-curseforge/schema"
)

type ModLoaderOption func(*ModLoaderRequest)

func NewModLoaderAPI(t http.RoundTripper) ModLoader {
	return func(modLoaderName string, o ...ModLoaderOption) (*schema.ApiResponseOfMinecraftModLoaderVersion, error) {
		r := new(ModLoaderRequest)
		for _, f := range o {
			f(r)
		}
		r.ModLoaderName = modLoaderName
		return schema.UnmarshalResponse[schema.ApiResponseOfMinecraftModLoaderVersion](r.Do(r.ctx, t))
	}
}

type ModLoader func(modLoaderName string, o ...ModLoaderOption) (*schema.ApiResponseOfMinecraftModLoaderVersion, error)

// https://docs.curseforge.com/#get-specific-minecraft-modloader
type ModLoaderRequest struct {
	ctx context.Context

	ModLoaderName string
}

func (r *ModLoaderRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodGet
		path   = fmt.Sprintf("%s/v1/minecraft/modloader/%s", schema.BaseUrl, r.ModLoaderName)
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

func (ModLoader) WithContext(ctx context.Context) ModLoaderOption {
	return func(o *ModLoaderRequest) {
		o.ctx = ctx
	}
}
