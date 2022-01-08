package godata

type GameSession struct {
	GameJoinUid string   `json:"gameJoinUid"`
	GameJoinUrl string   `json:"gameJoinUrl"`
	Players     []Player `json:"players"`
	GameInfo    GameInfo `json:"gameInfo"`
}
