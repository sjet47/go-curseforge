package enum

// https://docs.curseforge.com/#tocS_HashAlgo
type HashAlgo int

const (
	HashAlgoSHA1 HashAlgo = 1
	HashAlgoMD5  HashAlgo = 2
)
