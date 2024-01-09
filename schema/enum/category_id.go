package enum

import "strconv"

type CategoryID int32

func (c CategoryID) Param() string {
	return strconv.Itoa(int(c))
}
