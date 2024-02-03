package enum

import "strconv"

// https://docs.curseforge.com/#tocS_FileRelationType
type FileRelationType int

const (
	EmbeddedLibrary    FileRelationType = 1
	OptionalDependency FileRelationType = 2
	RequiredDependency FileRelationType = 3
	Tool               FileRelationType = 4
	Incompatible       FileRelationType = 5
	Include            FileRelationType = 6
)

func (fr FileRelationType) Param() string {
	return strconv.Itoa(int(fr))
}

func (fr FileRelationType) String() string {
	switch int(fr) {
	case int(EmbeddedLibrary):
		return "Embedded Library"
	case int(OptionalDependency):
		return "Optional Dependency"
	case int(RequiredDependency):
		return "Required Dependency"
	case int(Tool):
		return "Tool"
	case int(Incompatible):
		return "Incompatible"
	case int(Include):
		return "Include"
	default:
		return "Unknown"
	}
}
