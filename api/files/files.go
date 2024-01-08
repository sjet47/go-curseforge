package files

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ASjet/go-curseforge/schema"
)

func NewFilesAPI(t http.RoundTripper) Files {
	return func(modID schema.ModID, o ...func(*FilesRequest)) (*http.Response, error) {
		r := new(FilesRequest)
		for _, f := range o {
			f(r)
		}
		r.ModID = modID
		return r.Do(r.ctx, t)
	}
}

type Files func(modID schema.ModID, o ...func(*FilesRequest)) (*http.Response, error)

// https://docs.curseforge.com/#get-mod-files
type FilesRequest struct {
	ctx context.Context

	ModID             schema.ModID
	GameVersion       *string
	GameVersionTypeID *int
	ModLoader         *schema.ModLoader
	PageNum           int
	PageSize          int
}

func (r *FilesRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodGet
		params = make(map[string]string)
		path   = fmt.Sprintf("%s/v1/mods/%d/files", schema.BaseUrl, r.ModID)
	)

	if r.GameVersion != nil {
		params["gameVersion"] = *r.GameVersion
	}
	if r.GameVersionTypeID != nil {
		params["gameVersionTypeId"] = strconv.Itoa(*r.GameVersionTypeID)
	}
	if r.ModLoader != nil {
		params["modLoaderType"] = r.ModLoader.Param()
	}
	if r.PageNum > 0 {
		params["index"] = strconv.Itoa(r.PageNum)
	}
	if r.PageSize > 0 {
		params["pageSize"] = strconv.Itoa(r.PageSize)
	}

	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		return nil, err
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	if len(params) > 0 {
		query := req.URL.Query()
		for k, v := range params {
			query.Add(k, v)
		}
		req.URL.RawQuery = query.Encode()
	}

	return t.RoundTrip(req)
}

func (r Files) WithContext(ctx context.Context) func(*FilesRequest) {
	return func(o *FilesRequest) {
		o.ctx = ctx
	}
}

func (r Files) WithGameVersion(gameVersion string) func(*FilesRequest) {
	return func(o *FilesRequest) {
		o.GameVersion = &gameVersion
	}
}

func (r Files) WithGameVersionTypeID(gameVersionTypeID int) func(*FilesRequest) {
	return func(o *FilesRequest) {
		o.GameVersionTypeID = &gameVersionTypeID
	}
}

func (r Files) WithModLoader(modLoader schema.ModLoader) func(*FilesRequest) {
	return func(o *FilesRequest) {
		o.ModLoader = &modLoader
	}
}

func (r Files) WithPageNum(pageNum int) func(*FilesRequest) {
	return func(o *FilesRequest) {
		o.PageNum = pageNum
	}
}

func (r Files) WithPageSize(pageSize int) func(*FilesRequest) {
	return func(o *FilesRequest) {
		o.PageSize = pageSize
	}
}
