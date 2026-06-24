package service

import (
	"strings"

	"github.com/aliffatulmf/buku-tamu-apbj/internal/entity"
	"github.com/aliffatulmf/buku-tamu-apbj/internal/io"
	"github.com/aliffatulmf/buku-tamu-apbj/internal/repository"
	"github.com/aliffatulmf/buku-tamu-apbj/request"
)

type PemdaService struct {
	Repository   *repository.PemdaRepository
	Instansi     *repository.InstansiRepository
	ImageStorage *io.ImageStorage
}

func NewPemdaService(
	repository *repository.PemdaRepository,
	instansi *repository.InstansiRepository,
	imageStorage *io.ImageStorage,
) PemdaService {
	return PemdaService{
		Repository:   repository,
		Instansi:     instansi,
		ImageStorage: imageStorage,
	}
}

func (pemda PemdaService) Find(flt request.FilterQueryRequest) ([]entity.TypePemdaAgency, error) {
	return pemda.Repository.FindWithFilter(flt.SBN, flt.From, flt.To, flt.Permission)
}

func (pemda PemdaService) Create(req request.PemdaRequest) error {
	res, err := pemda.ImageStorage.Save(req.Image)
	if err != nil {
		return err
	}

	model := entity.Pemda{
		PemdaName:    req.Name,
		Phone:        req.Telephone,
		AgencyID:     req.Agency,
		SkpdOpd:      req.SkpdOpd,
		Destination:  strings.ToLower(req.Destination),
		Consultation: req.Consultation,
		Pokja:        req.Pokja,
		Image:        res,
	}

	switch entity.DestinationType(model.Destination) {
	case entity.DestinationAdvokasi:
		err = pemda.Repository.CreateWithOmit(&model, "Consultation", "Pokja")
	case entity.DestinationLPSE:
		err = pemda.Repository.CreateWithOmit(&model, "Pokja")
	case entity.DestinationPokja:
		err = pemda.Repository.CreateWithOmit(&model, "Consultation")
	default:
		err = pemda.Repository.CreateWithOmit(&model)
	}

	if err != nil {
		return err
	}
	return nil
}

func (pemda PemdaService) FindByID(id uint) (entity.TypePemdaAgency, error) {
	return pemda.Repository.FindByID(id)
}

func (pemda PemdaService) DeleteByID(id uint) error {
	return pemda.Repository.Delete(id)
}
func (pemda PemdaService) UpdatePermission(id uint) error {
	return pemda.Repository.UpdatePermission(id)
}
