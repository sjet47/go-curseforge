package games

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ASjet/go-curseforge/schema"
	"github.com/ASjet/go-curseforge/schema/enum"
)

type GameVersionsOption func(*GameVersionsRequest)

func NewGameVersionsAPI(t http.RoundTripper) GameVersions {
	return func(gameID enum.GameID, o ...GameVersionsOption) (*schema.GetVersionsResponse, error) {
		r := new(GameVersionsRequest)
		for _, f := range o {
			f(r)
		}
		r.GameID = gameID
		return schema.UnmarshalResponse[schema.GetVersionsResponse](r.Do(r.ctx, t))
	}
}

type GameVersions func(gameID enum.GameID, o ...GameVersionsOption) (*schema.GetVersionsResponse, error)

// https://docs.curseforge.com/#get-versions
type GameVersionsRequest struct {
	ctx context.Context

	GameID enum.GameID
}

func (r *GameVersionsRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodGet
		path   = fmt.Sprintf("%s/v1/games/%d/versions", schema.BaseUrl, r.GameID)
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

func (GameVersions) WithContext(ctx context.Context) GameVersionsOption {
	return func(o *GameVersionsRequest) {
		o.ctx = ctx
	}
}
