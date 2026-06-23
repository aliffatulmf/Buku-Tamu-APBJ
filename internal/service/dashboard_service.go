package service

import (
	"errors"
	"time"

	"github.com/aliffatulmf/buku-tamu-apbj/internal/entity"
	"github.com/aliffatulmf/buku-tamu-apbj/request"
)

type DashboardService struct {
	Pemda    entity.PemdaRepository
	Provider entity.PenyediaRepository
	Pokja    entity.PokjaRepository
	Instansi entity.InstansiRepository
	Exporter DataExporter
}

type DType string

var PemdaType DType = "Pemda"
var PenyediaType DType = "Penyedia"

func NewDashboardServices(
	pemda entity.PemdaRepository,
	provider entity.PenyediaRepository,
	pokja entity.PokjaRepository,
	instansi entity.InstansiRepository,
	exporter DataExporter,
) DashboardService {
	return DashboardService{
		Pemda:    pemda,
		Provider: provider,
		Pokja:    pokja,
		Instansi: instansi,
		Exporter: exporter,
	}
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
	model, _ := s.Pemda.FindByDateRange(start, end)
	return model
}

func (s DashboardService) FindProvider(start, end time.Time) []entity.Provider {
	model, _ := s.Provider.FindByDateRange(start, end)
	return model
}

func (s DashboardService) GoExport(t DType, req request.ExportQuery) error {
	switch t {
	case PemdaType:
		res := s.FindPemda(req.From, req.To)
		if err := s.Exporter.ExportPemda(res); err != nil {
			return err
		}
	case PenyediaType:
		res := s.FindProvider(req.From, req.To)
		if err := s.Exporter.ExportPenyedia(res); err != nil {
			return err
		}
	default:
		return errors.New("type cannot be recognized")
	}

	return nil
}
