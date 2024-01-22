package fingerprints

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ASjet/go-curseforge/schema"
	"github.com/ASjet/go-curseforge/schema/enum"
)

type FingerprintMatchesByGameOption func(*FingerprintMatchesByGameRequest)

func NewFingerMatchesByGameAPI(t http.RoundTripper) FingerprintMatchesByGame {
	return func(gameID enum.GameID, fingerprints []schema.Fingerprint,
		o ...FingerprintMatchesByGameOption) (*schema.GetFingerprintMatchesResponse, error) {
		r := new(FingerprintMatchesByGameRequest)
		for _, f := range o {
			f(r)
		}
		r.GameID = gameID
		r.Fingerprints = fingerprints
		return schema.UnmarshalResponse[schema.GetFingerprintMatchesResponse](r.Do(r.ctx, t))
	}
}

type FingerprintMatchesByGame func(gameID enum.GameID, fingerprints []schema.Fingerprint, o ...FingerprintMatchesByGameOption) (*schema.GetFingerprintMatchesResponse, error)

// https://docs.curseforge.com/#get-fingerprints-matches-by-game-id
type FingerprintMatchesByGameRequest struct {
	ctx context.Context

	GameID enum.GameID
	schema.GetFingerprintMatchesRequestBody
}

func (r *FingerprintMatchesByGameRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodPost
		path   = fmt.Sprintf("%s/v1/fingerprints/%d", schema.BaseUrl, r.GameID)
	)

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(&r.GetFingerprintMatchesRequestBody); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, path, buf)
	if err != nil {
		return nil, err
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	req.Header.Add("Content-Type", "application/json")

	return t.RoundTrip(req)
}

func (FingerprintMatchesByGame) WithContext(ctx context.Context) FingerprintMatchesByGameOption {
	return func(o *FingerprintMatchesByGameRequest) {
		o.ctx = ctx
	}
}
