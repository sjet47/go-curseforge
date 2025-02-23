package games

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/sjet47/go-curseforge/schema"
)

type GamesOption func(*GamesRequest)

func NewGamesAPI(t http.RoundTripper) Games {
	return func(o ...GamesOption) (*schema.GetGamesResponse, error) {
		r := new(GamesRequest)
		for _, f := range o {
			f(r)
		}
		return schema.UnmarshalResponse[schema.GetGamesResponse](r.Do(r.ctx, t))
	}
}

type Games func(o ...GamesOption) (*schema.GetGamesResponse, error)

// https://docs.curseforge.com/#get-games
type GamesRequest struct {
	ctx context.Context

	Index    int
	PageSize int
}

func (r *GamesRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodGet
		params = make(map[string]string)
		path   = fmt.Sprintf("%s/v1/games", schema.BaseUrl)
	)

	if r.Index > 0 {
		params["index"] = strconv.Itoa(r.Index)
	}
	if r.PageSize > 0 {
		params["pageSize"] = strconv.Itoa(r.PageSize)
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

func (Games) WithContext(ctx context.Context) GamesOption {
	return func(o *GamesRequest) {
		o.ctx = ctx
	}
}

func (Games) WithIndex(index int) GamesOption {
	return func(o *GamesRequest) {
		o.Index = index
	}
}

func (Games) WithPageSize(pageSize int) GamesOption {
	return func(o *GamesRequest) {
		o.PageSize = pageSize
	}
}
