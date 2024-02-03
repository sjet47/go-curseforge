package enum

// https://docs.curseforge.com/#tocS_HashAlgo
type HashAlgo int

const (
	HashAlgoSHA1 HashAlgo = 1
	HashAlgoMD5  HashAlgo = 2
)

func (ha HashAlgo) String() string {
	switch int(ha) {
	case int(HashAlgoSHA1):
		return "SHA1"
	case int(HashAlgoMD5):
		return "MD5"
	default:
		return "Unknown"
	}
}
