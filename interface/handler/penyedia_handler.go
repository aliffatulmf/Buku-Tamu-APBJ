package handler

import (
	"net/http"

	"github.com/aliffatulmf/buku-tamu-apbj/interface/service"
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
		ctx.HTML(http.StatusBadRequest, "error-400.html", gin.H{})
		return
	}

	res, err := penyedia.Service.Find(flt)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error-500.html", gin.H{})
		return
	}

	ctx.HTML(http.StatusOK, "terdaftar_provider.html", gin.H{"data": res})
}

func (penyedia PenyediaHandler) PenyediaDetail(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := penyedia.Service.FindByID(id)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "error-404.html", gin.H{})
		return
	}

	ctx.HTML(http.StatusOK, "detail_provider.html", gin.H{"data": res})
}

func (penyedia PenyediaHandler) PenyediaDelete(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := penyedia.Service.DeleteByID(id); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (penyedia PenyediaHandler) PenyediaUpdatePermission(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := penyedia.Service.UpdatePermission(id); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.Status(http.StatusNoContent)
}

// // create record
// func (h ProviderH) ProviderCreate(ctx *gin.Context) {
// 	req := request.ProviderRequest{}

// 	if err := ctx.ShouldBind(&req); err != nil {
// 		ctx.HTML(http.StatusBadRequest, "registrasi_provider.html", gin.H{
// 			"error": true,
// 		})
// 		return
// 	}

// get all saved records
// func get_destination(d DestinationService, obj gin.H) error {
// 	dest, err := d.Find()
// 	if err != nil {
// 		return err
// 	}

// 	cons, err := d.FindConsultation()
// 	if err != nil {
// 		return err
// 	}

// 	pok, err := d.FindPokja()
// 	if err != nil {
// 		return err
// 	}

// 	obj["data"] = gin.H{
// 		"destination":  dest,
// 		"consultation": cons,
// 		"pokja":        pok,
// 	}
// 	return nil
// }

// func (penyedia PenyediaHandler) PenyediaIndex(ctx *gin.Context) {
// 	obj := gin.H{}

// 	if err := get_destination(penyedia.destination, obj); err != nil {
// 		ctx.HTML(http.StatusInternalServerError, "error-500.html", gin.H{})
// 		return
// 	}

// 	ctx.HTML(http.StatusOK, "registrasi_provider.html", obj)
// }

// 	if err := h.service.Create(req); err != nil {
// 		ctx.HTML(http.StatusInternalServerError, "error-500.html", gin.H{})
// 		return
// 	}

// 	ctx.HTML(http.StatusOK, "registrasi_provider.html", gin.H{
// 		"success": true,
// 	})
// }

// func (h ProviderH) ProviderList(ctx *gin.Context) {
// 	var req request.FilterQueryRequest

// 	if err := ctx.ShouldBindQuery(&req); err != nil {
// 		ctx.HTML(http.StatusNotFound, "error-404.html", gin.H{})
// 		return
// 	}

// 	res, err := h.service.Find(req)
// 	if err != nil {
// 		ctx.HTML(http.StatusBadRequest, "error-404.html", gin.H{})
// 		return
// 	}

// 	ctx.HTML(http.StatusOK, "terdaftar_provider.html", gin.H{"data": res})
// }

// // get record by id
// func (h ProviderH) ProviderGetDetail(ctx *gin.Context) {
// 	id := ctx.Param("id")

// 	res, err := h.service.GetByID(id)
// 	if err != nil {
// 		ctx.HTML(http.StatusBadRequest, "error-404.html", gin.H{})
// 		return
// 	}

// 	ctx.HTML(http.StatusOK, "detail_provider.html", gin.H{"data": res})
// }

// func (h ProviderH) ProviderDelete(ctx *gin.Context) {
// 	id := ctx.Param("id")

// 	if err := h.service.DeleteByID(id); err != nil {
// 		ctx.HTML(http.StatusBadRequest, "error-404.html", gin.H{})
// 		return
// 	}

// 	ctx.Status(http.StatusOK)
// }

// func (h ProviderH) ProviderUpdateStatus(ctx *gin.Context) {
// 	id := ctx.Param("id")

// 	if err := h.service.UpdatePermission(id); err != nil {
// 		ctx.HTML(http.StatusBadRequest, "error-404.html", gin.H{})
// 		return
// 	}

// 	ctx.Status(http.StatusOK)
// }
