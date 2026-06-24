package service

import (
	"errors"
	"strings"

	"github.com/aliffatulmf/buku-tamu-apbj/internal/entity"
	"github.com/aliffatulmf/buku-tamu-apbj/internal/io"
	"github.com/aliffatulmf/buku-tamu-apbj/internal/repository"
	"github.com/aliffatulmf/buku-tamu-apbj/request"
)

type PenyediaService struct {
	Repository   *repository.PenyediaRepository
	ImageStorage *io.ImageStorage
}

func NewProviderService(
	repository *repository.PenyediaRepository,
	imageStorage *io.ImageStorage,
) PenyediaService {
	return PenyediaService{
		Repository:   repository,
		ImageStorage: imageStorage,
	}
}

func (penyedia PenyediaService) Create(req request.PenyediaRequest) error {
	img, err := penyedia.ImageStorage.Save(req.Image)
	if err != nil {
		return err
	}

	model := entity.Provider{
		ProviderName: req.Name,
		Phone:        req.Telephone,
		Company:      req.Company,
		Description:  req.Description,
		Destination:  strings.ToLower(req.Destination),
		Consultation: req.Consultation,
		Pokja:        req.Pokja,
		Image:        img,
	}

	switch entity.DestinationType(model.Destination) {
	case entity.DestinationAdvokasi:
		err = penyedia.Repository.CreateWithOmit(&model, "Consultation", "Pokja")
	case entity.DestinationLPSE:
		err = penyedia.Repository.CreateWithOmit(&model, "Pokja")
	case entity.DestinationPokja:
		err = penyedia.Repository.CreateWithOmit(&model, "Consultation")
	default:
		err = penyedia.Repository.CreateWithOmit(&model)
	}

	if err != nil {
		return errors.New("error: cannot create record")
	}

	return nil
}

func (penyedia PenyediaService) Find(flt request.FilterQueryRequest) ([]entity.Provider, error) {
	return penyedia.Repository.FindWithFilter(flt.SBN, flt.From, flt.To, flt.Permission)
}

func (penyedia PenyediaService) FindByID(id uint) (entity.Provider, error) {
	var model entity.Provider

	if err := penyedia.Repository.FindByID(&model, id); err != nil {
		return model, err
	}
	return model, nil
}

func (penyedia PenyediaService) DeleteByID(id uint) error {
	if err := penyedia.Repository.Delete(&entity.Provider{}, "id = ?", id); err != nil {
		return err
	}
	return nil
}

func (penyedia PenyediaService) UpdatePermission(id uint) error {
	if err := penyedia.Repository.UpdatePermission(id); err != nil {
		return err
	}
	return nil
}
