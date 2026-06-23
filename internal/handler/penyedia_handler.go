package handler

import (
	"net/http"
	"strconv"

	"github.com/aliffatulmf/buku-tamu-apbj/internal/service"
	"github.com/aliffatulmf/buku-tamu-apbj/request"

	"github.com/gin-gonic/gin"
)

type PenyediaHandler struct {
	Service service.PenyediaService
	Tujuan  service.TujuanService
}

func NewPenyediaHandler(service service.PenyediaService, destination service.TujuanService) PenyediaHandler {
	return PenyediaHandler{
		Service: service,
		Tujuan:  destination,
	}
}

func (penyedia PenyediaHandler) PenyediaIndex(ctx *gin.Context) {
	d, _ := penyedia.Tujuan.FindTujuan()
	c, _ := penyedia.Tujuan.FindConsultation()
	p, _ := penyedia.Tujuan.FindPokja()

	ctx.HTML(http.StatusOK, "registrasi_provider.html", gin.H{
		"data": gin.H{
			"destination":  d,
			"consultation": c,
			"pokja":        p,
		},
	})
}

func (penyedia PenyediaHandler) PenyediaCreate(ctx *gin.Context) {
	var req request.PenyediaRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.HTML(http.StatusBadRequest, "registrasi_provider.html", gin.H{"error": true})
		return
	}

	if err := penyedia.Service.Create(req); err != nil {
		ctx.HTML(http.StatusInternalServerError, "registrasi_provider.html", gin.H{"error": true})
		return
	}

	ctx.HTML(http.StatusOK, "registrasi_provider.html", gin.H{"success": true})
}

func (penyedia PenyediaHandler) PenyediaList(ctx *gin.Context) {
	var flt request.FilterQueryRequest

	if err := ctx.ShouldBindQuery(&flt); err != nil {
		ctx.Redirect(http.StatusFound, "/penyedia/terdaftar")
		return
	}

	res, err := penyedia.Service.Find(flt)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/penyedia/terdaftar")
		return
	}

	ctx.HTML(http.StatusOK, "terdaftar_provider.html", gin.H{"data": res})
}

func (penyedia PenyediaHandler) PenyediaDetail(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error-404.html", gin.H{})
		return
	}

	res, err := penyedia.Service.FindByID(uint(id))
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error-404.html", gin.H{})
		return
	}

	ctx.HTML(http.StatusOK, "detail_provider.html", gin.H{"data": res})
}

func (penyedia PenyediaHandler) PenyediaDelete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	if err := penyedia.Service.DeleteByID(uint(id)); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (penyedia PenyediaHandler) PenyediaUpdatePermission(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	if err := penyedia.Service.UpdatePermission(uint(id)); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.Status(http.StatusNoContent)
}
