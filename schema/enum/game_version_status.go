package enum

import "strconv"

// https://docs.curseforge.com/#tocS_GameVersionStatus
type GameVersionStatus int

const (
	GameVersionStatusApproved GameVersionStatus = 1
	GameVersionStatusDeleted  GameVersionStatus = 2
	GameVersionStatusNew      GameVersionStatus = 3
)

func (g GameVersionStatus) Param() string {
	return strconv.Itoa(int(g))
}

func (g GameVersionStatus) String() string {
	switch int(g) {
	case int(GameVersionStatusApproved):
		return "Approved"
	case int(GameVersionStatusDeleted):
		return "Deleted"
	case int(GameVersionStatusNew):
		return "New"
	default:
		return "Unknown"
	}
}
