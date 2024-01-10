package enum

import "strconv"

// https://docs.curseforge.com/#tocS_GameVersionType
type GameVersionType int

func (g GameVersionType) Param() string {
	return strconv.Itoa(int(g))
}
