package service

import (
	"github.com/aliffatulmf/buku-tamu-apbj/internal/entity"
)

type TujuanService struct {
	Repository entity.TujuanRepository
}

func NewTujuanService(repository entity.TujuanRepository) TujuanService {
	return TujuanService{Repository: repository}
}

func (tujuan TujuanService) FindTujuan() ([]entity.Destination, error) {
	var model []entity.Destination

	tx := tujuan.Repository.New(model)
	if err := tx.Find(&model).Error; err != nil {
		return model, err
	}
	return model, nil
}

func (tujuan TujuanService) FindConsultation() ([]entity.Consultation, error) {
	var model []entity.Consultation

	tx := tujuan.Repository.New(model)
	if err := tx.Find(&model).Error; err != nil {
		return model, err
	}
	return model, nil
}

func (tujuan TujuanService) FindPokja() ([]entity.Pokja, error) {
	var model []entity.Pokja

	tx := tujuan.Repository.New(model)
	if err := tx.Find(&model, "status = ?", true).Error; err != nil {
		return model, err
	}
	return model, nil
}
