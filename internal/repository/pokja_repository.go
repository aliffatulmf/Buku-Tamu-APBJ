package repository

import (
	"github.com/aliffatulmf/buku-tamu-apbj/internal/entity"

	"gorm.io/gorm"
)

type pokjaRepository struct {
	DB *gorm.DB
}

func NewPokjaRepository(db *gorm.DB) *pokjaRepository {
	return &pokjaRepository{DB: db}
}

func (pokja *pokjaRepository) New() *gorm.DB {
	tx := pokja.DB.Model(entity.Pokja{})
	tx.Session(&gorm.Session{
		QueryFields: true,
	})

	return tx
}

func (pokja *pokjaRepository) Create(model *entity.Pokja) error {
	tx := pokja.New()
	return tx.Transaction(func(trx *gorm.DB) error {
		if err := trx.Create(model).Error; err != nil {
			return err
		}
		return nil
	})
}

func (pokja *pokjaRepository) Find(model *[]entity.Pokja) error {
	tx := pokja.New()
	tx.Limit(100)
	if err := tx.Find(model).Error; err != nil {
		return err
	}

	return nil
}

func (pokja *pokjaRepository) UpdateStatus(id string, status bool) error {
	return pokja.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&entity.Pokja{}).Where("id = ?", id).Update("status", status)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
		return nil
	})
}

func (pokja *pokjaRepository) Delete(id string) error {
	return pokja.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Unscoped().Where("id = ?", id).Delete(&entity.Pokja{})
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
		return nil
	})
}

func (pokja *pokjaRepository) Count() int64 {
	var t int64

	tx := pokja.New()
	tx.Count(&t)

	return t
}

func (pokja *pokjaRepository) FindByID(id string) (entity.Pokja, error) {
	var model entity.Pokja
	tx := pokja.New()
	if err := tx.First(&model, "pokjas.id = ?", id).Error; err != nil {
		return entity.Pokja{}, err
	}
	return model, nil
}
