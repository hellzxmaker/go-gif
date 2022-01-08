package games

import "github.com/hellzxmaker/go-gif/players"

// Structs for API
type GameInfo struct {
	Round         int `json:"round"`
	ActivePlayers int `json:"activePlayers"`
}

type Question struct {
	Q    string `json:"q"`
	Type string `json:"type"`
}

type GameSession struct {
	GameJoinUid string           `json:"gameJoinUid"`
	GameJoinUrl string           `json:"gameJoinUrl"`
	Players     []players.Player `json:"players"`
	GameInfo    GameInfo         `json:"gameInfo"`
	Questions   []Question       `json:"questions"`
}

// Structs for Game Engine
