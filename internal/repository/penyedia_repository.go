package repository

import (
	"fmt"
	"time"

	"github.com/aliffatulmf/buku-tamu-apbj/internal/entity"

	"gorm.io/gorm"
)

type penyediaRepository struct {
	DB *gorm.DB
}

func NewPenyediaRepository(db *gorm.DB) *penyediaRepository {
	return &penyediaRepository{DB: db}
}

func (penyedia *penyediaRepository) New() *gorm.DB {
	tx := penyedia.DB.Model(entity.Provider{})
	tx.Session(&gorm.Session{
		QueryFields: true,
	})

	return tx
}

func (penyedia *penyediaRepository) Find() *gorm.DB {
	tx := penyedia.New()
	tx.Limit(100)

	return tx
}

func (penyedia *penyediaRepository) FindByID(model *entity.Provider, id uint) error {
	tx := penyedia.New()
	tx.First(model, "id = ?", id)

	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func (penyedia *penyediaRepository) Create(model *entity.Provider) error {
	return penyedia.New().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(model).Error; err != nil {
			return err
		}
		return nil
	})
}

func (penyedia *penyediaRepository) Delete(model *entity.Provider, conds ...interface{}) error {
	return penyedia.New().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(model, conds).Error; err != nil {
			return err
		}
		return nil
	})
}

func (penyedia *penyediaRepository) UpdatePermission(id uint) error {
	return penyedia.New().Transaction(func(tx *gorm.DB) error {
		result := tx.Where("id = ?", id).Update("verified", true)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
		return nil
	})
}

func (penyedia *penyediaRepository) Count() int64 {
	var col int64

	tx := penyedia.New()
	tx.Count(&col)
	return col
}

func (penyedia *penyediaRepository) CreateWithOmit(model *entity.Provider, omit ...string) error {
	return penyedia.New().Transaction(func(tx *gorm.DB) error {
		return tx.Omit(omit...).Create(model).Error
	})
}

func (penyedia *penyediaRepository) FindByDateRange(start, end time.Time) ([]entity.Provider, error) {
	var model []entity.Provider
	tx := penyedia.New()
	if !start.IsZero() && !end.IsZero() {
		tx.Where("created_at between ? and ?", start, end)
	}
	if err := tx.Scan(&model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (penyedia *penyediaRepository) FindWithFilter(sbn string, from, to time.Time, permission string) ([]entity.Provider, error) {
	var model []entity.Provider
	tx := penyedia.New()
	tx.Order("updated_at desc")

	if sbn != "" {
		arg := fmt.Sprintf("%%%s%%", sbn)
		tx.Where("provider_name like ?", arg)
	}

	if !from.IsZero() && !to.IsZero() {
		tx.Where("created_at between ? and ?", from, to)
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
