package touristUseCaseIntputPort

import (
	"github.com/Bigthugboy/TourWithUs/internals/domain/model"
)

type TouristUseCase interface {
	RegisterTouristUseCase(tourist *model.TouristDetails) (*model.TouristDetails, error)
}
