package enum

import "strconv"

// https://docs.curseforge.com/#tocS_FileReleaseType
type FileReleaseType int

const (
	FileReleaseTypeRelease FileReleaseType = 1
	FileReleaseTypeBeta    FileReleaseType = 2
	FileReleaseTypeAlpha   FileReleaseType = 3
)

func (f FileReleaseType) Param() string {
	return strconv.Itoa(int(f))
}

func (f FileReleaseType) String() string {
	switch int(f) {
	case int(FileReleaseTypeRelease):
		return "Release"
	case int(FileReleaseTypeBeta):
		return "Beta"
	case int(FileReleaseTypeAlpha):
		return "Alpha"
	default:
		return "Unknown"
	}
}
