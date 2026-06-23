package entity

import (
	"time"

	"gorm.io/gorm"
)

type PemdaRepository interface {
	New() *gorm.DB
	FindByID(id uint) (TypePemdaAgency, error)
	FindWithFilter(sbn string, from, to time.Time, permission string) ([]TypePemdaAgency, error)
	CreateWithOmit(model *Pemda, omit ...string) error
	Delete(id uint) error
	UpdatePermission(id uint) error
	Count() int64
	FindByDateRange(start, end time.Time) ([]TypePemdaAgency, error)
}

type PenyediaRepository interface {
	New() *gorm.DB
	Find() *gorm.DB
	FindByID(model *Provider, id uint) error
	FindWithFilter(sbn string, from, to time.Time, permission string) ([]Provider, error)
	CreateWithOmit(model *Provider, omit ...string) error
	Delete(model *Provider, conds ...interface{}) error
	UpdatePermission(id uint) error
	Count() int64
	FindByDateRange(start, end time.Time) ([]Provider, error)
}

type InstansiRepository interface {
	New() *gorm.DB
	Create(model *Agency) error
	Find(model *[]Agency) error
	FindBy(model *Agency, conds ...interface{}) error
	FindWithFilter(sbn string, from, to time.Time) ([]Agency, error)
	Update(model *Agency, conds ...interface{}) error
	Count() int64
}

type PokjaRepository interface {
	New() *gorm.DB
	Create(model *Pokja) error
	Find(model *[]Pokja) error
	FindByID(id string) (Pokja, error)
	UpdateStatus(id string, status bool) error
	Delete(id string) error
	Count() int64
}

type TujuanRepository interface {
	New(model interface{}) *gorm.DB
}
