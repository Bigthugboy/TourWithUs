package operatorController

import (
	"github.com/Bigthugboy/TourWithUs/internals/application.port/tourWithUs.port/input/tourOperatorUseCase"
	"github.com/Bigthugboy/TourWithUs/internals/domain/model/tourModel"
	model "github.com/Bigthugboy/TourWithUs/internals/domain/model/tourOpModel"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OperatorController struct {
	UseCase tourOperatorUseCase.TourOperatorUseCase
}

func NewController(useCase tourOperatorUseCase.TourOperatorUseCase) *OperatorController {
	return &OperatorController{
		UseCase: useCase,
	}
}

func (uc *OperatorController) RegisterOperator(ctx *gin.Context) {
	var operator model.TourOperator
	if err := ctx.ShouldBindJSON(&operator); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := uc.UseCase.RegisterTourOperator(operator)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

func (uc *OperatorController) LoginOperator(ctx *gin.Context) {
	var req model.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := uc.UseCase.Login(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

func (uc *OperatorController) CreateTour(ctx *gin.Context) {
	var tour tourModel.TourDto
	if err := ctx.ShouldBindJSON(&tour); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := uc.UseCase.CreateTour(tour)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})

}
