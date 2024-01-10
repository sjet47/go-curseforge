package enum

import "strconv"

type GameVersionType int

func (g GameVersionType) Param() string {
	return strconv.Itoa(int(g))
}
