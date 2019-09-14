package services

import (
	"errors"

	"github.com/gperreymond/osmo-coding-challenge/store"
	"github.com/moleculer-go/moleculer"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

// AchievementService ...
var AchievementService = moleculer.ServiceSchema{
	Name:   "Achievement",
	Events: []moleculer.Event{},
	Actions: []moleculer.Action{
		{
			Name: "ControlBruiserAward",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) interface{} {
				ctx.Logger().Info("params: ", params)
				// AggregateID is mandatory
				if params.Get("AggregateID").Exists() == false {
					ctx.Logger().Error("AggregateID is mandatory")
					return errors.New("AggregateID is mandatory")
				}
				aggregateID := params.Get("AggregateID").String()
				session, err := store.Session()
				defer session.Close()
				if err != nil {
					ctx.Logger().Error(err)
					return err
				}
				type Filter struct {
					AggregateID   string
					AggregateType string
					EventType     string
				}
				// ReQL: Translate to GO
				res, _ := r.Table("eventstore").Filter(Filter{
					AggregateID:   aggregateID,
					AggregateType: "Player",
					EventType:     "TotalAmountOfDamageDoneUpdated",
				}).Map(func(row r.Term) interface{} {
					return row.Field("Data")
				}).Max("TotalAmountOfDamageDone").Field("TotalAmountOfDamageDone").Ge(500).Run(session)
				var badge interface{}
				err = res.One(&badge)
				if err != nil {
					ctx.Logger().Error(err)
					return err
				}
				return map[string]interface{}{
					"Achievement": "BruiserAward",
					"Badge":       badge,
				}
			},
		},
	},
}

// GetAchievement ...
func GetAchievement() moleculer.ServiceSchema {
	return AchievementService
}
