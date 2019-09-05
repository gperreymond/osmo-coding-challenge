package models

import uuid "github.com/satori/go.uuid"

// Player ...
type Player struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	GamesWon   int    `json:"gamesWon"`
	GamesLoose int    `json:"gamesLoose"`
}

// Create ...
func Create(name string) Player {
	id := uuid.NewV4()
	player := Player{id.String(), name, 0, 0}
	return player
}
