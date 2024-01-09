package schema

import "strconv"

const (
	BaseUrl = "https://api.curseforge.com"
)

type (
	ModID           int32
	FileID          int32
	GameID          int32
	ClassID         int32
	CategoryID      int32
	GameVersionType int32
	AuthorID        int32
)

func (m ModID) Param() string {
	return strconv.Itoa(int(m))
}

func (f FileID) Param() string {
	return strconv.Itoa(int(f))
}

func (g GameID) Param() string {
	return strconv.Itoa(int(g))
}

func (c ClassID) Param() string {
	return strconv.Itoa(int(c))
}

func (c CategoryID) Param() string {
	return strconv.Itoa(int(c))
}

func (g GameVersionType) Param() string {
	return strconv.Itoa(int(g))
}

func (a AuthorID) Param() string {
	return strconv.Itoa(int(a))
}
