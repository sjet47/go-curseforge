package enum

import "strconv"

// https://docs.curseforge.com/#tocS_FileStatus
type FileStatus int

const (
	FileStatusProcessing         FileStatus = 1
	FileStatusChangesRequired    FileStatus = 2
	FileStatusUnderReview        FileStatus = 3
	FileStatusApproved           FileStatus = 4
	FileStatusRejected           FileStatus = 5
	FileStatusMalwareDetected    FileStatus = 6
	FileStatusDeleted            FileStatus = 7
	FileStatusArchived           FileStatus = 8
	FileStatusTesting            FileStatus = 9
	FileStatusReleased           FileStatus = 10
	FileStatusReadyForReview     FileStatus = 11
	FileStatusDeprecated         FileStatus = 12
	FileStatusBaking             FileStatus = 13
	FileStatusAwaitingPublishing FileStatus = 14
	FileStatusFailedPublishing   FileStatus = 15
)

func (fs FileStatus) Param() string {
	return strconv.Itoa(int(fs))
}

func (fs FileStatus) String() string {
	switch int(fs) {
	case int(FileStatusProcessing):
		return "Processing"
	case int(FileStatusChangesRequired):
		return "Changes Required"
	case int(FileStatusUnderReview):
		return "Under Review"
	case int(FileStatusApproved):
		return "Approved"
	case int(FileStatusRejected):
		return "Rejected"
	case int(FileStatusMalwareDetected):
		return "Malware Detected"
	case int(FileStatusDeleted):
		return "Deleted"
	case int(FileStatusArchived):
		return "Archived"
	case int(FileStatusTesting):
		return "Testing"
	case int(FileStatusReleased):
		return "Released"
	case int(FileStatusReadyForReview):
		return "Ready For Review"
	case int(FileStatusDeprecated):
		return "Deprecated"
	case int(FileStatusBaking):
		return "Baking"
	case int(FileStatusAwaitingPublishing):
		return "Awaiting Publishing"
	case int(FileStatusFailedPublishing):
		return "Failed Publishing"
	default:
		return "Unknown"
	}
}
