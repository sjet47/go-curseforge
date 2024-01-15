package games

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ASjet/go-curseforge/schema"
	"github.com/ASjet/go-curseforge/schema/enum"
)

type GameVersionTypesOption func(*GameVersionTypesRequest)

func NewGameVersionTypesAPI(t http.RoundTripper) GameVersionTypes {
	return func(gameID enum.GameID, o ...GameVersionTypesOption) (*schema.GetVersionTypesResponse, error) {
		r := new(GameVersionTypesRequest)
		for _, f := range o {
			f(r)
		}
		r.GameID = gameID
		return schema.UnmarshalResponse[schema.GetVersionTypesResponse](r.Do(r.ctx, t))
	}
}

type GameVersionTypes func(gameID enum.GameID, o ...GameVersionTypesOption) (*schema.GetVersionTypesResponse, error)

// https://docs.curseforge.com/#get-version-types
type GameVersionTypesRequest struct {
	ctx context.Context

	GameID enum.GameID
}

func (r *GameVersionTypesRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodGet
		path   = fmt.Sprintf("%s/v1/games/%d/version-types", schema.BaseUrl, r.GameID)
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

func (GameVersionTypes) WithContext(ctx context.Context) GameVersionTypesOption {
	return func(o *GameVersionTypesRequest) {
		o.ctx = ctx
	}
}
