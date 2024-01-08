package files

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ASjet/go-curseforge/schema"
)

func NewFileUrlAPI(t http.RoundTripper) FileUrl {
	return func(modID schema.ModID, fileID schema.FileID, o ...func(*FileUrlRequest)) (*http.Response, error) {
		r := new(FileUrlRequest)
		for _, f := range o {
			f(r)
		}
		r.ModID = modID
		r.FileID = fileID
		return r.Do(r.ctx, t)
	}
}

type FileUrl func(modID schema.ModID, fileID schema.FileID, o ...func(*FileUrlRequest)) (*http.Response, error)

// https://docs.curseforge.com/#get-mod-file-download-url
type FileUrlRequest struct {
	ctx context.Context

	ModID  schema.ModID
	FileID schema.FileID
}

func (r *FileUrlRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
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

func (r FileUrl) WithContext(ctx context.Context) func(*FileUrlRequest) {
	return func(o *FileUrlRequest) {
		o.ctx = ctx
	}
}
