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

type FingerprintFuzzyMatchesByGameOption func(*FingerprintFuzzyMatchesByGameRequest)

func NewFingerprintFuzzyMatchesByGameAPI(t http.RoundTripper) FingerprintFuzzyMatchesByGame {
	return func(gameID enum.GameID, o ...FingerprintFuzzyMatchesByGameOption) (*schema.GetFingerprintsFuzzyMatchesResponse, error) {
		r := new(FingerprintFuzzyMatchesByGameRequest)
		for _, f := range o {
			f(r)
		}
		r.GameID = gameID
		return schema.UnmarshalResponse[schema.GetFingerprintsFuzzyMatchesResponse](r.Do(r.ctx, t))
	}
}

type FingerprintFuzzyMatchesByGame func(gameID enum.GameID, o ...FingerprintFuzzyMatchesByGameOption) (*schema.GetFingerprintsFuzzyMatchesResponse, error)

// https://docs.curseforge.com/#get-fingerprints-fuzzy-matches-by-game-id
type FingerprintFuzzyMatchesByGameRequest struct {
	ctx context.Context

	schema.GetFuzzyMatchesRequestBody
}

func (r *FingerprintFuzzyMatchesByGameRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodPost
		path   = fmt.Sprintf("%s/v1/fingerprints/fuzzy/%d", schema.BaseUrl, r.GameID)
	)

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(&r.GetFuzzyMatchesRequestBody); err != nil {
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

func (FingerprintFuzzyMatchesByGame) WithContext(ctx context.Context) FingerprintFuzzyMatchesByGameOption {
	return func(o *FingerprintFuzzyMatchesByGameRequest) {
		o.ctx = ctx
	}
}

func (FingerprintFuzzyMatchesByGame) WithFolderFingerprint(ff ...schema.FolderFingerprint) FingerprintFuzzyMatchesByGameOption {
	return func(o *FingerprintFuzzyMatchesByGameRequest) {
		o.Fingerprints = ff
	}
}
