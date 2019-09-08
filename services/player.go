package main

import (
	eventstore "github.com/gperreymond/osmo-coding-challenge/store"
	player "github.com/gperreymond/osmo-coding-challenge/store"
	"github.com/gperreymond/osmo-coding-challenge/utils"

	"github.com/moleculer-go/moleculer"
	"github.com/moleculer-go/moleculer/broker"
)

var service = moleculer.ServiceSchema{
	Name: "Player",
	Actions: []moleculer.Action{
		{
			Name: "FindAggregateID",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) interface{} {
				ctx.Logger().Info("params: ", params)
				name := params.Get("name").String()
				data, err := player.GetAggregateID(name)
				if err != nil {
					return err
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
				res := <-ctx.Call("Player.FindAggregateID", map[string]interface{}{
					"name": name,
				})
				ctx.Logger().Info("control: ", res)
				if !res.IsError() {
					return false
				}
				// Prepare data
				data := player.Player{
					ID:         utils.UUID(),
					Name:       name,
					GamesWon:   0,
					GamesLoose: 0,
				}
				// Save Event to store
				err := eventstore.InsertEvent("Player", "Created", data.ID, data)
				if err != nil {
					return false
				}
				// Emit Event on bus
				ctx.Broadcast("Player.Created", data)
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
