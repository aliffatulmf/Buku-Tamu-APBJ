package entity

import (
	"time"

	"gorm.io/gorm"
)

type Agency struct {
	ID              uint   `gorm:"primaryKey"`
	AgencyName      string `gorm:"index:,unique;not null"`
	AgencyEmail     string
	AgencyTelephone string
	Pemdas          []Pemda `gorm:"foreignKey:AgencyID"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}
