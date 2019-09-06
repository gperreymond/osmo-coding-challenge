package main

import (
	"github.com/moleculer-go/moleculer"
	"github.com/moleculer-go/moleculer/broker"
)

var playerService = moleculer.ServiceSchema{
	Name: "player",
	Actions: []moleculer.Action{
		{
			Name: "create",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) interface{} {
				ctx.Logger().Info("hello")
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
	bkr.Publish(playerService)
	bkr.Start()

	// small hack to make the process run forever
	select {}
}
