package enum

// https://docs.curseforge.com/#tocS_ModStatus
type ModStatus int

const (
	ModStatusNew             ModStatus = 1
	ModStatusChangesRequired ModStatus = 2
	ModStatusUnderSoftReview ModStatus = 3
	ModStatusApproved        ModStatus = 4
	ModStatusRejected        ModStatus = 5
	ModStatusChangesMade     ModStatus = 6
	ModStatusInactive        ModStatus = 7
	ModStatusAbandoned       ModStatus = 8
	ModStatusDeleted         ModStatus = 9
	ModStatusUnderReview     ModStatus = 10
)

func (ms ModStatus) String() string {
	switch int(ms) {
	case int(ModStatusNew):
		return "New"
	case int(ModStatusChangesRequired):
		return "Changes Required"
	case int(ModStatusUnderSoftReview):
		return "Under Soft Review"
	case int(ModStatusApproved):
		return "Approved"
	case int(ModStatusRejected):
		return "Rejected"
	case int(ModStatusChangesMade):
		return "Changes Made"
	case int(ModStatusInactive):
		return "Inactive"
	case int(ModStatusAbandoned):
		return "Abandoned"
	case int(ModStatusDeleted):
		return "Deleted"
	case int(ModStatusUnderReview):
		return "Under Review"
	default:
		return "Unknown"
	}
}
