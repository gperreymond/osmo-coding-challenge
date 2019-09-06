package main

import (
	"log"
	"time"

	Player "github.com/gperreymond/osmo-coding-challenge/models"
	"github.com/moleculer-go/moleculer"
	"github.com/moleculer-go/moleculer/broker"
	"github.com/novalagung/gubrak"
)

func main() {
	bkr := broker.New(&moleculer.Config{
		Transporter: "nats://nats.docker.localhost:4222",
		LogLevel:    "info",
		Metrics:     true,
	})
	bkr.Start()
	time.Sleep(time.Millisecond * 1000)

	// Create the players
	names := []string{"Thrall", "Rexxar", "Gul'Dan", "Malfurion", "Garrosh", "Uther", "Anduin", "Valeera", "Morgl", "Medivh"}
	players := []Player.Player{}
	for _, name := range names {
		player := Player.Create(name)
		players = append(players, player)
		log.Println("player.created", player)
		result := <-bkr.Call("player.create", "")
		log.Println(result)
	}

	// Run a game: Select players and make teams
	shuffle, _ := gubrak.Shuffle(players)
	mates := gubrak.RandomInt(3, 5)
	log.Println("Number of players in each team: ", mates)
	teamBlue, _ := gubrak.Take(shuffle, mates)
	log.Println("Team Blue: ", teamBlue)
	teamRed, _ := gubrak.TakeRight(shuffle, mates)
	log.Println("Team Red: ", teamRed)

	// Run a game: Select players and make teams

	bkr.Stop()
}
