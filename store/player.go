package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// Player Model
type Player struct {
	ID         string `bson:"id" json:"id"`
	Name       string `bson:"name" json:"name"`
	GamesWon   int    `bson:"games_won" json:"games_won"`
	GamesLoose int    `bson:"games_loose" json:"games_loose"`
}

// GetPlayerAggregateID ...
func GetPlayerAggregateID(name string) (Player, error) {
	ctx := context.TODO()
	var result Player
	filter := bson.D{{Key: "data.name", Value: name}}
	conn, collection, err := Collection()
	if err != nil {
		return result, err
	}
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}
	Disconnect(conn)
	return result, nil
}
