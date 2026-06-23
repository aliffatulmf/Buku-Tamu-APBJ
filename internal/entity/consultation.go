package entity

import (
	"time"

	"gorm.io/gorm"
)

type Consultation struct {
	ID               string `gorm:"type:varchar(36);primaryKey"`
	ConsultationName string `gorm:"index:,unique;not null"`
	DestinationID    string `gorm:"type:varchar(36)"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}
