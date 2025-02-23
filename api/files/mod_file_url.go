package files

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sjet47/go-curseforge/schema"
)

type ModFileUrlOption func(*ModFileUrlRequest)

func NewModFileUrlAPI(t http.RoundTripper) ModFileUrl {
	return func(modID schema.ModID, fileID schema.FileID, o ...ModFileUrlOption) (*schema.StringResponse, error) {
		r := new(ModFileUrlRequest)
		for _, f := range o {
			f(r)
		}
		r.ModID = modID
		r.FileID = fileID
		return schema.UnmarshalResponse[schema.StringResponse](r.Do(r.ctx, t))
	}
}

type ModFileUrl func(modID schema.ModID, fileID schema.FileID, o ...ModFileUrlOption) (*schema.StringResponse, error)

// https://docs.curseforge.com/#get-mod-file-download-url
type ModFileUrlRequest struct {
	ctx context.Context

	ModID  schema.ModID
	FileID schema.FileID
}

func (r *ModFileUrlRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodGet
		path   = fmt.Sprintf("%s/v1/mods/%d/files/%d/download-url", schema.BaseUrl, r.ModID, r.FileID)
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

func (ModFileUrl) WithContext(ctx context.Context) ModFileUrlOption {
	return func(o *ModFileUrlRequest) {
		o.ctx = ctx
	}
}
