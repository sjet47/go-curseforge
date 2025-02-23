package schema

import (
	"time"

	"github.com/sjet47/go-curseforge/schema/enum"
)

// https://docs.curseforge.com/#tocS_Mod
type Mod struct {
	ID                            ModID           `json:"id"`
	GameID                        enum.GameID     `json:"gameId"`
	Name                          string          `json:"name"`
	Slug                          string          `json:"slug"`
	Links                         ModLinks        `json:"links"`
	Summary                       string          `json:"summary"`
	Status                        enum.ModStatus  `json:"status"`
	DownloadCount                 int64           `json:"downloadCount"`
	IsFeatured                    bool            `json:"isFeatured"`
	PrimaryCategoryID             enum.CategoryID `json:"primaryCategoryId"`
	Categories                    []Category      `json:"categories"`
	ClassID                       enum.ClassID    `json:"classId"`
	Authors                       []ModAuthor     `json:"authors"`
	Logo                          ModAsset        `json:"logo"`
	Screenshots                   []ModAsset      `json:"screenshots"`
	MainFileID                    FileID          `json:"mainFileId"`
	LatestFiles                   []File          `json:"latestFiles"`
	LatestFilesIndexes            []FileIndex     `json:"latestFilesIndexes"`
	LatestEarlyAccessFilesIndexes []FileIndex     `json:"latestEarlyAccessFilesIndexes"`
	DateCreated                   time.Time       `json:"dateCreated"`
	DateModified                  time.Time       `json:"dateModified"`
	DateReleased                  time.Time       `json:"dateReleased"`
	AllowModDistribution          bool            `json:"allowModDistribution"`
	GamePopularityRank            int32           `json:"gamePopularityRank"`
	IsAvailable                   bool            `json:"isAvailable"`
	ThumbsUpCount                 int32           `json:"thumbsUpCount"`
	Rating                        Rating          `json:"rating"`
}

// https://docs.curseforge.com/#tocS_ModAsset
type ModAsset struct {
	ID           ModAssetID `json:"id"`
	ModID        ModID      `json:"modId"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	ThumbnailUrl string     `json:"thumbnailUrl"`
	URL          string     `json:"url"`
}

// https://docs.curseforge.com/#tocS_ModAuthor
type ModAuthor struct {
	ID   AuthorID `json:"id"`
	Name string   `json:"name"`
	URL  string   `json:"url"`
}

// https://docs.curseforge.com/#tocS_ModLinks
type ModLinks struct {
	WebsiteURL string `json:"websiteUrl"`
	WikiURL    string `json:"wikiUrl"`
	IssueURL   string `json:"issueUrl"`
	SourceURL  string `json:"sourceUrl"`
}
