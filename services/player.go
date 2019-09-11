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
			Name: "Player.GameStarted",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) {
				ctx.Logger().Info("params: ", params)
				// Save Event to store
				err := store.InsertEvent("Player", "GameStarted", params)
				if err != nil {
					ctx.Logger().Error(err)
				}
			},
		},
		{
			Name: "Player.GameWonFinished",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) {
				ctx.Logger().Info("params: ", params)
				// Save Event to store
				err := store.InsertEvent("Player", "GameWonFinished", params)
				if err != nil {
					ctx.Logger().Error(err)
				}
			},
		},
		{
			Name: "Player.GameLooseFinished",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) {
				ctx.Logger().Info("params: ", params)
				// Save Event to store
				err := store.InsertEvent("Player", "GameLooseFinished", params)
				if err != nil {
					ctx.Logger().Error(err)
				}
			},
		},
		{
			Name: "Player.NumberOfKillsUpdated",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) {
				ctx.Logger().Info("params: ", params)
				// Save Event to store
				err := store.InsertEvent("Player", "NumberOfKillsUpdated", params)
				if err != nil {
					ctx.Logger().Error(err)
				}
			},
		},
		{
			Name: "Player.NumberOfFirstHitKillsUpdated",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) {
				ctx.Logger().Info("params: ", params)
				// Save Event to store
				err := store.InsertEvent("Player", "NumberOfFirstHitKillsUpdated", params)
				if err != nil {
					ctx.Logger().Error(err)
				}
			},
		},
		{
			Name: "Player.TotalAmountOfDamageDoneUpdated",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) {
				ctx.Logger().Info("params: ", params)
				// Save Event to store
				err := store.InsertEvent("Player", "TotalAmountOfDamageDoneUpdated", params)
				if err != nil {
					ctx.Logger().Error(err)
				}
			},
		},
		{
			Name: "Player.Created",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) {
				ctx.Logger().Info("params: ", params)
				// Save Event to store
				err := store.InsertEvent("Player", "Created", params)
				if err != nil {
					ctx.Logger().Error(err)
				}
			},
		},
	},
	Actions: []moleculer.Action{
		{
			Name: "AggregateStats",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) interface{} {
				ctx.Logger().Info("params: ", params)
				store.Aggregate(params)
				return true
			},
		},
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
				err = db.Create(&user).Error
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
