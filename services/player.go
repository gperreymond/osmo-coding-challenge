package services

import (
	"github.com/gperreymond/osmo-coding-challenge/models"
	"github.com/gperreymond/osmo-coding-challenge/store"
	"github.com/gperreymond/osmo-coding-challenge/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // ...
	"github.com/moleculer-go/moleculer"
)

// PlayerService ...
var PlayerService = moleculer.ServiceSchema{
	Name: "Player",
	Events: []moleculer.Event{
		{
			Name: "Player.Created",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) {
				ctx.Logger().Info("params: ", params)
				name := params.Get("Name").String()
				aggregateID := params.Get("AggregateID").String()
				// Save Event to store
				err := store.InsertEvent("Player", "Created", aggregateID, map[string]interface{}{
					"Name": name,
				})
				if err != nil {
					ctx.Logger().Error(err)
				}
			},
		},
	},
	Actions: []moleculer.Action{
		{
			Name: "Create",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) interface{} {
				ctx.Logger().Info("params: ", params)
				name := params.Get("Name").String()
				aggregateID := utils.UUID()
				db, err := gorm.Open("postgres", "host=localhost port=5432 user=infra dbname=osmo password=infra sslmode=disable")
				if err != nil {
					ctx.Logger().Error(err)
					return err
				}
				defer db.Close()
				// Migrate the schema
				db.AutoMigrate(&models.Player{})
				// Create a new user in postgres
				user := models.Player{Name: name, AggregateID: aggregateID}
				db.Create(&user)
				if err != nil {
					ctx.Logger().Error(err)
					return err
				}
				// Emit Event on bus
				ctx.Emit("Player.Created", map[string]interface{}{
					"Name":        name,
					"AggregateID": aggregateID,
				})
				return true
			},
		},
	},
}

// GetPlayer ...
func GetPlayer() moleculer.ServiceSchema {
	return PlayerService
}
