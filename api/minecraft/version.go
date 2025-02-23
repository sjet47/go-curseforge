package minecraft

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sjet47/go-curseforge/schema"
)

type MinecraftVersionOption func(*MinecraftVersionRequest)

func NewMinecraftVersionAPI(t http.RoundTripper) MinecraftVersion {
	return func(gameVersion schema.GameVersionStr, o ...MinecraftVersionOption) (*schema.ApiResponseOfMinecraftGameVersion, error) {
		r := new(MinecraftVersionRequest)
		for _, f := range o {
			f(r)
		}
		r.GameVersion = gameVersion
		return schema.UnmarshalResponse[schema.ApiResponseOfMinecraftGameVersion](r.Do(r.ctx, t))
	}
}

type MinecraftVersion func(gameVersion schema.GameVersionStr, o ...MinecraftVersionOption) (*schema.ApiResponseOfMinecraftGameVersion, error)

// https://docs.curseforge.com/#get-specific-minecraft-version
type MinecraftVersionRequest struct {
	ctx context.Context

	GameVersion schema.GameVersionStr
}

func (r *MinecraftVersionRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
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

func (MinecraftVersion) WithContext(ctx context.Context) MinecraftVersionOption {
	return func(o *MinecraftVersionRequest) {
		o.ctx = ctx
	}
}
