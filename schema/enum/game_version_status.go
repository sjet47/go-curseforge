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
