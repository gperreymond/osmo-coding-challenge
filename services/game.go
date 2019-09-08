package services

import (
	"github.com/moleculer-go/moleculer"
)

// GameService ...
var GameService = moleculer.ServiceSchema{
	Name:    "Game",
	Actions: []moleculer.Action{},
}

// GetGame ...
func GetGame() moleculer.ServiceSchema {
	return GameService
}
