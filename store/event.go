package store

import "time"

// Event stores the data for every event
type Event struct {
	ID            string      `bson:"id" json:"id"`
	AggregateID   string      `bson:"aggregate_id" json:"aggregate_id"`
	AggregateType string      `bson:"aggregate_type" json:"aggregate_type"`
	Type          string      `bson:"type" json:"type"`
	CreatedAt     time.Time   `bson:"created_at" json:"created_at"`
	Data          interface{} `bson:"data" json:"data"`
}
