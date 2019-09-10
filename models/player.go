package models

import (
	"github.com/jinzhu/gorm"
)

// Player Model
type Player struct {
	gorm.Model
	AggregateID string `gorm:"type:varchar(100);unique_index" json:"aggregate_id"`
	Name        string `gorm:"type:varchar(100);unique_index" json:"name"`
}
