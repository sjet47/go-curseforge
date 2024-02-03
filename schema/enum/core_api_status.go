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

func (cas CoreApiStatus) String() string {
	switch int(cas) {
	case int(CoreApiStatusPrivate):
		return "Private"
	case int(CoreApiStatusPublic):
		return "Public"
	default:
		return "Unknown"
	}
}
