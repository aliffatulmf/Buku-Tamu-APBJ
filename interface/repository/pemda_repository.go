package repository

import (
	"github.com/aliffatulmf/buku-tamu-apbj/entity"

	"gorm.io/gorm"
)

type PemdaRepository struct {
	DB *gorm.DB
}

type E *gorm.DB

func NewPemdaRepository(db *gorm.DB) PemdaRepository {
	return PemdaRepository{DB: db}
}

func (pemda PemdaRepository) New() *gorm.DB {
	tx := pemda.DB.Table("pemdas")
	tx.Joins("inner join agencies on agencies.id = pemdas.agency_id")
	tx.Where("pemdas.deleted_at is null")
	// tx.Session(&gorm.Session{
	// 	QueryFields: true,
	// })

	return tx
}

func (pemda PemdaRepository) Find() *gorm.DB {
	return pemda.New().Limit(100)
}

func (pemda PemdaRepository) FindOne(model *entity.TypePemdaAgency) error {
	tx := pemda.New()
	tx.Scan(model)

	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func (pemda PemdaRepository) Create(model *entity.Pemda) error {
	tx := pemda.New()
	tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(model).Error; err != nil {
			return err
		}
		return nil
	})

	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func (pemda PemdaRepository) Delete(id string) error {
	mtx := pemda.New()
	mtx.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(entity.Pemda{}, id).Error; err != nil {
			return err
		}
		return nil
	})

	if err := mtx.Error; err != nil {
		return err
	}
	return nil
}

func (pemda PemdaRepository) UpdatePermission(id string) error {
	var model entity.Pemda

	return pemda.New().Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&model, id).Updates(entity.Pemda{Verified: true}).Error; err != nil {
			return err
		}
		return nil
	})
}

func (pemda PemdaRepository) Count() int64 {
	var t int64

	tx := pemda.New()
	tx.Count(&t)

	return t
}
