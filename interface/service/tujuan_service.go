package service

import (
	"github.com/aliffatulmf/buku-tamu-apbj/entity"
	"github.com/aliffatulmf/buku-tamu-apbj/interface/repository"
)

type TujuanService struct {
	Repository repository.TujuanRepository
}

func NewTujuanService(repository repository.TujuanRepository) TujuanService {
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
