package handler

import (
	"github.com/aliffatulmf/buku-tamu-apbj/entity"

	"github.com/gin-gonic/gin"
)

type DestinationService interface {
	Find(conds ...interface{}) ([]entity.Destination, error)
	FindPokja() ([]entity.Pokja, error)
	FindPokjaByID(id int) (entity.Pokja, error)
	FindConsultation() ([]entity.Consultation, error)
}

type destinationHandler struct {
	service DestinationService
}

func NewDestinationHandler(service DestinationService) destinationHandler {
	return destinationHandler{service}
}

func (s destinationHandler) DestinationIndex(ctx *gin.Context) {
	res, err := s.service.Find()
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": res,
	})
}
