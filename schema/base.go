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
	GameVersionID          int32
	MinecraftGameVersionID int32
	MinecraftModLoaderID   int32
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

func (g GameVersionID) Param() string {
	return strconv.Itoa(int(g))
}
