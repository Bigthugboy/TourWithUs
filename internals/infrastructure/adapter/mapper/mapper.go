package mapper

import (
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto"
)

func MapModelToObject(tourist *dto.TouristDetails) dto.TouristObject {
	return dto.TouristObject{
		FirstName:  tourist.FirstName,
		LastName:   tourist.LastName,
		Email:      tourist.Email,
		Password:   tourist.Password,
		ProfilePic: tourist.ProfilePic,
		Username:   tourist.Username,
	}
}

func MapObjectToModel(obj *dto.TouristObject) *dto.TouristDetails {
	if obj == nil {
		return &dto.TouristDetails{}
	}
	return &dto.TouristDetails{
		FirstName:  obj.FirstName,
		LastName:   obj.LastName,
		Email:      obj.Email,
		Password:   obj.Password,
		ProfilePic: obj.ProfilePic,
		Username:   obj.Username,
	}
}
