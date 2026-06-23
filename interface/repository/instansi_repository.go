package repository

import (
	"github.com/aliffatulmf/buku-tamu-apbj/entity"

	"gorm.io/gorm"
)

type InstansiRepository struct {
	DB *gorm.DB
}

func NewAgencyRepository(gorm *gorm.DB) InstansiRepository {
	return InstansiRepository{DB: gorm}
}

func (instansi InstansiRepository) New() *gorm.DB {
	tx := instansi.DB.Model(entity.Agency{})
	return tx.Session(&gorm.Session{
		QueryFields: true,
	})
}

func (instansi InstansiRepository) Create(model *entity.Agency) error {
	return instansi.New().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(model).Error; err != nil {
			return err
		}
		return nil
	})
}

func (instansi InstansiRepository) Find(model *[]entity.Agency) error {
	tx := instansi.New()
	if err := tx.Find(model).Error; err != nil {
		return err
	}

	return nil
}

func (instansi InstansiRepository) FindBy(model *entity.Agency, conds ...interface{}) error {
	tx := instansi.New()
	if err := tx.First(model, conds...).Error; err != nil {
		return err
	}

	return nil
}

func (instansi InstansiRepository) Update(model *entity.Agency, conds ...interface{}) error {
	return instansi.New().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where(conds).Updates(model).Error; err != nil {
			return err
		}

		return nil
	})
}

func (instansi InstansiRepository) DeleteBy(conds ...interface{}) error {
	return instansi.New().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(entity.Agency{}, conds...).Error; err != nil {
			return err
		}

		return nil
	})
}

func (instansi InstansiRepository) Count() int64 {
	var col int64

	tx := instansi.New()
	tx.Count(&col)
	return col
}
