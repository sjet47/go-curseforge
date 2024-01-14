package schema

import (
	"time"

	"github.com/ASjet/go-curseforge/schema/enum"
)

// https://docs.curseforge.com/#tocS_MinecraftGameVersion
type MinecraftGameVersion struct {
	ID                    MinecraftGameVersionID     `json:"id"`
	GameVersionID         GameVersionID              `json:"gameVersionId"`
	VersionString         enum.GameVersion           `json:"versionString"`
	JarDownloadURL        string                     `json:"jarDownloadUrl"`
	JsonDownloadURL       string                     `json:"jsonDownloadUrl"`
	Approved              bool                       `json:"approved"`
	DateModified          time.Time                  `json:"dateModified"`
	GameVersionTypeID     enum.GameVersionType       `json:"gameVersionTypeId"`
	GameVersionStatus     enum.GameVersionStatus     `json:"gameVersionStatus"`
	GameVersionTypeStatus enum.GameVersionTypeStatus `json:"gameVersionTypeStatus"`
}

// https://docs.curseforge.com/#tocS_MinecraftModLoaderIndex
type MinecraftModLoaderIndex struct {
	Name         string           `json:"name"`
	GameVersion  enum.GameVersion `json:"gameVersion"`
	Latest       bool             `json:"latest"`
	Recommended  bool             `json:"recommended"`
	DateModified time.Time        `json:"dateModified"`
	Type         enum.ModLoader   `json:"type"`
}

// https://docs.curseforge.com/#tocS_MinecraftModLoaderVersion
type MinecraftModLoaderVersion struct {
	ID                             MinecraftModLoaderID        `json:"id"`
	GameVersionID                  GameVersionID               `json:"gameVersionId"`
	MinecraftGameVersionID         MinecraftGameVersionID      `json:"minecraftGameVersionId"`
	ForgeVersion                   string                      `json:"forgeVersion"`
	Name                           string                      `json:"name"`
	Type                           enum.ModLoader              `json:"type"`
	DownloadURL                    string                      `json:"downloadUrl"`
	FileName                       string                      `json:"filename"`
	InstallMethod                  enum.ModLoaderInstallMethod `json:"installMethod"`
	Latest                         bool                        `json:"latest"`
	Recommended                    bool                        `json:"recommended"`
	Approved                       bool                        `json:"approved"`
	DateModified                   time.Time                   `json:"dateModified"`
	MavenVersionString             string                      `json:"mavenVersionString"`
	VersionJson                    string                      `json:"versionJson"`
	LibrariesInstallLocation       string                      `json:"librariesInstallLocation"`
	MinecraftVersion               enum.GameVersion            `json:"minecraftVersion"`
	AdditionalFilesJson            string                      `json:"additionalFilesJson"`
	ModLoaderGameVersionID         GameVersionID               `json:"modLoaderGameVersionId"`
	ModLoaderGameVersionTypeID     enum.GameVersionType        `json:"modLoaderGameVersionTypeId"`
	ModLoaderGameVersionStatus     enum.GameVersionStatus      `json:"modLoaderGameVersionStatus"`
	ModLoaderGameVersionTypeStatus enum.GameVersionTypeStatus  `json:"modLoaderGameVersionTypeStatus"`
	McGameVersionID                MinecraftGameVersionID      `json:"mcGameVersionId"`
	McGameVersionTypeID            enum.GameVersionType        `json:"mcGameVersionTypeId"`
	McGameVersionStatus            enum.GameVersionStatus      `json:"mcGameVersionStatus"`
	McGameVersionTypeStatus        enum.GameVersionTypeStatus  `json:"mcGameVersionTypeStatus"`
	InstallProfileJson             string                      `json:"installProfileJson"`
}
