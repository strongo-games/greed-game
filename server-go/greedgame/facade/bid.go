package facade

import (
	"github.com/strongo-games/greed-game/server-go/greedgame/models"
	"github.com/strongo-games/arena/arena-go"
)

type BidOutput struct {
	RivalKey            arena.BattleID
	Game                models.Game
	User                models.User
	RivalUser           models.User
	UserContestant      arena.Contestant
	RivalUserContestant arena.Contestant
}
