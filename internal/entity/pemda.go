package entity

import (
	"time"

	"gorm.io/gorm"
)

type Pemda struct {
	ID           uint   `gorm:"primaryKey"`
	PemdaName    string `gorm:"not null"`
	Phone        string
	SkpdOpd      string `gorm:"not null"`
	AgencyID     uint   `gorm:"not null"`
	Destination  string
	Consultation string
	Pokja        string
	Image        string `gorm:"not null"`
	Verified     bool   `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
