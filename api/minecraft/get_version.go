package minecraft

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ASjet/go-curseforge/schema"
	"github.com/ASjet/go-curseforge/schema/enum"
)

type GetMinecraftVersionOption func(*GetMinecraftVersionRequest)

func NewGetMinecraftVersionAPI(t http.RoundTripper) GetMinecraftVersion {
	return func(gameVersion enum.GameVersion, o ...GetMinecraftVersionOption) (*schema.ApiResponseOfMinecraftGameVersion, error) {
		r := new(GetMinecraftVersionRequest)
		for _, f := range o {
			f(r)
		}
		r.GameVersion = gameVersion
		return schema.UnmarshalResponse[schema.ApiResponseOfMinecraftGameVersion](r.Do(r.ctx, t))
	}
}

type GetMinecraftVersion func(gameVersion enum.GameVersion, o ...GetMinecraftVersionOption) (*schema.ApiResponseOfMinecraftGameVersion, error)

// https://docs.curseforge.com/#get-specific-minecraft-version
type GetMinecraftVersionRequest struct {
	ctx context.Context

	GameVersion enum.GameVersion
}

func (r *GetMinecraftVersionRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodGet
		path   = fmt.Sprintf("%s/v1/minecraft/version/%s", schema.BaseUrl, r.GameVersion)
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

func (GetMinecraftVersion) WithContext(ctx context.Context) GetMinecraftVersionOption {
	return func(o *GetMinecraftVersionRequest) {
		o.ctx = ctx
	}
}
