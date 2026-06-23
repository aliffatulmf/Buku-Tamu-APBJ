package request

import "time"

type PemdaRequest struct {
	Name         string `form:"nama" binding:"required"`
	Telephone    string `form:"telepon" binding:"required,numeric"`
	Agency       uint   `form:"instansi" binding:"required,numeric"`
	SkpdOpd      string `form:"skpd_opd" binding:"required"`
	Destination  string `form:"tujuan"  binding:"required"`
	Consultation string `form:"konsultasi" binding:"required_if=Destination 2"`
	Pokja        string `form:"pokja" binding:"required_if=Destination 3"`
	Image        string `form:"image" binding:"required"`
}

type InstansiRequest struct {
	ID        uint
	Name      string `form:"name" binding:"required"`
	Email     string `form:"email"`
	Telephone string `form:"telephone"`
}

type PenyediaRequest struct {
	Name         string `form:"nama" binding:"required"`
	Telephone    string `form:"telepon" binding:"required,numeric"`
	Company      string `form:"company" binding:"required"`
	Description  string `form:"keterangan"`
	Destination  string `form:"tujuan"  binding:"required"`
	Consultation string `form:"konsultasi" binding:"required_if=Destination 2"`
	Pokja        string `form:"pokja" binding:"required_if=Destination 3"`
	Image        string `form:"image" binding:"required"`
}

type PokjaCreateRequest struct {
	Name string `form:"nama" binding:"required"`
}

type FilterQueryRequest struct {
	QF         string    `form:"qf"`
	SBN        string    `form:"sbn"`
	Permission string    `form:"permission"`
	From       time.Time `form:"from" time_format:"2006-01-02"`
	To         time.Time `form:"to" time_format:"2006-01-02"`
}

type ExportQuery struct {
	QF    string    `form:"qf"`
	Table string    `form:"tabel"`
	From  time.Time `form:"from" time_format:"2006-01-02"`
	To    time.Time `form:"to" time_format:"2006-01-02"`
}

type PokjaRequest struct {
	ID            string `form:"id"`
	CurrentStatus bool   `form:"cstat"`
	UpdateStatus  bool   `form:"ustat"`
}
