package domainMapper

import (
	"github.com/Bigthugboy/TourWithUs/internals/domain/model"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/touristDto"
)

func MapRegisterRequestToTouristDetails(details *model.RegisterRequest) touristDto.TouristDetails {
	return touristDto.TouristDetails{
		FirstName:  details.FirstName,
		LastName:   details.LastName,
		Email:      details.Email,
		Password:   details.Password,
		ProfilePic: details.ProfilePic,
		Username:   details.Username,
	}

}

func MapRegisterRequestToTouristObject(request *model.RegisterRequest) touristDto.TouristObject {
	return touristDto.TouristObject{
		FirstName:  request.FirstName,
		LastName:   request.LastName,
		Email:      request.Email,
		Password:   request.Password,
		ProfilePic: request.ProfilePic,
		Username:   request.Username,
	}
}

func MapLoginRequestToTouristDetails(details *model.LoginRequest) touristDto.RetrieveTourist {
	return touristDto.RetrieveTourist{
		Email:    details.Email,
		Password: details.Password,
	}
}
