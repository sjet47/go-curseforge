package schema

import (
	"errors"
	"strconv"
	"strings"
)

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

// https://docs.curseforge.com/#tocS_SortOrder
type SortOrder string

const (
	SortOrderAscending  SortOrder = "asc"
	SortOrderDescending SortOrder = "desc"
)

func (so SortOrder) Param() string {
	return string(so)
}

// https://docs.curseforge.com/#tocS_ModLoaderType
type ModLoader int

const (
	ModLoaderAny        ModLoader = 0
	ModLoaderForge      ModLoader = 1
	ModLoaderCauldron   ModLoader = 2
	ModLoaderLiteLoader ModLoader = 3
	ModLoaderFabric     ModLoader = 4
	ModLoaderQuilt      ModLoader = 5
	ModLoaderNeoForge   ModLoader = 6
)

func (ml ModLoader) String() string {
	switch ml {
	case ModLoaderForge:
		return "Forge"
	case ModLoaderCauldron:
		return "Cauldron"
	case ModLoaderLiteLoader:
		return "LiteLoader"
	case ModLoaderFabric:
		return "Fabric"
	case ModLoaderQuilt:
		return "Quilt"
	case ModLoaderNeoForge:
		return "NeoForge"
	default:
		return "Unknown"
	}
}

func (ml ModLoader) Param() string {
	return strconv.Itoa(int(ml))
}

var (
	ErrUnknownModLoaderType = errors.New("unknown mod loader type")
	modLoaderStrMap         = map[string]ModLoader{
		"any":        ModLoaderAny,
		"forge":      ModLoaderForge,
		"cauldron":   ModLoaderCauldron,
		"liteloader": ModLoaderLiteLoader,
		"fabric":     ModLoaderFabric,
		"quilt":      ModLoaderQuilt,
		"neoforge":   ModLoaderNeoForge,
	}
)

func ParseModLoader(s string) (ModLoader, error) {
	ml, ok := modLoaderStrMap[strings.ToLower(s)]
	if !ok {
		return 0, ErrUnknownModLoaderType
	}
	return ml, nil
}
