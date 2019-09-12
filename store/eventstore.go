package store

import (
	"errors"
	"log"
	"time"

	"github.com/moleculer-go/moleculer"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

// GetEvent ...
func GetEvent(params moleculer.Payload) error {
	log.Println("Store InsertEvent", params)
	// AggregateID is mandatory
	if params.Get("AggregateID").Exists() == false {
		return errors.New("AggregateID is mandatory")
	}
	// AggregateType is mandatory
	if params.Get("AggregateType").Exists() == false {
		return errors.New("AggregateType is mandatory")
	}
	// EventType is mandatory
	if params.Get("EventType").Exists() == false {
		return errors.New("EventType is mandatory")
	}
	aggregateID := params.Get("AggregateID").String()
	aggregateType := params.Get("AggregateType").String()
	eventType := params.Get("EventType").String()
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "osmo",
	})
	if err != nil {
		return err
	}
	defer session.Close()
	type Filter struct {
		AggregateID   string
		AggregateType string
		EventType     string
	}
	f := Filter{
		AggregateID:   aggregateID,
		AggregateType: aggregateType,
		EventType:     eventType,
	}
	err = r.Table("eventstore").Filter(f, r.FilterOpts{}).Exec(session)
	log.Println(err)
	if err != nil {
		return err
	}
	return nil
}

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
