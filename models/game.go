package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Game Model
type Game struct {
	gorm.Model
	AggregateID string    `gorm:"type:varchar(100);unique_index" json:"aggregate_id"`
	StartedAt   time.Time `json:"started_at"`
	FinishedAt  time.Time `json:"finished_at"`
}
