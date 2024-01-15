package schema

import (
	"time"

	"github.com/ASjet/go-curseforge/schema/enum"
)

// https://docs.curseforge.com/#tocS_File
type File struct {
	ID                   FileID                `json:"id"`
	GameID               enum.GameID           `json:"gameId"`
	ModID                ModID                 `json:"modId"`
	IsAvailable          bool                  `json:"isAvailable"`
	DisplayName          string                `json:"displayName"`
	FileName             string                `json:"fileName"`
	ReleaseType          enum.FileReleaseType  `json:"releaseType"`
	FileStatus           enum.FileStatus       `json:"fileStatus"`
	Hashes               []FileHash            `json:"hashes"`
	FileDate             time.Time             `json:"fileDate"`
	FileLength           int64                 `json:"fileLength"`
	DownloadCount        int64                 `json:"downloadCount"`
	FileSizeOnDisk       int64                 `json:"fileSizeOnDisk"`
	DownloadURL          string                `json:"downloadUrl"`
	GameVersions         []enum.GameVersion    `json:"gameVersions"`
	SortableGameVersions []SortableGameVersion `json:"sortableGameVersions"`
	Dependencies         []FileDependency      `json:"dependencies"`
	ExposeAsAlternative  bool                  `json:"exposeAsAlternative"`
	ParentProjectFileID  FileID                `json:"parentProjectFileId"`
	AlternateFileID      FileID                `json:"alternateFileId"`
	IsServerPack         bool                  `json:"isServerPack"`
	ServerPackFileID     FileID                `json:"serverPackFileId"`
	IsEarlyAccessContent bool                  `json:"isEarlyAccessContent"`
	EarlyAccessEndDate   time.Time             `json:"earlyAccessEndDate"`
	FileFingerprint      int64                 `json:"fileFingerprint"`
	Modules              []FileModule          `json:"modules"`
}

// https://docs.curseforge.com/#tocS_FileDependency
type FileDependency struct {
	ModID        ModID                 `json:"modId"`
	RelationType enum.FileRelationType `json:"relationType"`
}

// https://docs.curseforge.com/#tocS_FileHash
type FileHash struct {
	Value string        `json:"value"`
	Algo  enum.HashAlgo `json:"algo"`
}

// https://docs.curseforge.com/#tocS_FileIndex
type FileIndex struct {
	GameVersion       enum.GameVersion       `json:"gameVersion"`
	FileID            FileID                 `json:"fileId"`
	FileName          string                 `json:"fileName"`
	ReleaseType       enum.FileReleaseType   `json:"releaseType"`
	GameVersionTypeID enum.GameVersionTypeID `json:"gameVersionTypeId"`
	ModLoader         enum.ModLoader         `json:"modLoader"`
}

// https://docs.curseforge.com/#tocS_FileModule
type FileModule struct {
	Name        string `json:"name"`
	Fingerprint int64  `json:"fingerprint"`
}
