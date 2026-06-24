package repository

import (
	"fmt"
	"time"

	"github.com/aliffatulmf/buku-tamu-apbj/internal/entity"

	"gorm.io/gorm"
)

type PemdaRepository struct {
	DB *gorm.DB
}

func NewPemdaRepository(db *gorm.DB) *PemdaRepository {
	return &PemdaRepository{DB: db}
}

func (pemda *PemdaRepository) New() *gorm.DB {
	tx := pemda.DB.Table("pemdas")
	tx.Joins("inner join agencies on agencies.id = pemdas.agency_id")
	tx.Where("pemdas.deleted_at is null")

	return tx
}

func (pemda *PemdaRepository) Find() *gorm.DB {
	return pemda.New().Limit(100)
}

func (pemda *PemdaRepository) FindOne(model *entity.TypePemdaAgency) error {
	tx := pemda.New()
	tx.Scan(model)

	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func (pemda *PemdaRepository) Create(model *entity.Pemda) error {
	return pemda.New().Create(model).Error
}

func (pemda *PemdaRepository) Delete(id uint) error {
	return pemda.New().Delete(entity.Pemda{}, id).Error
}

func (pemda *PemdaRepository) UpdatePermission(id uint) error {
	var model entity.Pemda
	return pemda.New().First(&model, id).Updates(entity.Pemda{Verified: true}).Error
}

func (pemda *PemdaRepository) Count() int64 {
	var t int64

	tx := pemda.New()
	tx.Count(&t)

	return t
}

func (pemda *PemdaRepository) FindByID(id uint) (entity.TypePemdaAgency, error) {
	var model entity.TypePemdaAgency
	tx := pemda.New()
	if err := tx.Where("pemdas.id = ?", id).Scan(&model).Error; err != nil {
		return entity.TypePemdaAgency{}, err
	}
	if model.ID < 1 {
		return entity.TypePemdaAgency{}, gorm.ErrRecordNotFound
	}
	return model, nil
}

func (pemda *PemdaRepository) CreateWithOmit(model *entity.Pemda, omit ...string) error {
	return pemda.New().Omit(omit...).Create(model).Error
}

func (pemda *PemdaRepository) FindByDateRange(start, end time.Time) ([]entity.TypePemdaAgency, error) {
	var model []entity.TypePemdaAgency
	tx := pemda.New()
	if !start.IsZero() && !end.IsZero() {
		tx.Where("pemdas.created_at between ? and ?", start, end)
	}
	if err := tx.Scan(&model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (pemda *PemdaRepository) FindWithFilter(sbn string, from, to time.Time, permission string) ([]entity.TypePemdaAgency, error) {
	var model []entity.TypePemdaAgency
	tx := pemda.New()
	tx.Order("updated_at desc")

	if sbn != "" {
		arg := fmt.Sprintf("%%%s%%", sbn)
		tx.Where("pemda_name LIKE ?", arg)
	}

	if !from.IsZero() && !to.IsZero() {
		tx.Where("pemdas.created_at BETWEEN ? AND ?", from, to)
	}

	switch permission {
	case "allowed":
		tx.Where("verified = ?", true)
	case "notallowed":
		tx.Where("verified = ?", false)
	}

	if err := tx.Scan(&model).Error; err != nil {
		return nil, err
	}
	return model, nil
}
