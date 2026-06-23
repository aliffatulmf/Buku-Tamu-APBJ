package handler

import (
	"log"
	"net/http"

	"github.com/aliffatulmf/buku-tamu-apbj/internal/service"
	"github.com/aliffatulmf/buku-tamu-apbj/request"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	Service service.DashboardService
}

func NewDashboardHandler(service service.DashboardService) DashboardHandler {
	return DashboardHandler{service}
}

func (dash DashboardHandler) DashbordIndex(ctx *gin.Context) {
	cnt := dash.Service.DashboardCounts()
	ctx.HTML(http.StatusOK, "dashboard.html", gin.H{
		"pemda":    cnt.Pemda,
		"provider": cnt.Provider,
		"pokja":    cnt.Pokja,
		"instansi": cnt.Instansi,
	})
}

func (dash DashboardHandler) DashboardExport(ctx *gin.Context) {
	var req request.ExportQuery
	var param service.DType

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.HTML(http.StatusBadRequest, "error-404.html", gin.H{})
		return
	}

	switch req.Table {
	case "pemda":
		param = service.PemdaType
	case "penyedia":
		param = service.PenyediaType
	}

	if err := dash.Service.GoExport(param, req); err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadGateway)
	}

	ctx.Status(http.StatusOK)
}
