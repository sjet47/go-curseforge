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

type FingerprintFuzzyMatchesOption func(*FingerprintFuzzyMatchesRequest)

func NewFingerprintFuzzyMatchesAPI(t http.RoundTripper) FingerprintFuzzyMatches {
	return func(gameID enum.GameID, o ...FingerprintFuzzyMatchesOption) (*schema.GetFingerprintsFuzzyMatchesResponse, error) {
		r := new(FingerprintFuzzyMatchesRequest)
		for _, f := range o {
			f(r)
		}
		r.GameID = gameID
		return schema.UnmarshalResponse[schema.GetFingerprintsFuzzyMatchesResponse](r.Do(r.ctx, t))
	}
}

type FingerprintFuzzyMatches func(gameID enum.GameID, o ...FingerprintFuzzyMatchesOption) (*schema.GetFingerprintsFuzzyMatchesResponse, error)

// https://docs.curseforge.com/#get-fingerprints-fuzzy-matches
type FingerprintFuzzyMatchesRequest struct {
	ctx context.Context

	schema.GetFuzzyMatchesRequestBody
}

func (r *FingerprintFuzzyMatchesRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodPost
		path   = fmt.Sprintf("%s/v1/fingerprints/fuzzy", schema.BaseUrl)
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

func (FingerprintFuzzyMatches) WithContext(ctx context.Context) FingerprintFuzzyMatchesOption {
	return func(o *FingerprintFuzzyMatchesRequest) {
		o.ctx = ctx
	}
}

func (FingerprintFuzzyMatches) WithFolderFingerprint(ff ...schema.FolderFingerprint) FingerprintFuzzyMatchesOption {
	return func(o *FingerprintFuzzyMatchesRequest) {
		o.Fingerprints = ff
	}
}
