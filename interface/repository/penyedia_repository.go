package repository

import (
	"errors"

	"github.com/aliffatulmf/buku-tamu-apbj/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PenyediaRepository struct {
	DB *gorm.DB
}

func NewPenyediaRepository(db *gorm.DB) PenyediaRepository {
	return PenyediaRepository{DB: db}
}

func (penyedia PenyediaRepository) New() *gorm.DB {
	tx := penyedia.DB.Model(entity.Provider{})
	tx.Session(&gorm.Session{
		QueryFields: true,
	})

	return tx
}

func (penyedia PenyediaRepository) Find() *gorm.DB {
	tx := penyedia.New()
	tx.Limit(100)

	return tx
}

func (penyedia PenyediaRepository) FindByID(model *entity.Provider, id string) error {
	tx := penyedia.New()
	tx.First(model, "id = ?", id)

	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func (penyedia PenyediaRepository) Create(model *entity.Provider) error {
	return penyedia.New().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(model).Error; err != nil {
			return err
		}
		return nil
	})
}

func (penyedia PenyediaRepository) Delete(model *entity.Provider, conds ...interface{}) error {
	return penyedia.New().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(model, conds).Error; err != nil {
			return err
		}
		return nil
	})
}

func (penyedia PenyediaRepository) UpdatePermission(id string) error {
	var model entity.Provider

	return penyedia.New().Transaction(func(tx *gorm.DB) error {
		err := tx.First(&model, id).Update("verified", true).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return logger.ErrRecordNotFound
		}
		return nil
	})
}

func (penyedia PenyediaRepository) Count() int64 {
	var col int64

	tx := penyedia.New()
	tx.Count(&col)
	return col
}
