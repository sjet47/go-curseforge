package files

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ASjet/go-curseforge/schema"
)

func NewModFileChangelogAPI(t http.RoundTripper) ModFileChangelog {
	return func(modID schema.ModID, fileID schema.FileID, o ...func(*ModFileChangelogRequest)) (*http.Response, error) {
		r := new(ModFileChangelogRequest)
		for _, f := range o {
			f(r)
		}
		r.ModID = modID
		r.FileID = fileID
		return r.Do(r.ctx, t)
	}
}

type ModFileChangelog func(modID schema.ModID, fileID schema.FileID, o ...func(*ModFileChangelogRequest)) (*http.Response, error)

// https://docs.curseforge.com/#get-mod-file-changelog
type ModFileChangelogRequest struct {
	ctx context.Context

	ModID  schema.ModID
	FileID schema.FileID
}

func (r *ModFileChangelogRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
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

func (r ModFileChangelog) WithContext(ctx context.Context) func(*ModFileChangelogRequest) {
	return func(o *ModFileChangelogRequest) {
		o.ctx = ctx
	}
}
