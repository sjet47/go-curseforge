package files

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sjet47/go-curseforge/schema"
)

type ModFileOption func(*ModFileRequest)

func NewModFileAPI(t http.RoundTripper) ModFile {
	return func(modID schema.ModID, fileID schema.FileID, o ...ModFileOption) (*schema.GetModFileResponse, error) {
		r := new(ModFileRequest)
		for _, f := range o {
			f(r)
		}
		r.ModID = modID
		r.FileID = fileID
		return schema.UnmarshalResponse[schema.GetModFileResponse](r.Do(r.ctx, t))
	}
}

type ModFile func(modID schema.ModID, fileID schema.FileID, o ...ModFileOption) (*schema.GetModFileResponse, error)

// https://docs.curseforge.com/#get-mod-file
type ModFileRequest struct {
	ctx context.Context

	ModID  schema.ModID
	FileID schema.FileID
}

func (r *ModFileRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodGet
		path   = fmt.Sprintf("%s/v1/mods/%d/files/%d", schema.BaseUrl, r.ModID, r.FileID)
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

func (ModFile) WithContext(ctx context.Context) ModFileOption {
	return func(o *ModFileRequest) {
		o.ctx = ctx
	}
}
