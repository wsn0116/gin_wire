package controllers

import (
	"net/http"

	"example.com/gin_wire/app/usecases"
	"github.com/gin-gonic/gin"
)

type SampleController struct {
	sampleUsecase *usecases.SampleUsecase
}

func NewSampleController(
	sampleUsecase *usecases.SampleUsecase,
) *SampleController {
	return &SampleController{
		sampleUsecase: sampleUsecase,
	}
}

func (controller *SampleController) Get(c *gin.Context) {
	msg := controller.sampleUsecase.Get()
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}
