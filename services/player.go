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
	Actions: []moleculer.Action{
		{
			Name: "Create",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) interface{} {
				ctx.Logger().Info("params: ", params)
				name := params.Get("name").String()
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
				// Save Event to store
				data := map[string]interface{}{
					"Name":        name,
					"AggregateID": aggregateID,
				}
				err = store.InsertEvent("Player", "Created", aggregateID, data)
				if err != nil {
					ctx.Logger().Error(err)
					return err
				}
				// Emit Event on bus
				ctx.Emit("Player.Created", data)
				return true
			},
		},
	},
}

// GetPlayer ...
func GetPlayer() moleculer.ServiceSchema {
	return PlayerService
}
