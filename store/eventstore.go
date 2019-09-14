package store

import (
	"errors"
	"time"

	"github.com/moleculer-go/moleculer"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

// Config ...
type Config struct {
	Address  string
	Database string
}

var configStore Config

// SetConfigStore ...
func SetConfigStore(value Config) {
	configStore = value
}

// Session ...
func Session() (*r.Session, error) {
	if configStore.Address == "" {
		configStore = Config{
			Address:  "localhost:28015",
			Database: "osmo",
		}
	}
	session, err := r.Connect(r.ConnectOpts{
		Address:  configStore.Address,
		Database: configStore.Database,
	})
	if err != nil {
		return session, err
	}
	return session, nil
}

// InsertEvent ...
func InsertEvent(aggregateType string, eventType string, params moleculer.Payload) error {
	// AggregateID is mandatory
	if params.Get("AggregateID").Exists() == false {
		return errors.New("AggregateID is mandatory")
	}
	aggregateID := params.Get("AggregateID").String()
	session, err := Session()
	defer session.Close()
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
	_ = r.Table("eventstore").Insert(e).Exec(session, r.ExecOpts{
		NoReply: true,
	})
	return nil
}
