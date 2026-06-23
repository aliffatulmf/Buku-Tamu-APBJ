package service

import (
	"github.com/aliffatulmf/buku-tamu-apbj/entity"
	"github.com/aliffatulmf/buku-tamu-apbj/interface/repository"
	"github.com/aliffatulmf/buku-tamu-apbj/lib"
	"github.com/aliffatulmf/buku-tamu-apbj/request"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type PemdaService struct {
	Repository repository.PemdaRepository
	Instansi   repository.InstansiRepository
}

func NewPemdaService(repository repository.PemdaRepository, instansi repository.InstansiRepository) PemdaService {
	return PemdaService{
		Repository: repository,
		Instansi:   instansi,
	}
}

func (pemda PemdaService) Find(flt request.FilterQueryRequest) ([]entity.TypePemdaAgency, error) {
	var model []entity.TypePemdaAgency
	var et time.Time

	tx := pemda.Repository.New()
	tx.Order("updated_at desc")

	if flt.SBN != "" {
		arg := fmt.Sprintf("%%%s%%", flt.SBN)
		tx.Where("pemda_name LIKE ?", arg)
	}

	if flt.From != et && flt.To != et {
		tx.Where("pemdas.created_at BETWEEN ? AND ?", flt.From, flt.To)
	}

	switch flt.Permission {
	case "allowed":
		tx.Where("verified = ?", true)
	case "notallowed":
		tx.Where("verified = ?", false)
	}

	if err := tx.Scan(&model).Error; err != nil {
		return []entity.TypePemdaAgency{}, err
	}

	return model, nil
}

func (pemda PemdaService) Create(req request.PemdaRequest) error {
	res, err := lib.SaveImageFile(req.Image)
	if err != nil {
		return err
	}

	model := entity.Pemda{
		PemdaName:    req.Name,
		Phone:        req.Telephone,
		AgencyID:     req.Agency,
		SkpdOpd:      req.SkpdOpd,
		Destination:  req.Destination,
		Consultation: req.Consultation,
		Pokja:        req.Pokja,
		Image:        res,
	}

	tx := pemda.Repository.New()

	switch req.Destination {
	case "Advokasi":
		err = tx.Omit("Consultation", "Pokja").Create(&model).Error
	case "LPSE":
		err = tx.Omit("Pokja").Create(&model).Error
	case "POKJA":
		err = tx.Omit("Consultation").Create(&model).Error
	}

	if err != nil {
		return err
	}
	return nil
}

func (pemda PemdaService) FindByID(id string) (entity.TypePemdaAgency, error) {
	var model entity.TypePemdaAgency

	tx := pemda.Repository.New()
	tx.Where("pemdas.id = ?", id)

	if err := tx.Scan(&model).Error; err != nil {
		return entity.TypePemdaAgency{}, err
	}

	if model.ID < 1 {
		return entity.TypePemdaAgency{}, gorm.ErrRecordNotFound
	}

	return model, nil
}

func (pemda PemdaService) DeleteByID(id string) error {
	return pemda.Repository.Delete(id)
}
func (pemda PemdaService) UpdatePermission(id string) error {
	return pemda.Repository.UpdatePermission(id)
}
