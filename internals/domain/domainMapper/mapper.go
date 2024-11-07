package domainMapper

import (
	"github.com/Bigthugboy/TourWithUs/internals/domain/model"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/tourDto"
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

func MapModelTourDtoToTourDBObject(tour *model.TourDto) tourDto.TourObject {
	return tourDto.TourObject{
		OperatorID:      tour.OperatorID,
		TourTitle:       tour.TourTitle,
		Location:        tour.Location,
		StartTime:       tour.StartTime,
		LanguageOffered: tour.LanguageOffered,
		NumberOfTourist: tour.NumberOfTourist,
		Description:     tour.Description,
		TourGuide:       tour.TourGuide,
		TourOperator:    tour.TourOperator,
		OperatorContact: tour.OperatorContact,
		Category:        tour.Category,
		Activity:        tour.Activity,
		Date:            tour.Date,
		Price:           tour.Price,
		TouristEmail:    tour.TouristEmail,
		Availability:    tour.Availability,
		TourType:        tourDto.TourType(tour.TourType),
	}
}
func MapObjectDtoToModelDto(tour *tourDto.TourObject) model.TourDto {
	return model.TourDto{
		OperatorID:      tour.OperatorID,
		TourTitle:       tour.TourTitle,
		Location:        tour.Location,
		StartTime:       tour.StartTime,
		LanguageOffered: tour.LanguageOffered,
		NumberOfTourist: tour.NumberOfTourist,
		Description:     tour.Description,
		TourGuide:       tour.TourGuide,
		TourOperator:    tour.TourOperator,
		OperatorContact: tour.OperatorContact,
		Category:        tour.Category,
		Activity:        tour.Activity,
		Date:            tour.Date,
		Price:           tour.Price,
		TouristEmail:    tour.TouristEmail,
		Availability:    tour.Availability,
		TourType:        model.TourType(tour.TourType),
	}
}
