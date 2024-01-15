package enum

import "strconv"

// https://docs.curseforge.com/#tocS_GameVersionType
type GameVersionTypeID int

func (g GameVersionTypeID) Param() string {
	return strconv.Itoa(int(g))
}
