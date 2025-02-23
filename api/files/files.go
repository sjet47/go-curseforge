package files

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sjet47/go-curseforge/schema"
)

type FilesOption func(*FilesRequest)

func NewFilesAPI(t http.RoundTripper) Files {
	return func(fileIDs []schema.FileID, o ...FilesOption) (*schema.GetFilesResponse, error) {
		r := new(FilesRequest)
		for _, f := range o {
			f(r)
		}
		r.FileIDs = fileIDs
		return schema.UnmarshalResponse[schema.GetFilesResponse](r.Do(r.ctx, t))
	}
}

type Files func(fileIDs []schema.FileID, o ...FilesOption) (*schema.GetFilesResponse, error)

// https://docs.curseforge.com/#get-files
type FilesRequest struct {
	ctx context.Context

	schema.GetModFilesRequestBody
}

func (r *FilesRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodPost
		path   = fmt.Sprintf("%s/v1/mods/files", schema.BaseUrl)
	)

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(&r.GetModFilesRequestBody); err != nil {
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

func (Files) WithContext(ctx context.Context) FilesOption {
	return func(o *FilesRequest) {
		o.ctx = ctx
	}
}
