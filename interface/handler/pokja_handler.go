package handler

import (
	"net/http"

	"github.com/aliffatulmf/buku-tamu-apbj/interface/service"
	"github.com/aliffatulmf/buku-tamu-apbj/request"

	"github.com/gin-gonic/gin"
)

type PokjaHandler struct {
	Service service.PokjaService
}

func NewPokjaHandler(service service.PokjaService) PokjaHandler {
	return PokjaHandler{Service: service}
}

func (pokja PokjaHandler) PokjaIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "registrasi_pokja.html", gin.H{})
}

func (pokja PokjaHandler) PokjaCreate(ctx *gin.Context) {
	var req request.PokjaCreateRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.HTML(http.StatusBadRequest, "registrasi_pokja.html", gin.H{"errors": true})
		return
	}

	if err := pokja.Service.Create(req); err != nil {
		ctx.HTML(http.StatusInternalServerError, "registrasi_pokja.html", gin.H{"errors": true})
		return
	}

	ctx.HTML(http.StatusOK, "registrasi_pokja.html", gin.H{"success": true})
}

func (pokja PokjaHandler) PokjaList(ctx *gin.Context) {
	res, err := pokja.Service.Find()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error-500.html", gin.H{})
		return
	}

	ctx.HTML(http.StatusOK, "terdaftar_pokja.html", gin.H{"data": res})
}

func (pokja PokjaHandler) PokjaDetail(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := pokja.Service.FindByID(id)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error-404.html", gin.H{})
		return
	}

	ctx.HTML(http.StatusOK, "detail_pokja.html", gin.H{"data": res})
}

func (pokja PokjaHandler) PokjaUpdate(ctx *gin.Context) {
	var req request.PokjaRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.HTML(http.StatusBadRequest, "error-404.html", gin.H{})
		return
	}

	if err := pokja.Service.UpdateStatus(req); err != nil {
		ctx.HTML(http.StatusBadRequest, "error-404.html", gin.H{})
		return
	}

	ctx.Redirect(http.StatusFound, "/pokja/terdaftar")
}

func (pokja PokjaHandler) PokjaDelete(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := pokja.Service.Delete(id); err != nil {
		ctx.HTML(http.StatusBadRequest, "error-404.html", gin.H{})
		return
	}

	ctx.Status(http.StatusOK)
}
