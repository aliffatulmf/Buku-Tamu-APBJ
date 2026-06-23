package service

import "github.com/aliffatulmf/buku-tamu-apbj/internal/entity"

type ImageStorage interface {
	Save(base64Data string) (string, error)
}

type DataExporter interface {
	ExportPemda(records []entity.TypePemdaAgency) error
	ExportPenyedia(records []entity.Provider) error
}
