package enum

import "strconv"

// https://docs.curseforge.com/#tocS_CoreApiStatus
type CoreApiStatus int

const (
	CoreApiStatusPrivate CoreApiStatus = 1
	CoreApiStatusPublic  CoreApiStatus = 2
)

func (cas CoreApiStatus) Param() string {
	return strconv.Itoa(int(cas))
}
