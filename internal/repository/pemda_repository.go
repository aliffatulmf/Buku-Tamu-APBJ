package repository

import (
	"fmt"
	"time"

	"github.com/aliffatulmf/buku-tamu-apbj/internal/entity"

	"gorm.io/gorm"
)

type pemdaRepository struct {
	DB *gorm.DB
}

func NewPemdaRepository(db *gorm.DB) *pemdaRepository {
	return &pemdaRepository{DB: db}
}

func (pemda *pemdaRepository) New() *gorm.DB {
	tx := pemda.DB.Table("pemdas")
	tx.Joins("inner join agencies on agencies.id = pemdas.agency_id")
	tx.Where("pemdas.deleted_at is null")

	return tx
}

func (pemda *pemdaRepository) Find() *gorm.DB {
	return pemda.New().Limit(100)
}

func (pemda *pemdaRepository) FindOne(model *entity.TypePemdaAgency) error {
	tx := pemda.New()
	tx.Scan(model)

	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func (pemda *pemdaRepository) Create(model *entity.Pemda) error {
	return pemda.New().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(model).Error; err != nil {
			return err
		}
		return nil
	})
}

func (pemda *pemdaRepository) Delete(id uint) error {
	return pemda.New().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(entity.Pemda{}, id).Error; err != nil {
			return err
		}
		return nil
	})
}

func (pemda *pemdaRepository) UpdatePermission(id uint) error {
	var model entity.Pemda

	return pemda.New().Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&model, id).Updates(entity.Pemda{Verified: true}).Error; err != nil {
			return err
		}
		return nil
	})
}

func (pemda *pemdaRepository) Count() int64 {
	var t int64

	tx := pemda.New()
	tx.Count(&t)

	return t
}

func (pemda *pemdaRepository) FindByID(id uint) (entity.TypePemdaAgency, error) {
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

func (pemda *pemdaRepository) CreateWithOmit(model *entity.Pemda, omit ...string) error {
	return pemda.New().Transaction(func(tx *gorm.DB) error {
		return tx.Omit(omit...).Create(model).Error
	})
}

func (pemda *pemdaRepository) FindByDateRange(start, end time.Time) ([]entity.TypePemdaAgency, error) {
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

func (pemda *pemdaRepository) FindWithFilter(sbn string, from, to time.Time, permission string) ([]entity.TypePemdaAgency, error) {
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
