package service

import (
	"github.com/aliffatulmf/buku-tamu-apbj/internal/entity"
	"github.com/aliffatulmf/buku-tamu-apbj/internal/repository"
	"github.com/aliffatulmf/buku-tamu-apbj/request"
)

type PokjaService struct {
	Repository *repository.PokjaRepository
}

func NewPokjaService(repository *repository.PokjaRepository) PokjaService {
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
	return pokja.Repository.FindByID(id)
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
		DestinationID: string(entity.DestinationPokja),
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
