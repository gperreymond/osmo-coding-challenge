package services_test

import (
	"github.com/jinzhu/gorm"
	"github.com/moleculer-go/moleculer"
	"github.com/moleculer-go/moleculer/broker"
	"github.com/moleculer-go/moleculer/payload"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	models "github.com/gperreymond/osmo-coding-challenge/models"
	. "github.com/gperreymond/osmo-coding-challenge/services"
)

var _ = Describe("Service", func() {
	It("should delete a player if exists", func() {
		db, _ := gorm.Open("postgres", "host=localhost port=5432 user=infra dbname=osmo password=infra sslmode=disable")
		defer db.Close()
		db.Unscoped().Where("name LIKE ?", "%test%").Delete(models.Player{})
	})
	It("should not create a player, because Name is mandatory", func() {
		bkr := broker.New(&moleculer.Config{
			Transporter: "nats://localhost:4222",
			LogLevel:    "fatal",
			Metrics:     false,
		})
		bkr.Publish(
			GetPlayer(),
		)
		bkr.Start()
		res := <-bkr.Call("Player.Create", payload.Empty())
		Expect(res.IsError()).To(Equal(true))
		Expect(res.String()).To(Equal("Name is mandatory"))
	})
	It("should succesfully create a player", func() {
		bkr := broker.New(&moleculer.Config{
			Transporter: "nats://localhost:4222",
			LogLevel:    "fatal",
			Metrics:     false,
		})
		bkr.Publish(
			GetPlayer(),
		)
		bkr.Start()
		res := <-bkr.Call("Player.Create", payload.New(map[string]string{
			"Name": "test",
		}))
		Expect(res.IsError()).To(Equal(false))
	})
	It("should not create a player, because Player.Name is already used", func() {
		bkr := broker.New(&moleculer.Config{
			Transporter: "nats://localhost:4222",
			LogLevel:    "fatal",
			Metrics:     false,
		})
		bkr.Publish(
			GetPlayer(),
		)
		bkr.Start()
		res := <-bkr.Call("Player.Create", payload.New(map[string]string{
			"Name": "test",
		}))
		Expect(res.IsError()).To(Equal(true))
		Expect(res.String()).To(Equal("pq: duplicate key value violates unique constraint \"uix_players_name\""))
	})
	It("should delete a player if exists", func() {
		db, _ := gorm.Open("postgres", "host=localhost port=5432 user=infra dbname=osmo password=infra sslmode=disable")
		defer db.Close()
		db.Unscoped().Where("name LIKE ?", "%test%").Delete(models.Player{})
	})
	It("should successfully play a game", func() {
		bkr := broker.New(&moleculer.Config{
			Transporter: "nats://localhost:4222",
			LogLevel:    "fatal",
			Metrics:     false,
		})
		bkr.Publish(
			GetPlayer(),
			GetGame(),
		)
		bkr.Start()
		res := <-bkr.Call("Game.Play", payload.Empty())
		Expect(res.IsError()).To(Equal(false))
	})
})
