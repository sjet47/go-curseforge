package games

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ASjet/go-curseforge/schema"
	"github.com/ASjet/go-curseforge/schema/enum"
)

type GameVersion2V2Option func(*GameVersionsV2Request)

func NewGameVersionsV2API(t http.RoundTripper) GameVersionsV2 {
	return func(gameID enum.GameID, o ...GameVersion2V2Option) (*schema.GetVersionTypesResponseV2, error) {
		r := new(GameVersionsV2Request)
		for _, f := range o {
			f(r)
		}
		r.GameID = gameID
		return schema.UnmarshalResponse[schema.GetVersionTypesResponseV2](r.Do(r.ctx, t))
	}
}

type GameVersionsV2 func(gameID enum.GameID, o ...GameVersion2V2Option) (*schema.GetVersionTypesResponseV2, error)

// https://docs.curseforge.com/#get-versions-v2
type GameVersionsV2Request struct {
	ctx context.Context

	GameID enum.GameID
}

func (r *GameVersionsV2Request) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodGet
		path   = fmt.Sprintf("%s/v2/games/%d/versions", schema.BaseUrl, r.GameID)
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

func (GameVersionsV2) WithContext(ctx context.Context) GameVersion2V2Option {
	return func(o *GameVersionsV2Request) {
		o.ctx = ctx
	}
}
