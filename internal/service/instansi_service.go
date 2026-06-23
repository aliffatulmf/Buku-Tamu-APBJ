package service

import (
	"github.com/aliffatulmf/buku-tamu-apbj/internal/entity"
	"github.com/aliffatulmf/buku-tamu-apbj/request"
)

type InstansiService struct {
	Repository entity.InstansiRepository
}

func NewAgencyService(repository entity.InstansiRepository) InstansiService {
	return InstansiService{Repository: repository}
}

func (instansi InstansiService) Create(req request.InstansiRequest) error {
	model := entity.Agency{
		AgencyName:      req.Name,
		AgencyEmail:     req.Email,
		AgencyTelephone: req.Telephone,
	}

	if err := instansi.Repository.Create(&model); err != nil {
		return err
	}
	return nil
}

func (instansi InstansiService) Find() ([]entity.Agency, error) {
	var model []entity.Agency

	if err := instansi.Repository.Find(&model); err != nil {
		return model, err
	}

	return model, nil
}

func (instansi InstansiService) FindByID(id uint) (entity.Agency, error) {
	var model entity.Agency

	if err := instansi.Repository.FindBy(&model, "id = ?", id); err != nil {
		return model, err
	}

	return model, nil
}

func (instansi InstansiService) FindByFilter(req request.FilterQueryRequest) ([]entity.Agency, error) {
	return instansi.Repository.FindWithFilter(req.SBN, req.From, req.To)
}

func (instansi InstansiService) Update(req request.InstansiRequest) error {
	model := entity.Agency{
		AgencyName:      req.Name,
		AgencyEmail:     req.Email,
		AgencyTelephone: req.Telephone,
	}

	if err := instansi.Repository.Update(&model, "agencies.id = ?", req.ID); err != nil {
		return err
	}
	return nil
}
