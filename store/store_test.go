package store_test

import (
	"github.com/moleculer-go/moleculer/payload"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/gperreymond/osmo-coding-challenge/store"
)

var _ = Describe("Store", func() {
	It("should fail to insert an event", func() {
		err := InsertEvent("Player", "GameStarted", payload.Empty())
		Expect(err).To(HaveOccurred())
	})
})
