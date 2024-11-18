package touristUseCase

import (
	model "github.com/Bigthugboy/TourWithUs/internals/domain/model/touristModel"
)

type TouristUseCase interface {
	RegisterTouristUseCase(tourist *model.RegisterRequest) (*model.RegisterResponse, error)
	Login(details model.LoginRequest) (*model.LoginResponse, error)
}
