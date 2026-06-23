package service

import (
	"github.com/aliffatulmf/buku-tamu-apbj/entity"
	"github.com/aliffatulmf/buku-tamu-apbj/interface/repository"
	"github.com/aliffatulmf/buku-tamu-apbj/lib"
	"github.com/aliffatulmf/buku-tamu-apbj/request"
	"errors"
	"fmt"
	"time"
)

type PenyediaService struct {
	Repository repository.PenyediaRepository
}

func NewProviderService(repository repository.PenyediaRepository) PenyediaService {
	return PenyediaService{Repository: repository}
}

func (penyedia PenyediaService) Create(req request.PenyediaRequest) error {
	img, err := lib.SaveImageFile(req.Image)
	if err != nil {
		return err
	}

	model := entity.Provider{
		ProviderName: req.Name,
		Phone:        req.Telephone,
		Company:      req.Company,
		Description:  req.Description,
		Destination:  req.Destination,
		Consultation: req.Consultation,
		Pokja:        req.Pokja,
		Image:        img,
	}

	tx := penyedia.Repository.New()

	switch model.Destination {
	case "Advokasi":
		err = tx.Omit("Consultation", "Pokja").Create(&model).Error
	case "LPSE":
		err = tx.Omit("Pokja").Create(&model).Error
	case "POKJA":
		err = tx.Omit("Consultation").Create(&model).Error
	}

	if err != nil {
		return errors.New("ERROR: cannot create record.")
	}

	return nil
}

func (penyedia PenyediaService) Find(flt request.FilterQueryRequest) ([]entity.Provider, error) {
	var model []entity.Provider
	var et time.Time

	tx := penyedia.Repository.Find()
	tx.Order("updated_at desc")

	if flt.SBN != "" {
		arg := fmt.Sprintf("%%%s%%", flt.SBN)
		tx.Where("provider_name like ?", arg)
	}

	if flt.From != et && flt.To != et {
		tx.Where("created_at between ? and ?", flt.From, flt.To)
	}

	switch flt.Permission {
	case "allowed":
		tx.Where("verified = ?", true)
	case "notallowed":
		tx.Where("verified = ?", false)
	}

	if err := tx.Scan(&model).Error; err != nil {
		return []entity.Provider{}, err
	}

	return model, nil
}

func (penyedia PenyediaService) FindByID(id string) (entity.Provider, error) {
	var model entity.Provider

	if err := penyedia.Repository.FindByID(&model, id); err != nil {
		return model, err
	}
	return model, nil
}

func (penyedia PenyediaService) DeleteByID(id string) error {
	if err := penyedia.Repository.Delete(&entity.Provider{}, "id = ?", id); err != nil {
		return err
	}
	return nil
}

func (penyedia PenyediaService) UpdatePermission(id string) error {
	if err := penyedia.Repository.UpdatePermission(id); err != nil {
		return err
	}
	return nil
}
