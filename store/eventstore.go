package store

import (
	"errors"
	"log"
	"time"

	"github.com/moleculer-go/moleculer"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

// InsertEvent ...
func InsertEvent(aggregateType string, eventType string, params moleculer.Payload) error {
	log.Println("Store InsertEvent", params)
	// AggregateID is mandatory
	if params.Get("AggregateID").Exists() == false {
		return errors.New("AggregateID is mandatory")
	}
	aggregateID := params.Get("AggregateID").String()
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "osmo",
	})
	if err != nil {
		return err
	}
	defer session.Close()
	e := Event{
		AggregateID:   aggregateID,
		AggregateType: aggregateType,
		EventType:     eventType,
		CreatedAt:     time.Now().UTC(),
		Data:          params.Bson(),
	}
	err = r.Table("eventstore").Insert(e).Exec(session, r.ExecOpts{
		NoReply: true,
	})
	if err != nil {
		return err
	}
	return nil
}
