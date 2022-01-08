package godata

type Player struct {
	// Unique ID of for the player. Pulled from their local app instance
	PlayerUid string `json:"playerUid"`
	// The player's name
	PlayerName string `json:"playerName"`
	// A list of active session IDs
	ActiveSessions []string `json:"activeSessions"`
}
