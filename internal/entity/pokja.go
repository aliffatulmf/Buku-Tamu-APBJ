package entity

import (
	"time"

	"gorm.io/gorm"
)

type Pokja struct {
	ID            string `gorm:"type:varchar(36);primaryKey"`
	PokjaName     string `gorm:"index:,unique;not null"`
	Status        bool
	DestinationID string `gorm:"type:varchar(36)"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
