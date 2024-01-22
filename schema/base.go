package schema

import "strconv"

const (
	BaseUrl = "https://api.curseforge.com"
)

type (
	ModID                  int32
	ModAssetID             int32
	FileID                 int32
	AuthorID               int32
	Rating                 int
	GameVersionStr         string
	GameVersionID          int32
	GameVersionTypeID      int32
	MinecraftGameVersionID int32
	MinecraftModLoaderID   int32
	FingerprintID          int32
	Fingerprint            int
)

func (m ModID) Param() string {
	return strconv.Itoa(int(m))
}

func (f FileID) Param() string {
	return strconv.Itoa(int(f))
}

func (a AuthorID) Param() string {
	return strconv.Itoa(int(a))
}

func (r Rating) Param() string {
	return strconv.Itoa(int(r))
}

func (g GameVersionStr) Param() string {
	return string(g)
}

func (g GameVersionID) Param() string {
	return strconv.Itoa(int(g))
}

func (g GameVersionTypeID) Param() string {
	return strconv.Itoa(int(g))
}

func (m MinecraftGameVersionID) Param() string {
	return strconv.Itoa(int(m))
}

func (m MinecraftModLoaderID) Param() string {
	return strconv.Itoa(int(m))
}

func (f FingerprintID) Param() string {
	return strconv.Itoa(int(f))
}

func (f Fingerprint) Param() string {
	return strconv.Itoa(int(f))
}
