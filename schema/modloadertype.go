package schema

import (
	"errors"
	"strconv"
	"strings"
)

// https://docs.curseforge.com/#tocS_ModLoaderType
type ModLoader int

const (
	ModLoaderAny        = 0
	ModLoaderForge      = 1
	ModLoaderCauldron   = 2
	ModLoaderLiteLoader = 3
	ModLoaderFabric     = 4
	ModLoaderQuilt      = 5
	ModLoaderNeoForge   = 6
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
