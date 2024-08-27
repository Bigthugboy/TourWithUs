package touristUseCaseIntputPort

import "github.com/Bigthugboy/TourWithUs/internals/domain/model"

type AuthTouristUseCase interface {
	Login(details model.TouristDetails)
}
