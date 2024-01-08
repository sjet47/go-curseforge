package files

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ASjet/go-curseforge/schema"
)

func NewModFilesAPI(t http.RoundTripper) ModFiles {
	return func(modID schema.ModID, o ...func(*ModFilesRequest)) (*http.Response, error) {
		r := new(ModFilesRequest)
		for _, f := range o {
			f(r)
		}
		r.ModID = modID
		return r.Do(r.ctx, t)
	}
}

type ModFiles func(modID schema.ModID, o ...func(*ModFilesRequest)) (*http.Response, error)

// https://docs.curseforge.com/#get-mod-files
type ModFilesRequest struct {
	ctx context.Context

	ModID             schema.ModID
	GameVersion       *string
	GameVersionTypeID *int
	ModLoader         *schema.ModLoader
	PageNum           int
	PageSize          int
}

func (r *ModFilesRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
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

func (r ModFiles) WithContext(ctx context.Context) func(*ModFilesRequest) {
	return func(o *ModFilesRequest) {
		o.ctx = ctx
	}
}

func (r ModFiles) WithGameVersion(gameVersion string) func(*ModFilesRequest) {
	return func(o *ModFilesRequest) {
		o.GameVersion = &gameVersion
	}
}

func (r ModFiles) WithGameVersionTypeID(gameVersionTypeID int) func(*ModFilesRequest) {
	return func(o *ModFilesRequest) {
		o.GameVersionTypeID = &gameVersionTypeID
	}
}

func (r ModFiles) WithModLoader(modLoader schema.ModLoader) func(*ModFilesRequest) {
	return func(o *ModFilesRequest) {
		o.ModLoader = &modLoader
	}
}

func (r ModFiles) WithPageNum(pageNum int) func(*ModFilesRequest) {
	return func(o *ModFilesRequest) {
		o.PageNum = pageNum
	}
}

func (r ModFiles) WithPageSize(pageSize int) func(*ModFilesRequest) {
	return func(o *ModFilesRequest) {
		o.PageSize = pageSize
	}
}
