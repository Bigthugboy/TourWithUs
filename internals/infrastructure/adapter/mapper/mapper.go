package mapper

import (
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/touristDto"
)

func MapModelToObject(tourist *touristDto.TouristDetails) touristDto.TouristObject {
	return touristDto.TouristObject{
		FirstName:  tourist.FirstName,
		LastName:   tourist.LastName,
		Email:      tourist.Email,
		Password:   tourist.Password,
		ProfilePic: tourist.ProfilePic,
		Username:   tourist.Username,
	}
}

func MapObjectToModel(obj *touristDto.TouristObject) *touristDto.TouristDetails {
	if obj == nil {
		return &touristDto.TouristDetails{}
	}
	return &touristDto.TouristDetails{
		FirstName:  obj.FirstName,
		LastName:   obj.LastName,
		Email:      obj.Email,
		Password:   obj.Password,
		ProfilePic: obj.ProfilePic,
		Username:   obj.Username,
	}
}
