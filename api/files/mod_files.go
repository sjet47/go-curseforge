package files

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/sjet47/go-curseforge/schema"
	"github.com/sjet47/go-curseforge/schema/enum"
)

type ModFilesOption func(*ModFilesRequest)

func NewModFilesAPI(t http.RoundTripper) ModFiles {
	return func(modID schema.ModID, o ...ModFilesOption) (*schema.GetModFilesResponse, error) {
		r := new(ModFilesRequest)
		for _, f := range o {
			f(r)
		}
		r.ModID = modID
		return schema.UnmarshalResponse[schema.GetModFilesResponse](r.Do(r.ctx, t))
	}
}

type ModFiles func(modID schema.ModID, o ...ModFilesOption) (*schema.GetModFilesResponse, error)

// https://docs.curseforge.com/#get-mod-files
type ModFilesRequest struct {
	ctx context.Context

	ModID             schema.ModID
	GameVersion       *schema.GameVersionStr
	GameVersionTypeID *schema.GameVersionTypeID
	ModLoader         *enum.ModLoader
	Index             int // Not page number!
	PageSize          int
}

func (r *ModFilesRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodGet
		params = make(map[string]string)
		path   = fmt.Sprintf("%s/v1/mods/%d/files", schema.BaseUrl, r.ModID)
	)

	if r.GameVersion != nil {
		params["gameVersion"] = r.GameVersion.Param()
	}
	if r.GameVersionTypeID != nil {
		params["gameVersionTypeId"] = r.GameVersionTypeID.Param()
	}
	if r.ModLoader != nil {
		params["modLoaderType"] = r.ModLoader.Param()
	}
	if r.Index > 0 {
		params["index"] = strconv.Itoa(r.Index)
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

func (ModFiles) WithContext(ctx context.Context) ModFilesOption {
	return func(o *ModFilesRequest) {
		o.ctx = ctx
	}
}

func (ModFiles) WithGameVersion(gameVersion schema.GameVersionStr) ModFilesOption {
	return func(o *ModFilesRequest) {
		o.GameVersion = &gameVersion
	}
}

func (ModFiles) WithGameVersionTypeID(gameVersionTypeID schema.GameVersionTypeID) ModFilesOption {
	return func(o *ModFilesRequest) {
		o.GameVersionTypeID = &gameVersionTypeID
	}
}

func (ModFiles) WithModLoader(modLoader enum.ModLoader) ModFilesOption {
	return func(o *ModFilesRequest) {
		o.ModLoader = &modLoader
	}
}

func (ModFiles) WithIndex(index int) ModFilesOption {
	return func(o *ModFilesRequest) {
		o.Index = index
	}
}

func (ModFiles) WithPageSize(pageSize int) ModFilesOption {
	return func(o *ModFilesRequest) {
		o.PageSize = pageSize
	}
}
