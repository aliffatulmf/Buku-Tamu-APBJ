package entity

import (
	"time"

	"gorm.io/gorm"
)

type Destination struct {
	ID              string         `gorm:"type:varchar(36);primaryKey"`
	DestinationName string         `gorm:"index:,unique;not null"`
	Consultations   []Consultation `gorm:"foreignKey:DestinationID"`
	Pokjas          []Pokja        `gorm:"foreignKey:DestinationID"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}
