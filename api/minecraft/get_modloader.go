package minecraft

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ASjet/go-curseforge/schema"
)

type GetModLoaderOption func(*GetModLoaderRequest)

func NewGetModLoaderAPI(t http.RoundTripper) GetModLoader {
	return func(modLoaderName string, o ...GetModLoaderOption) (*schema.ApiResponseOfMinecraftModLoaderVersion, error) {
		r := new(GetModLoaderRequest)
		for _, f := range o {
			f(r)
		}
		r.ModLoaderName = modLoaderName
		return schema.UnmarshalResponse[schema.ApiResponseOfMinecraftModLoaderVersion](r.Do(r.ctx, t))
	}
}

type GetModLoader func(modLoaderName string, o ...GetModLoaderOption) (*schema.ApiResponseOfMinecraftModLoaderVersion, error)

// https://docs.curseforge.com/#get-specific-minecraft-modloader
type GetModLoaderRequest struct {
	ctx context.Context

	ModLoaderName string
}

func (r *GetModLoaderRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
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

func (GetModLoader) WithContext(ctx context.Context) GetModLoaderOption {
	return func(o *GetModLoaderRequest) {
		o.ctx = ctx
	}
}
