package schema

import (
	"time"

	"github.com/ASjet/go-curseforge/schema/enum"
)

// https://docs.curseforge.com/#tocS_Pagination
type Pagination struct {
	Index       int `json:"index"`
	PageSize    int `json:"pageSize"`
	ResultCount int `json:"resultCount"`
	TotalCount  int `json:"totalCount"`
}

// https://docs.curseforge.com/#tocS_Category
type Category struct {
	ID               enum.CategoryID `json:"id"`
	GameID           enum.GameID     `json:"gameId"`
	Name             string          `json:"name"`
	Slug             string          `json:"slug"`
	URL              string          `json:"url"`
	IconURL          string          `json:"iconUrl"`
	DateModified     time.Time       `json:"dateModified"`
	IsClass          bool            `json:"isClass"`
	ClassID          enum.ClassID    `json:"classId"`
	ParentCategoryID enum.CategoryID `json:"parentCategoryId"`
	DisplayIndex     int32           `json:"displayIndex"`
}
