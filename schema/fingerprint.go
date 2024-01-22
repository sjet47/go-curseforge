package schema

// https://docs.curseforge.com/#tocS_FingerprintFuzzyMatch
type FingerprintFuzzyMatch struct {
	ID           FingerprintID `json:"id"`
	File         File          `json:"file"`
	LatestFiles  []File        `json:"latestFiles"`
	Fingerprints []Fingerprint `json:"fingerprints"`
}

// https://docs.curseforge.com/#tocS_FingerprintFuzzyMatchResult
type FingerprintFuzzyMatchResult struct {
	FuzzyMatches []FingerprintFuzzyMatch `json:"fuzzyMatches"`
}

// https://docs.curseforge.com/#tocS_FingerprintMatch
type FingerprintMatch struct {
	ID          FingerprintID `json:"id"`
	File        File          `json:"file"`
	LatestFiles []File        `json:"latestFiles"`
}

// https://docs.curseforge.com/#tocS_FingerprintsMatchesResult
type FingerprintsMatchesResult struct {
	IsCacheBuilt             bool                   `json:"isCacheBuilt"`
	ExactMatches             []FingerprintMatch     `json:"exactMatches"`
	ExactFingerprints        []Fingerprint          `json:"exactFingerprints"`
	PartialMatches           []FingerprintMatch     `json:"partialMatches"`
	PartialMatchFingerprints map[string]Fingerprint `json:"partialMatchFingerprints"`
	InstalledFingerprints    []Fingerprint          `json:"installedFingerprints"`
	UnmatchedFingerprints    []Fingerprint          `json:"unmatchedFingerprints"`
}

// https://docs.curseforge.com/#tocS_FolderFingerprint
type FolderFingerprint struct {
	FolderName   string        `json:"foldername"`
	Fingerprints []Fingerprint `json:"fingerprints"`
}
