package operatorController

import (
	"github.com/Bigthugboy/TourWithUs/internals/application.port/tourWithUs.port/input/tourOperatorUseCase"
	"github.com/Bigthugboy/TourWithUs/internals/domain/model/tourModel"
	model "github.com/Bigthugboy/TourWithUs/internals/domain/model/tourOpModel"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type OperatorController struct {
	UseCase tourOperatorUseCase.TourOperatorUseCase
}

func NewController(useCase tourOperatorUseCase.TourOperatorUseCase) *OperatorController {
	return &OperatorController{
		UseCase: useCase,
	}
}
func BindRequest(req interface{}, ctx *gin.Context) error {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	return nil
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

func (uc *OperatorController) UpdateTour(ctx *gin.Context) {
	var tour tourModel.UpdateTourDto
	idStr := ctx.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tour ID"})
		return
	}

	if err := ctx.ShouldBindJSON(&tour); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := uc.UseCase.UpdateTour(uint(id), tour)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": res})
}
func (uc *OperatorController) DeleteTour(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := uc.UseCase.DeleteTour(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": res})
}

func (uc *OperatorController) GetTourList(ctx *gin.Context) {
	res, err := uc.UseCase.ListTours()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": res})
}

func (uc *OperatorController) ViewTourDetails(ctx *gin.Context) {
	idstr := ctx.Param("id")
	log.Println("--->", idstr)
	id, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tour ID"})
		return
	}
	res, err := uc.UseCase.ViewTourDetails(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": res})
}

func (uc *OperatorController) GetTourByAOperator(ctx *gin.Context) {
	id := ctx.Param("id")
	tourId := ctx.Param("tourId")
	res, err := uc.UseCase.GetTourByOperatorId(id, tourId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": res})

}

func (uc *OperatorController) GetToursByAOperator(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := uc.UseCase.GetAllTourByOperatorId(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": res})
}

func (uc *OperatorController) ChangeTourAvailability(ctx *gin.Context) {
	tourId := ctx.Param("tourId")
	id, err := strconv.ParseUint(tourId, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tour ID"})
		return
	}
	res, err := uc.UseCase.ManageAvailability(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": res})
}
