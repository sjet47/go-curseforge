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

func (cs CoreStatus) String() string {
	switch int(cs) {
	case int(CoreStatusDraft):
		return "Draft"
	case int(CoreStatusTest):
		return "Test"
	case int(CoreStatusPendingReview):
		return "Pending Review"
	case int(CoreStatusRejected):
		return "Rejected"
	case int(CoreStatusApproved):
		return "Approved"
	case int(CoreStatusLive):
		return "Live"
	default:
		return "Unknown"
	}
}
