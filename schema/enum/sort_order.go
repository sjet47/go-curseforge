package enum

// https://docs.curseforge.com/#tocS_SortOrder
type SortOrder string

const (
	SortOrderAscending  SortOrder = "asc"
	SortOrderDescending SortOrder = "desc"
)

func (so SortOrder) Param() string {
	return string(so)
}
