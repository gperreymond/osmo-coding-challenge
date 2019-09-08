package main

import (
	"os"
	"time"

	services "github.com/gperreymond/osmo-coding-challenge/services"
	"github.com/moleculer-go/moleculer"
	"github.com/moleculer-go/moleculer/broker"
)

// StartServices ...
func StartServices() {
	bkr := broker.New(&moleculer.Config{
		Transporter: "nats://nats.docker.localhost:4222",
		LogLevel:    "info",
		Metrics:     true,
	})
	bkr.Publish(
		services.GetGame(),
		services.GetPlayer(),
	)
	bkr.Start()
}

// Initialize ...
func Initialize() {
	bkr := broker.New(&moleculer.Config{
		Transporter: "nats://nats.docker.localhost:4222",
		LogLevel:    "info",
		Metrics:     true,
	})
	bkr.Start()
	time.Sleep(time.Millisecond * 2000)
	// Create the players
	names := []string{"Thrall", "Rexxar", "Gul'Dan", "Malfurion", "Garrosh", "Uther", "Anduin", "Valeera", "Morgl", "Medivh"}
	for _, name := range names {
		bkr.Call("Player.Create", map[string]interface{}{
			"name": name,
		})
	}
	time.Sleep(time.Millisecond * 2000)
	bkr.Stop()
}

func main() {
	arg := os.Args[1]
	switch arg {
	case "--initialize":
		Initialize()
	case "--start-services":
		StartServices()
		// small hack to make the process run forever
		select {}
	}
}

/* shuffle, _ := gubrak.Shuffle(players)
mates := gubrak.RandomInt(3, 5)
log.Println("Number of players in each team: ", mates)
teamBlue, _ := gubrak.Take(shuffle, mates)
log.Println("Team Blue: ", teamBlue)
teamRed, _ := gubrak.TakeRight(shuffle, mates)
log.Println("Team Red: ", teamRed) */
