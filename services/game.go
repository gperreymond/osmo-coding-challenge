package services

import (
	"log"
	"time"

	"github.com/gperreymond/osmo-coding-challenge/models"
	"github.com/gperreymond/osmo-coding-challenge/store"
	"github.com/gperreymond/osmo-coding-challenge/utils"
	"github.com/jinzhu/gorm"
	"github.com/moleculer-go/moleculer"
	"github.com/moleculer-go/moleculer/payload"
	"github.com/novalagung/gubrak"
)

const (
	// MaxHpLife ...
	MaxHpLife int = 100
)

// PlayerInGame ...
type PlayerInGame struct {
	AggregateID              string
	Name                     string
	Life                     int
	Loose                    bool
	NumberOfAttemptedAttacks int
	NumberOfHits             int
	TotalAmountOfDamageDone  int
	NumberOfKills            int
	NumberOfFirstHitKills    int
	NumberOfAssists          int
	TotalNumberOfSpellsCast  int
	TotalSpellDamageDone     int
}

// GetRandomAlivePlayerInTeam ...
func GetRandomAlivePlayerInTeam(team []PlayerInGame) (*PlayerInGame, bool) {
	loose := true
	var indexes []int
	for i, p := range team {
		log.Println("......", p.Name, p.Life, "life left")
		if p.Life > 0 {
			indexes = append(indexes, i)
			loose = false
		}
	}
	if loose == true {
		return nil, loose
	}
	rand := gubrak.RandomInt(1, len(indexes))
	i := indexes[rand-1]
	return &team[i], loose
}

// Even ...
func Even(number int) bool {
	return number%2 == 0
}

// GameService ...
var GameService = moleculer.ServiceSchema{
	Name: "Game",
	Events: []moleculer.Event{
		{
			Name: "Game.Created",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) {
				ctx.Logger().Info("params: ", params)
				// Save Event to store
				err := store.InsertEvent("Game", "Created", params)
				if err != nil {
					ctx.Logger().Error(err)
				}
			},
		},
		{
			Name: "Game.FinishUpdated",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) {
				ctx.Logger().Info("params: ", params)
				// Save Event to store
				err := store.InsertEvent("Game", "FinishUpdated", params)
				if err != nil {
					ctx.Logger().Error(err)
				}
			},
		},
	},
	Actions: []moleculer.Action{
		{
			Name: "Finish",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) interface{} {
				ctx.Logger().Info("params: ", params)
				aggregateID := params.Get("AggregateID").String()
				players := params.Get("Players")
				db, err := gorm.Open("postgres", "host=localhost port=5432 user=infra dbname=osmo password=infra sslmode=disable")
				defer db.Close()
				if err != nil {
					ctx.Logger().Error(err)
					return err
				}
				// **********************
				// Update a game in postgres
				// **********************
				duration := int64(gubrak.RandomInt(15, 30))
				finishedAt := time.Now().UTC().Add(15 * time.Minute)
				game := models.Game{StartedAt: finishedAt, AggregateID: aggregateID}
				err = db.Model(&game).Where("aggregate_id = ?", aggregateID).Updates(map[string]interface{}{"finishedAt": finishedAt}).Error
				if err != nil {
					ctx.Logger().Error(err)
					return err
				}
				// **********************
				// Emit All Events on bus
				// **********************
				ctx.Emit("Game.FinishUpdated", map[string]interface{}{
					"AggregateID": aggregateID,
					"Duration":    duration,
					"Players":     players,
				})
				participants := params.Get("Players").Array()
				for _, value := range participants {
					participant, _ := value.Value().(PlayerInGame)
					if participant.Loose == true {
						ctx.Emit("Player.GameLooseFinished", map[string]interface{}{
							"AggregateID": participant.AggregateID,
							"Game":        aggregateID,
						})
					} else {
						ctx.Emit("Player.GameWonFinished", map[string]interface{}{
							"AggregateID": participant.AggregateID,
							"Game":        aggregateID,
						})
					}
					ctx.Emit("Player.TotalTimePlayedUpdated", map[string]interface{}{
						"Game":                   aggregateID,
						"AggregateID":            participant.AggregateID,
						"TotalTimePlayedUpdated": duration,
					})
					ctx.Emit("Player.TotalAmountOfDamageDoneUpdated", map[string]interface{}{
						"Game":                    aggregateID,
						"AggregateID":             participant.AggregateID,
						"TotalAmountOfDamageDone": participant.TotalAmountOfDamageDone,
					})
					ctx.Emit("Player.NumberOfFirstHitKillsUpdated", map[string]interface{}{
						"Game":                  aggregateID,
						"AggregateID":           participant.AggregateID,
						"NumberOfFirstHitKills": participant.NumberOfFirstHitKills,
					})
					ctx.Emit("Player.NumberOfKillsUpdated", map[string]interface{}{
						"Game":          aggregateID,
						"AggregateID":   participant.AggregateID,
						"NumberOfKills": participant.NumberOfKills,
					})
					ctx.Emit("Player.NumberOfHitsUpdated", map[string]interface{}{
						"Game":         aggregateID,
						"AggregateID":  participant.AggregateID,
						"NumberOfHits": participant.NumberOfHits,
					})
					ctx.Emit("Player.NumberOfAttemptedAttacksUpdated", map[string]interface{}{
						"Game":                     aggregateID,
						"AggregateID":              participant.AggregateID,
						"NumberOfAttemptedAttacks": participant.NumberOfAttemptedAttacks,
					})
				}
				return map[string]interface{}{
					"AggregateID": aggregateID,
					"Duration":    duration,
				}
			},
		},
		{
			Name: "Create",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) interface{} {
				ctx.Logger().Info("params: ", params)
				aggregateID := utils.UUID()
				db, err := gorm.Open("postgres", "host=localhost port=5432 user=infra dbname=osmo password=infra sslmode=disable")
				defer db.Close()
				if err != nil {
					ctx.Logger().Error(err)
					return err
				}
				// Migrate the schema
				db.AutoMigrate(&models.Game{})
				// Create a new game in postgres
				startedAt := time.Now().UTC()
				game := models.Game{StartedAt: startedAt, AggregateID: aggregateID}
				err = db.Create(&game).Error
				if err != nil {
					ctx.Logger().Error(err)
					return err
				}
				// Emit Event on bus
				ctx.Emit("Game.Created", payload.New(map[string]interface{}{
					"StartedAt":   startedAt,
					"AggregateID": aggregateID,
				}))
				return map[string]string{
					"AggregateID": aggregateID,
				}
			},
		},
		{
			Name: "Play",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) interface{} {
				ctx.Logger().Info("params: ", params)
				log.Println("*********** ------------ ***********")
				log.Println("*********** GAME STARTED ***********")
				log.Println("*********** ------------ ***********")
				game := <-ctx.Call("Game.Create", payload.Empty())
				if game.IsError() {
					ctx.Logger().Error(game.Error())
					return game.Error()
				}
				gameAggregateID := game.Get("AggregateID").String()
				log.Println("Game aggregateID:", gameAggregateID)
				// Get all players
				db, err := gorm.Open("postgres", "host=localhost port=5432 user=infra dbname=osmo password=infra sslmode=disable")
				defer db.Close()
				if err != nil {
					ctx.Logger().Error(err)
					return err
				}
				var players []models.Player
				var pool []PlayerInGame
				var reporting []PlayerInGame
				db.Find(&players)
				for _, p := range players {
					inGame := PlayerInGame{
						AggregateID: p.AggregateID,
						Name:        p.Name,
						Life:        MaxHpLife,
					}
					pool = append(pool, inGame)
				}
				// Contruct teams
				var teamBlue []PlayerInGame
				var teamRed []PlayerInGame
				shuffle, _ := gubrak.Shuffle(pool)
				mates := gubrak.RandomInt(3, 5)
				log.Println("NumberOf players in each team: ", mates)
				t1, _ := gubrak.Take(shuffle, mates)
				teamBlue, _ = t1.([]PlayerInGame)
				teamBlueNames := ""
				for _, p := range teamBlue {
					teamBlueNames += p.Name + ", "
					// Emit Event on bus
					ctx.Emit("Player.GameStarted", payload.New(map[string]string{
						"Game":        gameAggregateID,
						"AggregateID": p.AggregateID,
					}))
				}
				log.Println("... Team Blue: ", teamBlueNames)
				t2, _ := gubrak.TakeRight(shuffle, mates)
				teamRed, _ = t2.([]PlayerInGame)
				teamRedNames := ""
				for _, p := range teamRed {
					teamRedNames += p.Name + ", "
					// Emit Event on bus
					ctx.Emit("Player.GameStarted", payload.New(map[string]string{
						"Game":        gameAggregateID,
						"AggregateID": p.AggregateID,
					}))
				}
				log.Println("... Team Red: ", teamRedNames)
				// *****************************************
				// FIGHT: ALL TURNS ARE MADE HERE!
				// *****************************************
				currentTurn := 1
				firstHitKills := false
				for {
					log.Println("Turn", currentTurn)
					currentTeam := &teamRed
					currentTargetTeam := &teamBlue
					if Even(currentTurn) {
						currentTeam = &teamBlue
						currentTargetTeam = &teamRed
					}
					teamWinnersNames := ""
					teamWLosersNames := ""
					log.Println("... Choose current player")
					currentPlayer, loose := GetRandomAlivePlayerInTeam(*currentTeam)
					log.Println("... Choose current target player")
					currentTargetPlayer, _ := GetRandomAlivePlayerInTeam(*currentTargetTeam)
					// Loose ? Then emit all events
					if loose == true {
						for _, p := range *currentTeam {
							teamWLosersNames += p.Name + ", "
							p.Loose = true
							reporting = append(reporting, p)
						}
						log.Println("Losers:", teamWLosersNames)
						for _, p := range *currentTargetTeam {
							teamWinnersNames += p.Name + ", "
							p.Loose = false
							reporting = append(reporting, p)
						}
						log.Println("Winners:", teamWinnersNames)
						break
					} else {
						log.Println("... CurrentPlayer: ", currentPlayer.Name, "has", currentPlayer.Life, "life")
						log.Println("... CurrentTargetPlayer", currentTargetPlayer.Name, "has", currentTargetPlayer.Life, "life")
					}
					// Attack
					damage := gubrak.RandomInt(1, 10) + gubrak.RandomInt(1, 10)
					miss := gubrak.RandomInt(1, 100)
					currentPlayer.NumberOfAttemptedAttacks++
					if miss <= 20 {
						log.Println("... Attack:", currentPlayer.Name, "miss", currentTargetPlayer.Name)
					} else {
						log.Println("... Attack:", currentPlayer.Name, "hits", damage, "damage to", currentTargetPlayer.Name)
						currentPlayer.NumberOfHits++
						currentPlayer.TotalAmountOfDamageDone += damage
						currentTargetPlayer.Life -= damage
						if currentTargetPlayer.Life < 0 {
							currentTargetPlayer.Life = 0
						}
					}
					if currentTargetPlayer.Life == 0 {
						currentPlayer.NumberOfKills++
						// First kill in the game
						if firstHitKills == false {
							log.Println("... ", currentPlayer.Name, "made first kill")
							currentPlayer.NumberOfFirstHitKills = 1
							firstHitKills = true
						}
					}
					// End of turn
					currentTurn++
				}
				res := <-ctx.Call("Game.Finish", payload.New(map[string]interface{}{
					"AggregateID": gameAggregateID,
					"Players":     reporting,
				}))
				if res.IsError() {
					ctx.Logger().Error(res.Error())
					return res.Error()
				}
				log.Println("*********** ------------- ***********")
				log.Println("*********** GAME FINISHED ***********")
				log.Println("*********** ------------- ***********")
				return map[string]interface{}{
					"Done": true,
				}
			},
		},
	},
}

// GetGame ...
func GetGame() moleculer.ServiceSchema {
	return GameService
}
