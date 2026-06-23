package repository

import (
	"github.com/aliffatulmf/buku-tamu-apbj/entity"

	"gorm.io/gorm"
)

type PokjaRepository struct {
	DB *gorm.DB
}

func NewPokjaRepository(db *gorm.DB) PokjaRepository {
	return PokjaRepository{DB: db}
}

func (pokja PokjaRepository) New() *gorm.DB {
	tx := pokja.DB.Model(entity.Pokja{})
	tx.Session(&gorm.Session{
		QueryFields: true,
	})

	return tx
}

func (pokja PokjaRepository) Create(model *entity.Pokja) error {
	tx := pokja.New()
	return tx.Transaction(func(trx *gorm.DB) error {
		if err := trx.Create(model).Error; err != nil {
			return err
		}
		return nil
	})
}

func (pokja PokjaRepository) Find(model *[]entity.Pokja) error {
	tx := pokja.New()
	tx.Limit(100)
	if err := tx.Find(model).Error; err != nil {
		return err
	}

	return nil
}

func (pokja PokjaRepository) UpdateStatus(id string, status bool) error {
	return pokja.DB.Transaction(func(tx *gorm.DB) error {
		stx := tx.Model(entity.Pokja{})
		stx.Where("id = ?", id)
		stx.First(&entity.Pokja{})
		stx.Update("status", status)

		if err := stx.Error; err != nil {
			return err
		}
		return nil
	})
}

func (pokja PokjaRepository) Delete(id string) error {
	var model entity.Pokja
	return pokja.DB.Transaction(func(tx *gorm.DB) error {
		stx := tx.Model(entity.Pokja{})
		stx.Where("id = ?", id)
		stx.Unscoped()
		stx.Delete(&model)

		if err := stx.Error; err != nil {
			return err
		}
		return nil
	})
}

func (pokja PokjaRepository) Count() int64 {
	var t int64

	tx := pokja.New()
	tx.Count(&t)

	return t
}
