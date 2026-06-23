package entity

import (
	"time"

	"gorm.io/gorm"
)

type TypeAgency struct {
	ID              uint
	AgencyName      string
	AgencyEmail     string
	AgencyTelephone string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

type TypePemdaAgency struct {
	ID           uint
	PemdaName    string
	Phone        string
	SkpdOpd      string
	Destination  string
	Consultation string
	Pokja        string
	Image        string
	Verified     bool
	Agency       TypeAgency `gorm:"embedded"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
