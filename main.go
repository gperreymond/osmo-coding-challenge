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
		bkr.Call("Player.Create", map[string]interface{}{
			"Name": name,
		})
	}
	bkr.Stop()
}

// StatsPlayer ...
func StatsPlayer(aggregateID string) {
	bkr := broker.New(&moleculer.Config{
		Transporter: "nats://nats.docker.localhost:4222",
		LogLevel:    "info",
		Metrics:     true,
	})
	bkr.Start()
	time.Sleep(time.Millisecond * 1000)
	// Play a game
	res := <-bkr.Call("Player.AggregateStats", map[string]interface{}{
		"AggregateID": aggregateID,
	})
	if res.IsError() {
		log.Println(res.Error())
	}
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
	case "--stats-player":
		StatsPlayer(os.Args[2])
	case "--play-game":
		PlayGame()
	case "--initialize":
		Initialize()
	case "--start-services":
		StartServices()
		// small hack to make the process run forever
		select {}
	}
}
