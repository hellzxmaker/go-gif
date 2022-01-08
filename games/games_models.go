package games

import "github.com/hellzxmaker/go-gif/players"

type GameInfo struct {
	Round         int `json:"round"`
	ActivePlayers int `json:"activePlayers"`
}

type GameSession struct {
	GameJoinUid string           `json:"gameJoinUid"`
	GameJoinUrl string           `json:"gameJoinUrl"`
	Players     []players.Player `json:"players"`
	GameInfo    GameInfo         `json:"gameInfo"`
}
