package domainMapper

import (
	"github.com/Bigthugboy/TourWithUs/internals/domain/model"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto"
)

func MapRegisterRequestToTouristDetails(details *model.RegisterRequest) dto.TouristDetails {
	return dto.TouristDetails{
		FirstName:  details.FirstName,
		LastName:   details.LastName,
		Email:      details.Email,
		Password:   details.Password,
		ProfilePic: details.ProfilePic,
		Username:   details.Username,
	}

}

func MapRegisterRequestToTouristObject(request *model.RegisterRequest) dto.TouristObject {
	return dto.TouristObject{
		FirstName:  request.FirstName,
		LastName:   request.LastName,
		Email:      request.Email,
		Password:   request.Password,
		ProfilePic: request.ProfilePic,
		Username:   request.Username,
	}
}

func MapLoginRequestToTouristDetails(details *model.LoginRequest) dto.RetrieveTourist {
	return dto.RetrieveTourist{
		Email:    details.Email,
		Password: details.Password,
	}
}
