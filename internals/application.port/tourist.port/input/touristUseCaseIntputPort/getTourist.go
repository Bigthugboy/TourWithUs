package touristUseCaseIntputPort

import (
	"github.com/Bigthugboy/TourWithUs/internals/domain/model"
	"github.com/gin-gonic/gin"
)

type GetTouristUseCase interface {
	GetTouristByEMail(email string) (model.TouristDetails, error)
	gin.HandlerFunc
}
