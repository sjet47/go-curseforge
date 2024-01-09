package enum

import "strconv"

type ClassID int32

func (c ClassID) Param() string {
	return strconv.Itoa(int(c))
}
