package repository

import (
	"fmt"
	"time"

	"github.com/aliffatulmf/buku-tamu-apbj/internal/entity"

	"gorm.io/gorm"
)

type InstansiRepository struct {
	DB *gorm.DB
}

func NewAgencyRepository(db *gorm.DB) *InstansiRepository {
	return &InstansiRepository{DB: db}
}

func (instansi *InstansiRepository) New() *gorm.DB {
	tx := instansi.DB.Model(entity.Agency{})
	return tx.Session(&gorm.Session{
		QueryFields: true,
	})
}

func (instansi *InstansiRepository) Create(model *entity.Agency) error {
	return instansi.New().Create(model).Error
}

func (instansi *InstansiRepository) Find(model *[]entity.Agency) error {
	tx := instansi.New()
	if err := tx.Find(model).Error; err != nil {
		return err
	}

	return nil
}

func (instansi *InstansiRepository) FindBy(model *entity.Agency, conds ...interface{}) error {
	tx := instansi.New()
	if err := tx.First(model, conds...).Error; err != nil {
		return err
	}

	return nil
}

func (instansi *InstansiRepository) Update(model *entity.Agency, conds ...interface{}) error {
	return instansi.New().Where(conds).Updates(model).Error
}

func (instansi *InstansiRepository) DeleteBy(conds ...interface{}) error {
	return instansi.New().Delete(entity.Agency{}, conds...).Error
}

func (instansi *InstansiRepository) Count() int64 {
	var col int64

	tx := instansi.New()
	tx.Count(&col)
	return col
}

func (instansi *InstansiRepository) FindWithFilter(sbn string, from, to time.Time) ([]entity.Agency, error) {
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
