package repository

import (
	"gorm.io/gorm"
)

type TujuanRepository struct {
	DB *gorm.DB
}

func NewTujuanRepository(db *gorm.DB) *TujuanRepository {
	return &TujuanRepository{DB: db}
}

func (tujuan *TujuanRepository) New(model interface{}) *gorm.DB {
	tx := tujuan.DB.Model(model)
	return tx.Session(&gorm.Session{
		QueryFields: true,
	})
}
