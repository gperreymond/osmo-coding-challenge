package services

import (
	"log"

	"github.com/gperreymond/osmo-coding-challenge/models"
	"github.com/jinzhu/gorm"
	"github.com/moleculer-go/moleculer"
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
	NumberOfAttemptedAttacks int
	NumberOfHits             int
	TotalAmountOfDamageDone  int
	NumberOfKills            int
	NumberOfFirstHitKills    int
	NumberOfAssists          int
	TotalNumberOfSpellsCast  int
	TotalSpellDamageDone     int
	TotalTimePlayed          int
}

// GetRandomAlivePlayerInTeam ...
func GetRandomAlivePlayerInTeam(team []PlayerInGame) (*PlayerInGame, bool) {
	loose := true
	var indexes []int
	for i, p := range team {
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

// PhysicalAttackSuccess ...
func PhysicalAttackSuccess(number int) bool {
	i := gubrak.RandomInt(1, 100)
	return i < 50
}

// MagicAttackSuccess ...
func MagicAttackSuccess(number int) bool {
	i := gubrak.RandomInt(1, 100)
	return i < 50
}

// GameService ...
var GameService = moleculer.ServiceSchema{
	Name: "Game",
	Actions: []moleculer.Action{
		{
			Name: "Play",
			Handler: func(ctx moleculer.Context, params moleculer.Payload) interface{} {
				ctx.Logger().Info("params: ", params)
				log.Println("*********** ------------ ***********")
				log.Println("*********** GAME STARTED ***********")
				log.Println("*********** ------------ ***********")
				// Get all players
				db, err := gorm.Open("postgres", "host=localhost port=5432 user=infra dbname=osmo password=infra sslmode=disable")
				if err != nil {
					ctx.Logger().Error(err)
					return err
				}
				defer db.Close()
				var players []models.Player
				var pool []PlayerInGame
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
				}
				log.Println("... Team Blue: ", teamBlueNames)
				t2, _ := gubrak.TakeRight(shuffle, mates)
				teamRed, _ = t2.([]PlayerInGame)
				teamRedNames := ""
				for _, p := range teamRed {
					teamRedNames += p.Name + ", "
				}
				log.Println("... Team Red: ", teamRedNames)
				// FIGHT
				currentTurn := 1
				firstHitKills := false
				for {
					currentTeam := &teamRed
					currentTargetTeam := &teamBlue
					if Even(currentTurn) {
						currentTeam = &teamBlue
						currentTargetTeam = &teamRed
					}
					currentPlayer, loose := GetRandomAlivePlayerInTeam(*currentTeam)
					if loose == true {
						break
					}
					currentTargetPlayer, loose := GetRandomAlivePlayerInTeam(*currentTargetTeam)
					if loose == true {
						break
					}
					// Logs
					log.Println("Turn", currentTurn, ":", currentPlayer.Name, "attacks", currentTargetPlayer.Name)
					log.Println("... ", currentPlayer.Name, "has", currentPlayer.Life, "life")
					log.Println("... ", currentTargetPlayer.Name, "has", currentTargetPlayer.Life, "life")
					// Attack
					damage := gubrak.RandomInt(1, 10)
					log.Println("... ", currentPlayer.Name, "made", damage, "damage point(s)")
					currentPlayer.NumberOfAttemptedAttacks++
					currentPlayer.TotalAmountOfDamageDone += damage
					currentTargetPlayer.Life -= currentPlayer.TotalAmountOfDamageDone
					if currentTargetPlayer.Life < 0 {
						currentTargetPlayer.Life = 0
						// First kill in the game
						if firstHitKills == false {
							log.Println("... ", currentPlayer.Name, "made first kill")
							currentPlayer.NumberOfFirstHitKills++
							firstHitKills = true
						}
					}
					// End of turn
					currentTurn++
				}
				log.Println("*********** ------------- ***********")
				log.Println("*********** GAME FINISHED ***********")
				log.Println("*********** ------------- ***********")
				return true
			},
		},
	},
}

// GetGame ...
func GetGame() moleculer.ServiceSchema {
	return GameService
}
