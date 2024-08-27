package mapper

import (
	"github.com/Bigthugboy/TourWithUs/internals/domain/model"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/output/persistence/Db/query"
)

func MapModelToObject(tourist *model.TouristDetails) query.TouristObject {
	return query.TouristObject{
		FirstName:  tourist.FirstName,
		LastName:   tourist.LastName,
		Email:      tourist.Email,
		Password:   tourist.Password,
		ProfilePic: tourist.ProfilePic,
		Username:   tourist.Username,
	}
}

func MapObjectToModel(obj *query.TouristObject) *model.TouristDetails {
	return &model.TouristDetails{
		FirstName:  obj.FirstName,
		LastName:   obj.LastName,
		Email:      obj.Email,
		Password:   obj.Password,
		ProfilePic: obj.ProfilePic,
		Username:   obj.Username,
	}
}
