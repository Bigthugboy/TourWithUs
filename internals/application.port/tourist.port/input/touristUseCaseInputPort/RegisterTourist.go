package touristUseCaseInputPort

import (
	"github.com/Bigthugboy/TourWithUs/internals/domain/model"
)

type TouristUseCase interface {
	RegisterTouristUseCase(tourist *model.RegisterRequest) (*model.RegisterResponse, error)
	Login(details model.LoginRequest) (*model.LoginResponse, error)
}
