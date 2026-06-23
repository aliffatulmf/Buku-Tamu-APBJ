package handler

import (
	"github.com/aliffatulmf/buku-tamu-apbj/interface/service"
	"github.com/aliffatulmf/buku-tamu-apbj/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PemdaHandler struct {
	Service  service.PemdaService
	Tujuan   service.TujuanService
	Instansi service.InstansiService
}

func NewPemdaHandler(service service.PemdaService, tujuan service.TujuanService, instansi service.InstansiService) PemdaHandler {
	return PemdaHandler{
		Service:  service,
		Tujuan:   tujuan,
		Instansi: instansi,
	}
}

func (pemda PemdaHandler) PemdaIndex(ctx *gin.Context) {
	dest, err := pemda.Tujuan.FindTujuan()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Destination error", err.Error())
		return
	}

	pok, err := pemda.Tujuan.FindPokja()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Pokja error", err.Error())
		return
	}

	con, err := pemda.Tujuan.FindConsultation()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Consultation error", err.Error())
		return
	}

	ins, err := pemda.Instansi.Find()
	if err != nil {
		ctx.String(http.StatusOK, "Got error")
	}

	ctx.HTML(http.StatusOK, "registrasi_pemda.html", gin.H{
		"data": gin.H{
			"destination":  dest,
			"pokja":        pok,
			"consultation": con,
			"instansi":     ins,
		},
	})
}

func (pemda PemdaHandler) PemdaCreate(ctx *gin.Context) {
	req := request.PemdaRequest{}

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.HTML(http.StatusOK, "registrasi_pemda.html", gin.H{
			"error": true,
		})
		return
	}

	if err := pemda.Service.Create(req); err != nil {
		ctx.String(http.StatusBadRequest, "Error :\n%+v", req)
	}

	ctx.HTML(http.StatusOK, "registrasi_pemda.html", gin.H{
		"success": true,
	})
}

func (pemda PemdaHandler) PemdaList(ctx *gin.Context) {
	var flt request.FilterQueryRequest

	if err := ctx.ShouldBindQuery(&flt); err != nil {
		ctx.HTML(http.StatusBadRequest, "error-404.html", gin.H{})
		ctx.Abort()
	}

	res, err := pemda.Service.Find(flt)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error-500.html", gin.H{})
		ctx.Abort()
	}

	ctx.HTML(http.StatusOK, "terdaftar_pemda.html", gin.H{"data": res})
}

func (pemda PemdaHandler) PemdaDetail(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := pemda.Service.FindByID(id)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error-404.html", gin.H{})
		return
	}

	ctx.HTML(http.StatusOK, "detail_pemda.html", gin.H{
		"data": res,
	})
}

func (pemda PemdaHandler) DeleteByID(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := pemda.Service.DeleteByID(id); err != nil {
		ctx.HTML(http.StatusBadRequest, "error-404.html", gin.H{})
		return
	}

	ctx.Status(http.StatusOK)
}

func (pemda PemdaHandler) UpdatePermission(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := pemda.Service.UpdatePermission(id); err != nil {
		ctx.HTML(http.StatusBadRequest, "error-404.html", gin.H{})
		return
	}

	ctx.Status(http.StatusOK)
}
