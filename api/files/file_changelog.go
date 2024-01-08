package files

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ASjet/go-curseforge/schema"
)

func NewFileChangelogAPI(t http.RoundTripper) FileChangelog {
	return func(modID schema.ModID, fileID schema.FileID, o ...func(*FileChangelogRequest)) (*http.Response, error) {
		r := new(FileChangelogRequest)
		for _, f := range o {
			f(r)
		}
		r.ModID = modID
		r.FileID = fileID
		return r.Do(r.ctx, t)
	}
}

type FileChangelog func(modID schema.ModID, fileID schema.FileID, o ...func(*FileChangelogRequest)) (*http.Response, error)

// https://docs.curseforge.com/#get-mod-file-changelog
type FileChangelogRequest struct {
	ctx context.Context

	ModID  schema.ModID
	FileID schema.FileID
}

func (r *FileChangelogRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodGet
		path   = fmt.Sprintf("%s/v1/mods/%d/files/%d/changelog", schema.BaseUrl, r.ModID, r.FileID)
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

func (r FileChangelog) WithContext(ctx context.Context) func(*FileChangelogRequest) {
	return func(o *FileChangelogRequest) {
		o.ctx = ctx
	}
}
