package files

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sjet47/go-curseforge/schema"
)

type ModFileChangelogOption func(*ModFileChangelogRequest)

func NewModFileChangelogAPI(t http.RoundTripper) ModFileChangelog {
	return func(modID schema.ModID, fileID schema.FileID, o ...ModFileChangelogOption) (*schema.StringResponse, error) {
		r := new(ModFileChangelogRequest)
		for _, f := range o {
			f(r)
		}
		r.ModID = modID
		r.FileID = fileID
		return schema.UnmarshalResponse[schema.StringResponse](r.Do(r.ctx, t))
	}
}

type ModFileChangelog func(modID schema.ModID, fileID schema.FileID, o ...ModFileChangelogOption) (*schema.StringResponse, error)

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

func (ModFileChangelog) WithContext(ctx context.Context) ModFileChangelogOption {
	return func(o *ModFileChangelogRequest) {
		o.ctx = ctx
	}
}
