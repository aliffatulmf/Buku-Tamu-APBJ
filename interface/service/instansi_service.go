package service

import (
	"github.com/aliffatulmf/buku-tamu-apbj/entity"
	"github.com/aliffatulmf/buku-tamu-apbj/interface/repository"
	"github.com/aliffatulmf/buku-tamu-apbj/request"
	"fmt"
	"time"
)

type InstansiService struct {
	Repository repository.InstansiRepository
}

func NewAgencyService(repository repository.InstansiRepository) InstansiService {
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

func (instansi InstansiService) FindByID(id string) (entity.Agency, error) {
	var model entity.Agency

	if err := instansi.Repository.FindBy(&model, "id = ?", id); err != nil {
		return model, err
	}

	return model, nil
}

//FindByFilter SBN min 3 char or more
func (instansi InstansiService) FindByFilter(req request.FilterQueryRequest) ([]entity.Agency, error) {
	var model []entity.Agency
	var et time.Time

	tx := instansi.Repository.New()
	if len(req.SBN) > 2 {
		tx.Where("agencies.agency_name like ?", fmt.Sprintf("%%%s%%", req.SBN))
	}

	if req.From != et && req.To != et {
		tx.Where("agencies.created_at between ? and ?", req.From, req.To)
	}

	if err := tx.Scan(&model).Error; err != nil {
		return model, err
	}

	return model, nil
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
