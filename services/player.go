package services

import (
	"errors"

	"github.com/gperreymond/osmo-coding-challenge/models"
	"github.com/gperreymond/osmo-coding-challenge/store"
	"github.com/gperreymond/osmo-coding-challenge/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // ...
	"github.com/moleculer-go/moleculer"
	"github.com/moleculer-go/moleculer/payload"
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
			Name: "Player.NumberOfLandingAttacksUpdated",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) {
				ctx.Logger().Info("params: ", params)
				// Save Event to store
				err := store.InsertEvent("Player", "NumberOfLandingAttacksUpdated", params)
				if err != nil {
					ctx.Logger().Error(err)
				}
			},
		},
		{
			Name: "Player.NumberOfHitsUpdated",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) {
				ctx.Logger().Info("params: ", params)
				// Save Event to store
				err := store.InsertEvent("Player", "NumberOfHitsUpdated", params)
				if err != nil {
					ctx.Logger().Error(err)
				}
			},
		},
		{
			Name: "Player.NumberOfAttemptedAttacksUpdated",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) {
				ctx.Logger().Info("params: ", params)
				// Save Event to store
				err := store.InsertEvent("Player", "NumberOfAttemptedAttacksUpdated", params)
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
			Name: "Player.TotalTimePlayedUpdated",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) {
				ctx.Logger().Info("params: ", params)
				// Save Event to store
				err := store.InsertEvent("Player", "TotalTimePlayedUpdated", params)
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
			Name: "Create",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) interface{} {
				ctx.Logger().Info("params: ", params)
				// AggregateID is mandatory
				if params.Get("Name").Exists() == false {
					return errors.New("Name is mandatory")
				}
				name := params.Get("Name").String()
				aggregateID := utils.UUID()
				db, err := gorm.Open("postgres", "host=localhost port=5432 user=infra dbname=osmo password=infra sslmode=disable")
				defer db.Close()
				if err != nil {
					ctx.Logger().Error(err)
					return err
				}
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
				ctx.Emit("Player.Created", payload.New(map[string]string{
					"Name":        name,
					"AggregateID": aggregateID,
				}))
				return map[string]bool{
					"Done": true,
				}
			},
		},
	},
}

// GetPlayer ...
func GetPlayer() moleculer.ServiceSchema {
	return PlayerService
}
