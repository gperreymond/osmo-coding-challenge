package main

import (
	"log"
	"os"
	"time"

	"github.com/gperreymond/osmo-coding-challenge/services"
	"github.com/moleculer-go/moleculer"
	"github.com/moleculer-go/moleculer/broker"
	"github.com/moleculer-go/moleculer/payload"
	metrics "github.com/tevjef/go-runtime-metrics"
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
		services.GetAchievement(),
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
	time.Sleep(time.Millisecond * 1000)
	// Create the players
	names := []string{"Thrall", "Rexxar", "Gul'Dan", "Malfurion", "Garrosh", "Uther", "Anduin", "Valeera", "Morgl", "Medivh"}
	for _, name := range names {
		bkr.Call("Player.Create", payload.New(map[string]string{
			"Name": name,
		}))
	}
	bkr.Stop()
}

// AchievementsPlayer ...
func AchievementsPlayer(aggregateID string) {
	bkr := broker.New(&moleculer.Config{
		Transporter: "nats://nats.docker.localhost:4222",
		LogLevel:    "info",
		Metrics:     true,
	})
	bkr.Start()
	time.Sleep(time.Millisecond * 1000)
	// Control all Achievement for a player
	res := <-bkr.Call("Achievement.ControlBruiserAward", payload.New(map[string]string{
		"AggregateID":   aggregateID,
		"AggregateType": "Player",
		"EventType":     "TotalAmountOfDamageDoneUpdated",
	}))
	if res.IsError() {
		log.Println(res.Error())
	}
	log.Println(res)
	bkr.Stop()
}

// PlayGame ...
func PlayGame() {
	bkr := broker.New(&moleculer.Config{
		Transporter: "nats://nats.docker.localhost:4222",
		LogLevel:    "info",
		Metrics:     true,
	})
	bkr.Start()
	time.Sleep(time.Millisecond * 1000)
	// Play a game
	res := <-bkr.Call("Game.Play", payload.Empty())
	if res.IsError() {
		log.Println(res.Error())
	}
	bkr.Stop()
}

func main() {
	arg := os.Args[1]
	metrics.DefaultConfig.CollectionInterval = time.Second
	if err := metrics.RunCollector(metrics.DefaultConfig); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	switch arg {
	case "--achievements-player":
		AchievementsPlayer(os.Args[2])
	case "--play-game":
		PlayGame()
	case "--initialize":
		Initialize()
	case "--start-services":
		StartServices()
		// small hack to make the process run forever
		select {}
	default:
		log.Println("This mode is not allowed:", arg)
		os.Exit(0)
	}
}
