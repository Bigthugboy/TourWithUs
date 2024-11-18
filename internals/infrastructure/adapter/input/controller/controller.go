package controller

import (
	usecase "github.com/Bigthugboy/TourWithUs/internals/application.port/tourWithUs.port/input/touristUseCase"
	"github.com/Bigthugboy/TourWithUs/internals/domain/model/touristModel"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TouristController struct {
	UseCase usecase.TouristUseCase
}

func NewTouristController(useCase usecase.TouristUseCase) *TouristController {
	return &TouristController{UseCase: useCase}
}

func (c *TouristController) RegisterTourist(ctx *gin.Context) {
	var tourist touristModel.RegisterRequest
	if err := ctx.ShouldBindJSON(&tourist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	savedTourist, err := c.UseCase.RegisterTouristUseCase(&tourist)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, savedTourist)
}

func (c *TouristController) LoginTourist(ctx *gin.Context) {
	var tourist touristModel.LoginRequest
	if err := ctx.ShouldBindJSON(&tourist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	resp, err := c.UseCase.Login(tourist)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
