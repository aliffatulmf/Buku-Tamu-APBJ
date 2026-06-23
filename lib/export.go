package lib

import (
	"github.com/aliffatulmf/buku-tamu-apbj/entity"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type ExportPemdaAgency struct {
	ID uint
	entity.Pemda
	Agency struct {
		AgencyName      string
		AgencyEmail     string
		AgencyTelephone string
	} `gorm:"embedded"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

var WorkingDir, _ = os.Getwd()

func NewExportPemda(records []entity.TypePemdaAgency) error {
	f := excelize.NewFile()
	index := f.NewSheet("Sheet1")
	header := []string{"NO", "Nama", "Telepon", "SKPD/OPD", "Nama Instansi", "Telepon Instansi", "Email Instansi", "Tujuan", "Konsultasi", "Pokja", "Status", "Tanggal"}
	sheet := "Sheet1"

	for i := 1; i < len(header)+1; i++ {
		f.SetCellValue(sheet, fmt.Sprintf("%s1", string(64+i)), header[i-1])
	}

	for idx, rec := range records {
		f.SetCellValue(sheet, fmt.Sprintf("A%d", idx+2), idx+1)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", idx+2), rec.PemdaName)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", idx+2), rec.Phone)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", idx+2), rec.SkpdOpd)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", idx+2), rec.Agency.AgencyName)
		f.SetCellValue(sheet, fmt.Sprintf("F%d", idx+2), rec.Agency.AgencyTelephone)
		f.SetCellValue(sheet, fmt.Sprintf("G%d", idx+2), rec.Agency.AgencyEmail)
		f.SetCellValue(sheet, fmt.Sprintf("H%d", idx+2), rec.Destination)
		f.SetCellValue(sheet, fmt.Sprintf("I%d", idx+2), rec.Consultation)
		f.SetCellValue(sheet, fmt.Sprintf("J%d", idx+2), rec.Pokja)
		if rec.Verified {
			f.SetCellValue(sheet, fmt.Sprintf("K%d", idx+2), "Diizinkan")
		} else {
			f.SetCellValue(sheet, fmt.Sprintf("K%d", idx+2), "Tidak Diizinkan")
		}
		f.SetCellValue(sheet, fmt.Sprintf("L%d", idx+2), rec.CreatedAt.Format("02 January 2006 15:04:05.000"))
	}

	f.SetActiveSheet(index)

	nf := fmt.Sprintf("PEMDA %s.xlsx", time.Now().Format("2006_01_02-15_04_05"))
	fp := filepath.Join(WorkingDir, "Documents/Pemda", nf)

	if err := f.SaveAs(fp); err != nil {
		return err
	}

	return nil
}

func NewExportPenyedia(records []entity.Provider) error {
	f := excelize.NewFile()
	index := f.NewSheet("Sheet1")
	header := []string{"NO", "Nama", "Telepon", "Nama Perusahaan", "Keterangan", "Tujuan", "Konsultasi", "Pokja", "Status", "Tanggal"}
	sheet := "Sheet1"

	for i := 1; i < len(header)+1; i++ {
		f.SetCellValue(sheet, fmt.Sprintf("%s1", string(64+1)), header[i-1])
	}

	for idx, rec := range records {
		f.SetCellValue(sheet, fmt.Sprintf("A%d", idx+2), idx+1)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", idx+2), rec.ProviderName)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", idx+2), rec.Phone)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", idx+2), rec.Company)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", idx+2), rec.Description)
		f.SetCellValue(sheet, fmt.Sprintf("F%d", idx+2), rec.Destination)
		f.SetCellValue(sheet, fmt.Sprintf("G%d", idx+2), rec.Consultation)
		f.SetCellValue(sheet, fmt.Sprintf("H%d", idx+2), rec.Pokja)
		if rec.Verified {
			f.SetCellValue(sheet, fmt.Sprintf("I%d", idx+2), "Diizinkan")
		} else {
			f.SetCellValue(sheet, fmt.Sprintf("I%d", idx+2), "Tidak Diizinkan")
		}
		f.SetCellValue(sheet, fmt.Sprintf("J%d", idx+2), rec.CreatedAt.Format("02 January 2006 15:04:05.000"))
	}

	f.SetActiveSheet(index)

	nf := fmt.Sprintf("PENYEDIA %s.xlsx", time.Now().Format("2006_01_02-15_04_05"))
	fp := filepath.Join(WorkingDir, "Documents/Penyedia", nf)

	if err := f.SaveAs(fp); err != nil {
		return err
	}

	return nil
}
