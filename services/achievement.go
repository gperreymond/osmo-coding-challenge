package services

import (
	"errors"
	"log"

	"github.com/moleculer-go/moleculer"
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
				// AggregateType is mandatory
				if params.Get("AggregateType").Exists() == false {
					ctx.Logger().Error("AggregateType is mandatory")
					return errors.New("AggregateType is mandatory")
				}
				// EventType is mandatory
				if params.Get("EventType").Exists() == false {
					ctx.Logger().Error("EventType is mandatory")
					return errors.New("EventType is mandatory")
				}
				aggregateID := params.Get("AggregateID").String()
				aggregateType := params.Get("AggregateType").String()
				eventType := params.Get("EventType").String()
				log.Println(aggregateID, aggregateType, eventType)
				//res := store.AggregateByEventType(params)
				// log.Println(res)
				return map[string]interface{}{
					"Done": true,
				}
			},
		},
	},
}

// GetAchievement ...
func GetAchievement() moleculer.ServiceSchema {
	return AchievementService
}
