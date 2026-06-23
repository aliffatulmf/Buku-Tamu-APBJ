package repository

import (
	"gorm.io/gorm"
)

type tujuanRepository struct {
	DB *gorm.DB
}

func NewTujuanRepository(db *gorm.DB) *tujuanRepository {
	return &tujuanRepository{DB: db}
}

func (tujuan *tujuanRepository) New(model interface{}) *gorm.DB {
	tx := tujuan.DB.Model(model)
	return tx.Session(&gorm.Session{
		QueryFields: true,
	})
}
