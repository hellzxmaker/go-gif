package godata

type GameInfo struct {
	Round         int `json:"round"`
	ActivePlayers int `json:"activePlayers"`
}
