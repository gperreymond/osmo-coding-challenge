package store

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/moleculer-go/moleculer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect ...
func Connect() (*mongo.Client, error) {
	ctx := context.TODO()
	opts := options.Client().ApplyURI("mongodb://infra:infra@localhost:27017")
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Disconnect ...
func Disconnect(conn *mongo.Client) {
	ctx := context.TODO()
	conn.Disconnect(ctx)
}

// Collection ...
func Collection() (*mongo.Client, *mongo.Collection, error) {
	conn, err := Connect()
	if err != nil {
		return nil, nil, err
	}
	collection := conn.Database("osmo").Collection("eventstore")
	return conn, collection, nil
}

// Aggregate ...
func Aggregate(params moleculer.Payload) error {
	ctx := context.TODO()
	// AggregateID is mandatory
	if params.Get("AggregateID").Exists() == false {
		return errors.New("AggregateID is mandatory")
	}
	aggregateID := params.Get("AggregateID").String()
	conn, collection, err := Collection()
	if err != nil {
		return err
	}
	filter := bson.D{
		{Key: "aggregate_id", Value: aggregateID},
	}
	var items []Event
	err = collection.Find(ctx, filter).Sort("-created_at").All(&items)
	if err != nil {
		return err
	}
	// cursor.Sort(sort)
	for cursor.Next(ctx) {
		elem := &bson.D{}
		err = cursor.Decode(elem)
		log.Println(elem)
		if err != nil {
			return err
		}
	}
	Disconnect(conn)
	return nil
}

// InsertEvent ...
func InsertEvent(aggregateType string, eventType string, params moleculer.Payload) error {
	// AggregateID is mandatory
	if params.Get("AggregateID").Exists() == false {
		return errors.New("AggregateID is mandatory")
	}
	aggregateID := params.Get("AggregateID").String()
	ctx := context.TODO()
	conn, collection, err := Collection()
	if err != nil {
		return err
	}
	e := Event{
		AggregateID:   aggregateID,
		AggregateType: aggregateType,
		EventType:     eventType,
		CreatedAt:     time.Now().UTC(),
		Data:          params.Bson(),
	}
	_, err = collection.InsertOne(ctx, e)
	if err != nil {
		return err
	}
	Disconnect(conn)
	return nil
}
