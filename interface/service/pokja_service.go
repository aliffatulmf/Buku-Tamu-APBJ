package service

import (
	"github.com/aliffatulmf/buku-tamu-apbj/entity"
	"github.com/aliffatulmf/buku-tamu-apbj/interface/repository"
	"github.com/aliffatulmf/buku-tamu-apbj/request"
)

type PokjaService struct {
	Repository repository.PokjaRepository
}

func NewPokjaService(repository repository.PokjaRepository) PokjaService {
	return PokjaService{Repository: repository}
}

func (pokja PokjaService) Find() ([]entity.Pokja, error) {
	var model []entity.Pokja

	if err := pokja.Repository.Find(&model); err != nil {
		return model, err
	}

	return model, nil
}

func (pokja PokjaService) FindByID(id string) (entity.Pokja, error) {
	var model entity.Pokja

	tx := pokja.Repository.New()
	if err := tx.First(&model, "pokjas.id = ?", id).Error; err != nil {
		return model, err
	}

	return model, nil
}

func (pokja PokjaService) UpdateStatus(req request.PokjaRequest) error {
	if err := pokja.Repository.UpdateStatus(req.ID, req.UpdateStatus); err != nil {
		return err
	}

	return nil
}

func (pokja PokjaService) Create(req request.PokjaCreateRequest) error {
	model := entity.Pokja{
		PokjaName:     req.Name,
		Status:        true,
		DestinationID: 3,
	}

	if err := pokja.Repository.Create(&model); err != nil {
		return err
	}
	return nil
}

func (pokja PokjaService) Delete(id string) error {
	if err := pokja.Repository.Delete(id); err != nil {
		return err
	}

	return nil
}
