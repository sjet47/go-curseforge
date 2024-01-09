package enum

import "strconv"

// https://api.curseforge.com/v1/games
type GameID int32

const (
	MinecraftGameID GameID = 432
)

func (g GameID) Param() string {
	return strconv.Itoa(int(g))
}
