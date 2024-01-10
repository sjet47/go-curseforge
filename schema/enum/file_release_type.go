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
