package repository

import (
	"fmt"
	"time"

	"github.com/aliffatulmf/buku-tamu-apbj/internal/entity"

	"gorm.io/gorm"
)

type instansiRepository struct {
	DB *gorm.DB
}

func NewAgencyRepository(db *gorm.DB) *instansiRepository {
	return &instansiRepository{DB: db}
}

func (instansi *instansiRepository) New() *gorm.DB {
	tx := instansi.DB.Model(entity.Agency{})
	return tx.Session(&gorm.Session{
		QueryFields: true,
	})
}

func (instansi *instansiRepository) Create(model *entity.Agency) error {
	return instansi.New().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(model).Error; err != nil {
			return err
		}
		return nil
	})
}

func (instansi *instansiRepository) Find(model *[]entity.Agency) error {
	tx := instansi.New()
	if err := tx.Find(model).Error; err != nil {
		return err
	}

	return nil
}

func (instansi *instansiRepository) FindBy(model *entity.Agency, conds ...interface{}) error {
	tx := instansi.New()
	if err := tx.First(model, conds...).Error; err != nil {
		return err
	}

	return nil
}

func (instansi *instansiRepository) Update(model *entity.Agency, conds ...interface{}) error {
	return instansi.New().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where(conds).Updates(model).Error; err != nil {
			return err
		}

		return nil
	})
}

func (instansi *instansiRepository) DeleteBy(conds ...interface{}) error {
	return instansi.New().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(entity.Agency{}, conds...).Error; err != nil {
			return err
		}

		return nil
	})
}

func (instansi *instansiRepository) Count() int64 {
	var col int64

	tx := instansi.New()
	tx.Count(&col)
	return col
}

func (instansi *instansiRepository) FindWithFilter(sbn string, from, to time.Time) ([]entity.Agency, error) {
	var model []entity.Agency
	tx := instansi.New()

	if len(sbn) > 2 {
		tx.Where("agencies.agency_name like ?", fmt.Sprintf("%%%s%%", sbn))
	}

	if !from.IsZero() && !to.IsZero() {
		tx.Where("agencies.created_at between ? and ?", from, to)
	}

	if err := tx.Scan(&model).Error; err != nil {
		return nil, err
	}
	return model, nil
}
