package services

import (
	store "github.com/gperreymond/osmo-coding-challenge/store"
	"github.com/gperreymond/osmo-coding-challenge/utils"

	"github.com/moleculer-go/moleculer"
)

// PlayerService ...
var PlayerService = moleculer.ServiceSchema{
	Name: "Player",
	Actions: []moleculer.Action{
		{
			Name: "FindAggregateID",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) interface{} {
				ctx.Logger().Info("params: ", params)
				name := params.Get("name").String()
				data, err := store.GetPlayerAggregateID(name)
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
				data := store.Player{
					ID:         utils.UUID(),
					Name:       name,
					GamesWon:   0,
					GamesLoose: 0,
				}
				// Save Event to store
				err := store.InsertEvent("Player", "Created", data.ID, data)
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

// GetPlayer ...
func GetPlayer() moleculer.ServiceSchema {
	return PlayerService
}
