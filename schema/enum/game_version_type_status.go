package enum

type GameVersionTypeStatus int

const (
	GameVersionTypeStatusNormal  GameVersionTypeStatus = 1
	GameVersionTypeStatusDeleted GameVersionTypeStatus = 2
)

func (g GameVersionTypeStatus) String() string {
	switch int(g) {
	case int(GameVersionTypeStatusNormal):
		return "Normal"
	case int(GameVersionTypeStatusDeleted):
		return "Deleted"
	default:
		return "Unknown"
	}
}
