package mods

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/ASjet/go-curseforge/schema"
	"github.com/ASjet/go-curseforge/schema/enum"
)

func NewSearchModAPI(t http.RoundTripper) SearchMod {
	return func(gameID enum.GameID, o ...func(*SearchModRequest)) (*schema.SearchModsResponse, error) {
		r := new(SearchModRequest)
		for _, f := range o {
			f(r)
		}
		r.GameID = gameID
		return schema.UnmarshalResponse[schema.SearchModsResponse](r.Do(r.ctx, t))
	}
}

type SearchMod func(gameID enum.GameID, o ...func(*SearchModRequest)) (*schema.SearchModsResponse, error)

// https://docs.curseforge.com/#search-mods
type SearchModRequest struct {
	ctx context.Context

	GameID            enum.GameID
	ClassID           *enum.ClassID
	CategoryID        *enum.CategoryID
	CategoryIDs       *string
	GameVersion       *enum.GameVersion
	GameVersions      *string
	SearchFilter      *string
	SortField         *enum.ModsSearchSortField
	SortOrder         *enum.SortOrder
	ModLoaderType     *enum.ModLoader
	ModLoaderTypes    *string
	GameVersionTypeID *enum.GameVersionType
	AuthorID          *schema.AuthorID
	PrimaryAuthorID   *schema.AuthorID
	Slug              *string
	Index             int // Page number
	PageSize          int
}

func (r *SearchModRequest) Do(ctx context.Context, t http.RoundTripper) (*http.Response, error) {
	var (
		method = http.MethodGet
		params = make(map[string]string)
		path   = fmt.Sprintf("%s/v1/mods/search", schema.BaseUrl)
	)

	params["gameId"] = r.GameID.Param()

	if r.ClassID != nil {
		params["classId"] = r.ClassID.Param()
	}
	if r.CategoryID != nil {
		params["categoryId"] = r.CategoryID.Param()
	}
	if r.CategoryIDs != nil {
		params["categoryIds"] = *r.CategoryIDs
	}
	if r.GameVersion != nil {
		params["gameVersion"] = r.GameVersion.Param()
	}
	if r.GameVersions != nil {
		params["gameVersions"] = *r.GameVersions
	}
	if r.SearchFilter != nil {
		params["searchFilter"] = *r.SearchFilter
	}
	if r.SortField != nil {
		params["sortField"] = r.SortField.Param()
	}
	if r.SortOrder != nil {
		params["sortOrder"] = r.SortOrder.Param()
	}
	if r.ModLoaderType != nil {
		params["modLoaderType"] = r.ModLoaderType.Param()
	}
	if r.ModLoaderTypes != nil {
		params["modLoaderTypes"] = *r.ModLoaderTypes
	}
	if r.GameVersionTypeID != nil {
		params["gameVersionTypeId"] = r.GameVersionTypeID.Param()
	}
	if r.AuthorID != nil {
		params["authorId"] = r.AuthorID.Param()
	}
	if r.PrimaryAuthorID != nil {
		params["primaryAuthorId"] = r.PrimaryAuthorID.Param()
	}
	if r.Slug != nil {
		params["slug"] = *r.Slug
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

func (SearchMod) WithContext(ctx context.Context) func(*SearchModRequest) {
	return func(r *SearchModRequest) {
		r.ctx = ctx
	}
}

func (SearchMod) WithClassID(classID enum.ClassID) func(*SearchModRequest) {
	return func(r *SearchModRequest) {
		r.ClassID = &classID
	}
}

func (SearchMod) WithCategoryID(categoryID enum.CategoryID) func(*SearchModRequest) {
	return func(r *SearchModRequest) {
		r.CategoryID = &categoryID
	}
}

// NOTE: The maximum allowed category ids per query is 10
func (SearchMod) WithCategoryIDs(categoryIDs ...enum.CategoryID) func(*SearchModRequest) {
	if len(categoryIDs) > 10 {
		categoryIDs = categoryIDs[:10]
	}
	strIDs := make([]string, 0, len(categoryIDs))
	for _, id := range categoryIDs {
		strIDs = append(strIDs, strconv.Quote(id.Param()))
	}
	categoryIDsStr := "[" + strings.Join(strIDs, ",") + "]"
	return func(r *SearchModRequest) {
		r.CategoryIDs = &categoryIDsStr
	}
}

func (SearchMod) WithGameVersion(gameVersion enum.GameVersion) func(*SearchModRequest) {
	return func(r *SearchModRequest) {
		r.GameVersion = &gameVersion
	}
}

// NOTE: The maximum allowed game versions per query is 4
func (SearchMod) WithGameVersions(gameVersion ...string) func(*SearchModRequest) {
	if len(gameVersion) > 4 {
		gameVersion = gameVersion[:4]
	}
	for i, v := range gameVersion {
		gameVersion[i] = strconv.Quote(v)
	}
	gameVersionsStr := "[" + strings.Join(gameVersion, ",") + "]"
	return func(r *SearchModRequest) {
		r.GameVersions = &gameVersionsStr
	}
}

func (SearchMod) WithSearchFilter(searchFilter string) func(*SearchModRequest) {
	return func(r *SearchModRequest) {
		r.SearchFilter = &searchFilter
	}
}

func (SearchMod) WithSortField(sortField enum.ModsSearchSortField) func(*SearchModRequest) {
	return func(r *SearchModRequest) {
		r.SortField = &sortField
	}
}

func (SearchMod) WithSortOrder(sortOrder enum.SortOrder) func(*SearchModRequest) {
	return func(r *SearchModRequest) {
		r.SortOrder = &sortOrder
	}
}

func (SearchMod) WithModLoaderType(modLoaderType enum.ModLoader) func(*SearchModRequest) {
	return func(r *SearchModRequest) {
		r.ModLoaderType = &modLoaderType
	}
}

// NOTE: The maximum allowed mod loader types per query is 5
func (SearchMod) WithModLoaderTypes(modLoaderType ...enum.ModLoader) func(*SearchModRequest) {
	if len(modLoaderType) > 5 {
		modLoaderType = modLoaderType[:5]
	}
	strModLoaders := make([]string, 0, len(modLoaderType))
	for _, t := range modLoaderType {
		strModLoaders = append(strModLoaders, t.String())
	}
	modLoaderTypesStr := "[" + strings.Join(strModLoaders, ",") + "]"
	return func(r *SearchModRequest) {
		r.ModLoaderTypes = &modLoaderTypesStr
	}
}

func (SearchMod) WithGameVersionTypeID(gameVersionTypeID enum.GameVersionType) func(*SearchModRequest) {
	return func(r *SearchModRequest) {
		r.GameVersionTypeID = &gameVersionTypeID
	}
}

func (SearchMod) WithAuthorID(authorID schema.AuthorID) func(*SearchModRequest) {
	return func(r *SearchModRequest) {
		r.AuthorID = &authorID
	}
}

func (SearchMod) WithPrimaryAuthorID(primaryAuthorID schema.AuthorID) func(*SearchModRequest) {
	return func(r *SearchModRequest) {
		r.PrimaryAuthorID = &primaryAuthorID
	}
}

func (SearchMod) WithSlug(slug string) func(*SearchModRequest) {
	return func(r *SearchModRequest) {
		r.Slug = &slug
	}
}

func (SearchMod) WithIndex(index int) func(*SearchModRequest) {
	return func(r *SearchModRequest) {
		r.Index = index
	}
}

func (SearchMod) WithPageSize(pageSize int) func(*SearchModRequest) {
	return func(r *SearchModRequest) {
		r.PageSize = pageSize
	}
}
