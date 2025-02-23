package categories

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/sjet47/go-curseforge/schema"
	"github.com/sjet47/go-curseforge/schema/enum"
)

type CategoriesOption func(*CategoriesRequest)

func NewCategoriesAPI(t http.RoundTripper) Categories {
	return func(gameID enum.GameID, o ...CategoriesOption) (*schema.GetCategoriesResponse, error) {
		r := new(CategoriesRequest)
		for _, f := range o {
			f(r)
		}
		r.GameID = gameID
		return schema.UnmarshalResponse[schema.GetCategoriesResponse](r.Do(r.ctx, t))
	}
}

type Categories func(gameID enum.GameID, o ...CategoriesOption) (*schema.GetCategoriesResponse, error)

// https://docs.curseforge.com/#get-categories
type CategoriesRequest struct {
	ctx context.Context

	GameID      enum.GameID
	ClassID     *enum.ClassID
	ClassesOnly *bool
}

func (r *CategoriesRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodGet
		params = make(map[string]string)
		path   = fmt.Sprintf("%s/v1/categories", schema.BaseUrl)
	)

	params["gameId"] = r.GameID.Param()
	if r.ClassID != nil {
		params["classId"] = r.ClassID.Param()
	}
	if r.ClassesOnly != nil {
		params["classesOnly"] = strconv.FormatBool(*r.ClassesOnly)
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

func (Categories) WithContext(ctx context.Context) CategoriesOption {
	return func(o *CategoriesRequest) {
		o.ctx = ctx
	}
}

func (Categories) WithClassID(id enum.ClassID) CategoriesOption {
	return func(o *CategoriesRequest) {
		o.ClassID = &id
	}
}

func (Categories) WithClassesOnly(classesOnly bool) CategoriesOption {
	return func(o *CategoriesRequest) {
		o.ClassesOnly = &classesOnly
	}
}
