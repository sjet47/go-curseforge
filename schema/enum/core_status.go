package enum

import (
	"strconv"
)

// https://docs.curseforge.com/#tocS_CoreStatus
type CoreStatus int

const (
	CoreStatusDraft         CoreStatus = 1
	CoreStatusTest          CoreStatus = 2
	CoreStatusPendingReview CoreStatus = 3
	CoreStatusRejected      CoreStatus = 4
	CoreStatusApproved      CoreStatus = 5
	CoreStatusLive          CoreStatus = 6
)

func (cs CoreStatus) Param() string {
	return strconv.Itoa(int(cs))
}
