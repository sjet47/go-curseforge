package enum

import "strconv"

// https://docs.curseforge.com/#tocS_ModsSearchSortField
type ModsSearchSortField int

const (
	ModsSearchSortFieldFeatured         ModsSearchSortField = 1
	ModsSearchSortFieldPopularity       ModsSearchSortField = 2
	ModsSearchSortFieldLastUpdated      ModsSearchSortField = 3
	ModsSearchSortFieldName             ModsSearchSortField = 4
	ModsSearchSortFieldAuthor           ModsSearchSortField = 5
	ModsSearchSortFieldTotalDownloads   ModsSearchSortField = 6
	ModsSearchSortFieldCategory         ModsSearchSortField = 7
	ModsSearchSortFieldGameVersion      ModsSearchSortField = 8
	ModsSearchSortFieldEarlyAccess      ModsSearchSortField = 9
	ModsSearchSortFieldFeaturedReleased ModsSearchSortField = 10
	ModsSearchSortFieldReleasedDate     ModsSearchSortField = 11
	ModsSearchSortFieldRating           ModsSearchSortField = 12
)

func (ms ModsSearchSortField) Param() string {
	return strconv.Itoa(int(ms))
}
