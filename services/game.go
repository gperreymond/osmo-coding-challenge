package services

import (
	"log"

	"github.com/gperreymond/osmo-coding-challenge/models"
	"github.com/jinzhu/gorm"
	"github.com/moleculer-go/moleculer"
	"github.com/novalagung/gubrak"
)

// GameService ...
var GameService = moleculer.ServiceSchema{
	Name: "Game",
	Actions: []moleculer.Action{
		{
			Name: "Play",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) interface{} {
				ctx.Logger().Info("params: ", params)
				// Get all players
				db, err := gorm.Open("postgres", "host=localhost port=5432 user=infra dbname=osmo password=infra sslmode=disable")
				if err != nil {
					ctx.Logger().Error(err)
					return err
				}
				defer db.Close()
				var players []models.Player
				db.Find(&players)
				shuffle, _ := gubrak.Shuffle(players)
				mates := gubrak.RandomInt(3, 5)
				log.Println("Number of players in each team: ", mates)
				teamBlue, _ := gubrak.Take(shuffle, mates)
				log.Println("Team Blue: ", teamBlue)
				teamRed, _ := gubrak.TakeRight(shuffle, mates)
				log.Println("Team Red: ", teamRed)
				return true
			},
		},
	},
}

// GetGame ...
func GetGame() moleculer.ServiceSchema {
	return GameService
}
