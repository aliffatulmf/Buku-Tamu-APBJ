package service

import (
	"github.com/aliffatulmf/buku-tamu-apbj/entity"
	"github.com/aliffatulmf/buku-tamu-apbj/interface/repository"
	"github.com/aliffatulmf/buku-tamu-apbj/lib"
	"github.com/aliffatulmf/buku-tamu-apbj/request"
	"errors"
	"time"
)

type DashboardService struct {
	Pemda    repository.PemdaRepository
	Provider repository.PenyediaRepository
	Pokja    repository.PokjaRepository
	Instansi repository.InstansiRepository
}

type DType string

var PemdaType DType = "Pemda"
var PenyediaType DType = "Penyedia"

func NewDashboardServices(model DashboardService) DashboardService {
	return model
}

type DCnt struct {
	Pemda    int64
	Provider int64
	Pokja    int64
	Instansi int64
}

func (s DashboardService) DashboardCounts() DCnt {
	return DCnt{
		Pemda:    s.Pemda.Count(),
		Provider: s.Provider.Count(),
		Pokja:    s.Pokja.Count(),
		Instansi: s.Instansi.Count(),
	}
}

func (s DashboardService) FindPemda(start, end time.Time) []entity.TypePemdaAgency {
	var model []entity.TypePemdaAgency
	var et time.Time

	tx := s.Pemda.New()
	if start != et && end != et {
		tx.Where("pemdas.created_at between ? and ?", start.String(), end.String())
	}
	tx.Scan(&model)

	return model
}

func (s DashboardService) FindProvider(start, end time.Time) []entity.Provider {
	var model []entity.Provider
	var et time.Time

	tx := s.Provider.New()
	if start != et && end != et {
		tx.Where("created_at between ? and ?", start.String(), end.String())
	}
	tx.Scan(&model)

	return model
}

func (s DashboardService) GoExport(t DType, req request.ExportQuery) error {
	switch t {
	case PemdaType:
		res := s.FindPemda(req.From, req.To)
		if err := lib.NewExportPemda(res); err != nil {
			return err
		}
	case PenyediaType:
		res := s.FindProvider(req.From, req.To)
		if err := lib.NewExportPenyedia(res); err != nil {
			return err
		}
	default:
		return errors.New("type cannot be recognized")
	}

	return nil
}
