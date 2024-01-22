package fingerprints

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ASjet/go-curseforge/schema"
)

type FingerprintMatchesOption func(*FingerprintMatchesRequest)

func NewFingerprintMatchesAPI(t http.RoundTripper) FingerprintMatches {
	return func(fingerprints []schema.Fingerprint,
		o ...FingerprintMatchesOption) (*schema.GetFingerprintMatchesResponse, error) {
		r := new(FingerprintMatchesRequest)
		for _, f := range o {
			f(r)
		}
		r.Fingerprints = fingerprints
		return schema.UnmarshalResponse[schema.GetFingerprintMatchesResponse](r.Do(r.ctx, t))
	}
}

type FingerprintMatches func(fingerprints []schema.Fingerprint, o ...FingerprintMatchesOption) (*schema.GetFingerprintMatchesResponse, error)

// https://docs.curseforge.com/#get-fingerprints-matches
type FingerprintMatchesRequest struct {
	ctx context.Context

	schema.GetFingerprintMatchesRequestBody
}

func (r *FingerprintMatchesRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodPost
		path   = fmt.Sprintf("%s/v1/fingerprints", schema.BaseUrl)
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

func (FingerprintMatches) WithContext(ctx context.Context) FingerprintMatchesOption {
	return func(o *FingerprintMatchesRequest) {
		o.ctx = ctx
	}
}
