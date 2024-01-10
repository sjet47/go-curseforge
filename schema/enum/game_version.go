package enum

type GameVersion string

func (g GameVersion) Param() string {
	return string(g)
}

// TODO: add version enums
