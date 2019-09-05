package main

import (
	"log"

	Player "github.com/gperreymond/osmo-coding-challenge/models"
	"github.com/novalagung/gubrak"
)

func main() {
	// Create the players
	names := []string{"Thrall", "Rexxar", "Gul'Dan", "Malfurion", "Garrosh", "Uther", "Anduin", "Valeera", "Morgl", "Medivh"}
	players := []Player.Player{}
	for _, name := range names {
		player := Player.Create(name)
		players = append(players, player)
		log.Println("player.created", player)
	}
	// Run a game: Select players and make teams
	shuffle, _ := gubrak.Shuffle(players)
	mates := gubrak.RandomInt(3, 5)
	log.Println("Number of players in each team: ", mates)
	teamBlue, _ := gubrak.Take(shuffle, mates)
	log.Println("Team Blue: ", teamBlue)
	teamRed, _ := gubrak.TakeRight(shuffle, mates)
	log.Println("Team Red: ", teamRed)
}
