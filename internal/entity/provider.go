package entity

import (
	"time"

	"gorm.io/gorm"
)

type Provider struct {
	ID           uint   `gorm:"primaryKey"`
	ProviderName string `gorm:"index,unique;not null"`
	Phone        string
	Company      string `gorm:"not null"`
	Description  string
	Destination  string `gorm:"not null"`
	Consultation string
	Pokja        string
	Image        string `gorm:"not null"`
	Verified     bool   `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
