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
		_ = f.SetCellValue(sheet, string(rune('A'+i))+"1", h)
	}

	for idx, r := range records {
		row := idx + 2
		_ = f.SetCellValue(sheet, fmt.Sprintf("A%d", row), idx+1)
		_ = f.SetCellValue(sheet, fmt.Sprintf("B%d", row), r.PemdaName)
		_ = f.SetCellValue(sheet, fmt.Sprintf("C%d", row), r.Phone)
		_ = f.SetCellValue(sheet, fmt.Sprintf("D%d", row), r.SkpdOpd)
		_ = f.SetCellValue(sheet, fmt.Sprintf("E%d", row), r.Agency.AgencyName)
		_ = f.SetCellValue(sheet, fmt.Sprintf("F%d", row), r.Agency.AgencyTelephone)
		_ = f.SetCellValue(sheet, fmt.Sprintf("G%d", row), r.Agency.AgencyEmail)
		_ = f.SetCellValue(sheet, fmt.Sprintf("H%d", row), r.Destination)
		_ = f.SetCellValue(sheet, fmt.Sprintf("I%d", row), r.Consultation)
		_ = f.SetCellValue(sheet, fmt.Sprintf("J%d", row), r.Pokja)
		_ = f.SetCellValue(sheet, fmt.Sprintf("K%d", row), statusLabel(r.Verified))
		_ = f.SetCellValue(sheet, fmt.Sprintf("L%d", row), r.CreatedAt.Format("02 January 2006 15:04:05.000"))
	}

	f.SetActiveSheet(index)

	_ = os.MkdirAll(e.pemdaDir, os.ModeDir)
	name := fmt.Sprintf("PEMDA %s.xlsx", time.Now().Format("2006_01_02-15_04_05"))
	return f.SaveAs(filepath.Join(e.pemdaDir, name))
}

func (e *ExcelExporter) ExportPenyedia(records []entity.Provider) error {
	f := excelize.NewFile()
	sheet := "Sheet1"
	index := f.NewSheet(sheet)

	headers := []string{"NO", "Nama", "Telepon", "Nama Perusahaan", "Keterangan", "Tujuan", "Konsultasi", "Pokja", "Status", "Tanggal"}
	for i, h := range headers {
		_ = f.SetCellValue(sheet, string(rune('A'+i))+"1", h)
	}

	for idx, r := range records {
		row := idx + 2
		_ = f.SetCellValue(sheet, fmt.Sprintf("A%d", row), idx+1)
		_ = f.SetCellValue(sheet, fmt.Sprintf("B%d", row), r.ProviderName)
		_ = f.SetCellValue(sheet, fmt.Sprintf("C%d", row), r.Phone)
		_ = f.SetCellValue(sheet, fmt.Sprintf("D%d", row), r.Company)
		_ = f.SetCellValue(sheet, fmt.Sprintf("E%d", row), r.Description)
		_ = f.SetCellValue(sheet, fmt.Sprintf("F%d", row), r.Destination)
		_ = f.SetCellValue(sheet, fmt.Sprintf("G%d", row), r.Consultation)
		_ = f.SetCellValue(sheet, fmt.Sprintf("H%d", row), r.Pokja)
		_ = f.SetCellValue(sheet, fmt.Sprintf("I%d", row), statusLabel(r.Verified))
		_ = f.SetCellValue(sheet, fmt.Sprintf("J%d", row), r.CreatedAt.Format("02 January 2006 15:04:05.000"))
	}

	f.SetActiveSheet(index)

	_ = os.MkdirAll(e.penyediaDir, os.ModeDir)
	name := fmt.Sprintf("PENYEDIA %s.xlsx", time.Now().Format("2006_01_02-15_04_05"))
	return f.SaveAs(filepath.Join(e.penyediaDir, name))
}

func statusLabel(verified bool) string {
	if verified {
		return "Diizinkan"
	}
	return "Tidak Diizinkan"
}
