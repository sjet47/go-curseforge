package files

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ASjet/go-curseforge/schema"
)

func NewFileAPI(t http.RoundTripper) File {
	return func(modID schema.ModID, fileID schema.FileID, o ...func(*FileRequest)) (*http.Response, error) {
		r := new(FileRequest)
		for _, f := range o {
			f(r)
		}
		r.ModID = modID
		r.FileID = fileID
		return r.Do(r.ctx, t)
	}
}

type File func(modID schema.ModID, fileID schema.FileID, o ...func(*FileRequest)) (*http.Response, error)

// https://docs.curseforge.com/#get-mod-file
type FileRequest struct {
	ctx context.Context

	ModID  schema.ModID
	FileID schema.FileID
}

func (r *FileRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
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

func (r File) WithContext(ctx context.Context) func(*FileRequest) {
	return func(o *FileRequest) {
		o.ctx = ctx
	}
}
