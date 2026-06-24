package repository

import (
	"github.com/aliffatulmf/buku-tamu-apbj/internal/entity"

	"gorm.io/gorm"
)

type PokjaRepository struct {
	DB *gorm.DB
}

func NewPokjaRepository(db *gorm.DB) *PokjaRepository {
	return &PokjaRepository{DB: db}
}

func (pokja *PokjaRepository) New() *gorm.DB {
	tx := pokja.DB.Model(entity.Pokja{})
	tx.Session(&gorm.Session{
		QueryFields: true,
	})

	return tx
}

func (pokja *PokjaRepository) Create(model *entity.Pokja) error {
	return pokja.New().Create(model).Error
}

func (pokja *PokjaRepository) Find(model *[]entity.Pokja) error {
	tx := pokja.New()
	tx.Limit(100)
	if err := tx.Find(model).Error; err != nil {
		return err
	}

	return nil
}

func (pokja *PokjaRepository) UpdateStatus(id string, status bool) error {
	result := pokja.DB.Model(&entity.Pokja{}).Where("id = ?", id).Update("status", status)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (pokja *PokjaRepository) Delete(id string) error {
	result := pokja.DB.Unscoped().Where("id = ?", id).Delete(&entity.Pokja{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (pokja *PokjaRepository) Count() int64 {
	var t int64

	tx := pokja.New()
	tx.Count(&t)

	return t
}

func (pokja *PokjaRepository) FindByID(id string) (entity.Pokja, error) {
	var model entity.Pokja
	tx := pokja.New()
	if err := tx.First(&model, "pokjas.id = ?", id).Error; err != nil {
		return entity.Pokja{}, err
	}
	return model, nil
}
