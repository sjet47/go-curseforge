package enum

import "strconv"

type GameVersionType int32

func (g GameVersionType) Param() string {
	return strconv.Itoa(int(g))
}
