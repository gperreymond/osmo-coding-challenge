package store_test

import (
	"github.com/moleculer-go/moleculer/payload"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/gperreymond/osmo-coding-challenge/store"
)

var _ = Describe("Store", func() {
	It("should fail to insert an event because AggregateID is mandatory", func() {
		err := InsertEvent("Player", "GameStarted", payload.Empty())
		Expect(err).To(HaveOccurred())
	})
	It("should fail to insert an event", func() {
		SetConfigStore(Config{
			Address:  "test:28015",
			Database: "osmo",
		})
		err := InsertEvent("Player", "GameStarted", payload.New(map[string]string{
			"AggregateID": "test",
		}))
		Expect(err).To(HaveOccurred())
	})
	It("should successfully inserting an event", func() {
		SetConfigStore(Config{
			Address:  "localhost:28015",
			Database: "osmo",
		})
		err := InsertEvent("Player", "GameStarted", payload.New(map[string]string{
			"AggregateID":   "test",
			"AggregateType": "test",
			"EventType":     "test",
		}))
		Expect(err).NotTo(HaveOccurred())
	})
})
