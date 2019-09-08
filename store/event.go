package store

import (
	"time"
)

// Event stores the data for every event
type Event struct {
	AggregateID   string      `bson:"aggregate_id"`
	AggregateType string      `bson:"aggregate_type"`
	EventType     string      `bson:"event_type"`
	CreatedAt     time.Time   `bson:"created_at"`
	Data          interface{} `bson:"data"`
}
