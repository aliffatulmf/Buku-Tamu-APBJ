package handler

import (
	"net/http"
	"strconv"

	"github.com/aliffatulmf/buku-tamu-apbj/internal/service"
	"github.com/aliffatulmf/buku-tamu-apbj/request"

	"github.com/gin-gonic/gin"
)

type InstansiHandler struct {
	Service service.InstansiService
}

func NewAgencyHandler(service service.InstansiService) InstansiHandler {
	return InstansiHandler{Service: service}
}

func (instansi InstansiHandler) InstansiIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "registrasi_instansi.html", gin.H{})
}

func (instansi InstansiHandler) InstansiCreate(ctx *gin.Context) {
	var req request.InstansiRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.HTML(http.StatusBadRequest, "registrasi_instansi.html", gin.H{"error": true})
		return
	}

	if err := instansi.Service.Create(req); err != nil {
		ctx.HTML(http.StatusInternalServerError, "registrasi_instansi.html", gin.H{"error": true})
		return
	}

	ctx.HTML(http.StatusOK, "registrasi_instansi.html", gin.H{"success": true})
}

func (instansi InstansiHandler) InstansiFind(ctx *gin.Context) {
	res, err := instansi.Service.Find()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error-500.html", gin.H{})
		return
	}

	ctx.HTML(http.StatusOK, "terdaftar_instansi.html", gin.H{"data": res})
}

func (instansi InstansiHandler) InstansiDetail(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error-404.html", gin.H{})
		return
	}

	res, err := instansi.Service.FindByID(uint(id))
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error-404.html", gin.H{})
		return
	}

	ctx.HTML(http.StatusOK, "detail_instansi.html", gin.H{"data": res})
}

func (instansi InstansiHandler) InstansiUpdate(ctx *gin.Context) {
	var req request.InstansiRequest
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error-404.html", gin.H{})
		return
	}
	req.ID = uint(id)

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.HTML(http.StatusBadRequest, "error-404.html", gin.H{})
		return
	}

	res, err := instansi.Service.FindByID(req.ID)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error-404.html", gin.H{})
		return
	}

	if err := instansi.Service.Update(req); err != nil {
		ctx.HTML(http.StatusOK, "detail_instansi.html", gin.H{"data": res, "error": true})
		return
	}

	res.AgencyName = req.Name
	res.AgencyEmail = req.Email
	res.AgencyTelephone = req.Telephone

	ctx.HTML(http.StatusOK, "detail_instansi.html", gin.H{"data": res, "success": true})
}
