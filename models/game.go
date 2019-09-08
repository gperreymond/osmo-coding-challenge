package models

import "github.com/jinzhu/gorm"

// Game Model
type Game struct {
	gorm.Model
	AggregateID string `gorm:"type:varchar(100);unique_index" json:"aggregate_id"`
}
