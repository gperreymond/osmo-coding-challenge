package main

import (
	"log"

	player "github.com/gperreymond/osmo-coding-challenge/store"

	"github.com/moleculer-go/moleculer"
	"github.com/moleculer-go/moleculer/broker"
)

var service = moleculer.ServiceSchema{
	Name: "Player",
	Actions: []moleculer.Action{
		{
			Name: "GetAggregateID",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) interface{} {
				ctx.Logger().Info("params: ", params)
				name := params.Get("name").String()
				data, err := player.GetAggregateID(name)
				if err != nil {
					log.Println(err)
					return false
				}
				return data
			},
		},
		{
			Name: "Create",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) interface{} {
				ctx.Logger().Info("params: ", params)
				name := params.Get("name").String()
				// Control if player already created
				control := <-ctx.Call("Player.GetAggregateID", map[string]interface{}{
					"name": name,
				})
				ctx.Logger().Info("control: ", control)
				// Prepare data
				/* data := player.Player{
					ID:         utils.UUID(),
					Name:       name,
					GamesWon:   0,
					GamesLoose: 0,
				}
				// Save Event to store
				err := eventstore.InsertEvent(data.ID, data)
				if err != nil {
					return false
				} */
				// Emit event on bus
				// ctx.Broadcast("player.created", data)
				return true
			},
		},
	},
}

func main() {
	bkr := broker.New(&moleculer.Config{
		Transporter: "nats://nats.docker.localhost:4222",
		LogLevel:    "info",
		Metrics:     true,
	})
	bkr.Publish(service)
	bkr.Start()

	// small hack to make the process run forever
	select {}
}
