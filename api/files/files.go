package files

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ASjet/go-curseforge/schema"
)

func NewFilesAPI(t http.RoundTripper) Files {
	return func(fileIDs []schema.FileID, o ...func(*FilesRequest)) (*http.Response, error) {
		r := new(FilesRequest)
		for _, f := range o {
			f(r)
		}
		r.FileIDs = fileIDs
		return r.Do(r.ctx, t)
	}
}

type Files func(fileIDs []schema.FileID, o ...func(*FilesRequest)) (*http.Response, error)

// https://docs.curseforge.com/#get-files
type FilesRequest struct {
	ctx context.Context

	FileIDs []schema.FileID
}

func (r *FilesRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodPost
		path   = fmt.Sprintf("%s/v1/mods/files", schema.BaseUrl)
	)

	body := &schema.GetModFilesRequestBody{
		FileIDs: r.FileIDs,
	}

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(body); err != nil {
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

func (Files) WithContext(ctx context.Context) func(*FilesRequest) {
	return func(o *FilesRequest) {
		o.ctx = ctx
	}
}
