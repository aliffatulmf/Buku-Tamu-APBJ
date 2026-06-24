package io

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/aliffatulmf/buku-tamu-apbj/internal/entity"
	"github.com/xuri/excelize/v2"
)

type ExcelExporter struct {
	pemdaDir    string
	penyediaDir string
}

func NewExcelExporter() *ExcelExporter {
	wd, _ := os.Getwd()
	return &ExcelExporter{
		pemdaDir:    filepath.Join(wd, "Documents", "Pemda"),
		penyediaDir: filepath.Join(wd, "Documents", "Penyedia"),
	}
}

func (e *ExcelExporter) ExportPemda(records []entity.TypePemdaAgency) error {
	f := excelize.NewFile()
	sheet := "Sheet1"
	index := f.NewSheet(sheet)

	headers := []string{"NO", "Nama", "Telepon", "SKPD/OPD", "Nama Instansi", "Telepon Instansi", "Email Instansi", "Tujuan", "Konsultasi", "Pokja", "Status", "Tanggal"}
	for i, h := range headers {
		f.SetCellValue(sheet, string(rune('A'+i))+"1", h)
	}

	for idx, r := range records {
		row := idx + 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), idx+1)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), r.PemdaName)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), r.Phone)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), r.SkpdOpd)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), r.Agency.AgencyName)
		f.SetCellValue(sheet, fmt.Sprintf("F%d", row), r.Agency.AgencyTelephone)
		f.SetCellValue(sheet, fmt.Sprintf("G%d", row), r.Agency.AgencyEmail)
		f.SetCellValue(sheet, fmt.Sprintf("H%d", row), r.Destination)
		f.SetCellValue(sheet, fmt.Sprintf("I%d", row), r.Consultation)
		f.SetCellValue(sheet, fmt.Sprintf("J%d", row), r.Pokja)
		f.SetCellValue(sheet, fmt.Sprintf("K%d", row), statusLabel(r.Verified))
		f.SetCellValue(sheet, fmt.Sprintf("L%d", row), r.CreatedAt.Format("02 January 2006 15:04:05.000"))
	}

	f.SetActiveSheet(index)

	os.MkdirAll(e.pemdaDir, os.ModeDir)
	name := fmt.Sprintf("PEMDA %s.xlsx", time.Now().Format("2006_01_02-15_04_05"))
	return f.SaveAs(filepath.Join(e.pemdaDir, name))
}

func (e *ExcelExporter) ExportPenyedia(records []entity.Provider) error {
	f := excelize.NewFile()
	sheet := "Sheet1"
	index := f.NewSheet(sheet)

	headers := []string{"NO", "Nama", "Telepon", "Nama Perusahaan", "Keterangan", "Tujuan", "Konsultasi", "Pokja", "Status", "Tanggal"}
	for i, h := range headers {
		f.SetCellValue(sheet, string(rune('A'+i))+"1", h)
	}

	for idx, r := range records {
		row := idx + 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), idx+1)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), r.ProviderName)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), r.Phone)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), r.Company)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), r.Description)
		f.SetCellValue(sheet, fmt.Sprintf("F%d", row), r.Destination)
		f.SetCellValue(sheet, fmt.Sprintf("G%d", row), r.Consultation)
		f.SetCellValue(sheet, fmt.Sprintf("H%d", row), r.Pokja)
		f.SetCellValue(sheet, fmt.Sprintf("I%d", row), statusLabel(r.Verified))
		f.SetCellValue(sheet, fmt.Sprintf("J%d", row), r.CreatedAt.Format("02 January 2006 15:04:05.000"))
	}

	f.SetActiveSheet(index)

	os.MkdirAll(e.penyediaDir, os.ModeDir)
	name := fmt.Sprintf("PENYEDIA %s.xlsx", time.Now().Format("2006_01_02-15_04_05"))
	return f.SaveAs(filepath.Join(e.penyediaDir, name))
}

func statusLabel(verified bool) string {
	if verified {
		return "Diizinkan"
	}
	return "Tidak Diizinkan"
}
