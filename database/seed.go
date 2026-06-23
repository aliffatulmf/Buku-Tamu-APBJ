package database

import (
	"github.com/aliffatulmf/buku-tamu-apbj/internal/entity"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	seedDestinations(db)
	seedConsultations(db)
	seedPokjas(db)
	seedAgencies(db)
}

func seedDestinations(db *gorm.DB) {
	var count int64
	db.Model(&entity.Destination{}).Count(&count)
	if count > 0 {
		return
	}

	destinations := []entity.Destination{
		{ID: "advokasi", DestinationName: "Advokasi"},
		{ID: "lpse", DestinationName: "LPSE"},
		{ID: "pokja", DestinationName: "POKJA"},
	}
	db.Create(&destinations)
}

func seedConsultations(db *gorm.DB) {
	var count int64
	db.Model(&entity.Consultation{}).Count(&count)
	if count > 0 {
		return
	}

	consultations := []entity.Consultation{
		{ID: "konsultasi_spse", ConsultationName: "Konsultasi SPSE", DestinationID: "lpse"},
		{ID: "konsultasi_sirup", ConsultationName: "Konsultasi Sirup", DestinationID: "lpse"},
		{ID: "konsultasi_blangkon_jateng", ConsultationName: "Konsultasi Blangkon Jateng", DestinationID: "lpse"},
	}
	db.Create(&consultations)
}

func seedPokjas(db *gorm.DB) {
	var count int64
	db.Model(&entity.Pokja{}).Count(&count)
	if count > 0 {
		return
	}

	pokjas := []entity.Pokja{
		{ID: "pokja_pemilihan_1a", PokjaName: "Pokja Pemilihan 1a", Status: true, DestinationID: "pokja"},
		{ID: "pokja_pemilihan_2a", PokjaName: "Pokja Pemilihan 2a", Status: true, DestinationID: "pokja"},
		{ID: "pokja_pemilihan_3a", PokjaName: "Pokja Pemilihan 3a", Status: true, DestinationID: "pokja"},
		{ID: "pokja_pemilihan_4a", PokjaName: "Pokja Pemilihan 4a", Status: true, DestinationID: "pokja"},
		{ID: "pokja_pemilihan_5a", PokjaName: "Pokja Pemilihan 5a", Status: true, DestinationID: "pokja"},
		{ID: "pokja_pemilihan_6a", PokjaName: "Pokja Pemilihan 6a", Status: true, DestinationID: "pokja"},
		{ID: "pokja_pemilihan_7a", PokjaName: "Pokja Pemilihan 7a", Status: true, DestinationID: "pokja"},
		{ID: "pokja_pemilihan_8a", PokjaName: "Pokja Pemilihan 8a", Status: true, DestinationID: "pokja"},
		{ID: "pokja_pemilihan_9a", PokjaName: "Pokja Pemilihan 9a", Status: true, DestinationID: "pokja"},
		{ID: "pokja_pemilihan_10a", PokjaName: "Pokja Pemilihan 10a", Status: true, DestinationID: "pokja"},
		{ID: "pokja_pemilihan_11a", PokjaName: "Pokja Pemilihan 11a", Status: true, DestinationID: "pokja"},
		{ID: "pokja_pemilihan_12a", PokjaName: "Pokja Pemilihan 12a", Status: true, DestinationID: "pokja"},
		{ID: "pokja_pemilihan_13a", PokjaName: "Pokja Pemilihan 13a", Status: true, DestinationID: "pokja"},
	}
	db.Create(&pokjas)
}

func seedAgencies(db *gorm.DB) {
	var count int64
	db.Model(&entity.Agency{}).Count(&count)
	if count > 0 {
		return
	}

	agencies := []entity.Agency{
		{AgencyName: "Pemerintah Provinsi Jawa Tengah"},
		{AgencyName: "Pemerintah Kabupaten Cilacap"},
		{AgencyName: "Pemerintah Kabupaten Banyumas"},
		{AgencyName: "Pemerintah Kabupaten Purbalingga"},
		{AgencyName: "Pemerintah Kabupaten Banjarnegara"},
		{AgencyName: "Pemerintah Kabupaten Kebumen"},
		{AgencyName: "Pemerintah Kabupaten Purworejo"},
		{AgencyName: "Pemerintah Kabupaten Wonosobo"},
		{AgencyName: "Pemerintah Kabupaten Magelang"},
		{AgencyName: "Pemerintah Kabupaten Boyolali"},
		{AgencyName: "Pemerintah Kabupaten Klaten"},
		{AgencyName: "Pemerintah Kabupaten Sukoharjo"},
		{AgencyName: "Pemerintah Kabupaten Wonogiri"},
		{AgencyName: "Pemerintah Kabupaten Karanganyar"},
		{AgencyName: "Pemerintah Kabupaten Seragen"},
		{AgencyName: "Pemerintah Kabupaten Grobogan"},
		{AgencyName: "Pemerintah Kabupaten Blora"},
		{AgencyName: "Pemerintah Kabupaten Rembang"},
		{AgencyName: "Pemerintah Kabupaten Pati"},
		{AgencyName: "Pemerintah Kabupaten Kudus"},
		{AgencyName: "Pemerintah Kabupaten Jepara"},
		{AgencyName: "Pemerintah Kabupaten Demak"},
		{AgencyName: "Pemerintah Kabupaten Semarang"},
		{AgencyName: "Pemerintah Kabupaten Temanggung"},
		{AgencyName: "Pemerintah Kabupaten Kendal"},
		{AgencyName: "Pemerintah Kabupaten Batang"},
		{AgencyName: "Pemerintah Kabupaten Pekalongan"},
		{AgencyName: "Pemerintah Kabupaten Pemalang"},
		{AgencyName: "Pemerintah Kabupaten Tegal"},
		{AgencyName: "Pemerintah Kabupaten Brebes"},
		{AgencyName: "Pemerintah Kota Magelang"},
		{AgencyName: "Pemerintah Kota Surakarta"},
		{AgencyName: "Pemerintah Kota Salatiga"},
		{AgencyName: "Pemerintah Kota Semarang"},
		{AgencyName: "Pemerintah Kota Pekalongan"},
		{AgencyName: "Pemerintah Kota Tegal"},
	}
	db.Create(&agencies)
}
