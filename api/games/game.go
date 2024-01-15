package games

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ASjet/go-curseforge/schema"
	"github.com/ASjet/go-curseforge/schema/enum"
)

type GameOption func(*GameRequest)

func NewGameAPI(t http.RoundTripper) Game {
	return func(gameID enum.GameID, o ...GameOption) (*schema.GetGameResponse, error) {
		r := new(GameRequest)
		for _, f := range o {
			f(r)
		}
		r.GameID = gameID
		return schema.UnmarshalResponse[schema.GetGameResponse](r.Do(r.ctx, t))
	}
}

type Game func(gameID enum.GameID, o ...GameOption) (*schema.GetGameResponse, error)

// https://docs.curseforge.com/#get-game
type GameRequest struct {
	ctx context.Context

	GameID enum.GameID
}

func (r *GameRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodGet
		path   = fmt.Sprintf("%s/v1/games/%d", schema.BaseUrl, r.GameID)
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

func (Game) WithContext(ctx context.Context) GameOption {
	return func(o *GameRequest) {
		o.ctx = ctx
	}
}
