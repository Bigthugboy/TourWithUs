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
func MapToursToDto(tours []tourDto.TourObject) []model.TourDto {
	var dtoTours []model.TourDto
	for _, tour := range tours {
		dtoTours = append(dtoTours, MapObjectDtoToModelDto(&tour))
	}
	return dtoTours
}

func MapUpdateTourDtoToTourObject(tour *model.UpdateTourDto) map[string]interface{} {
	updates := make(map[string]interface{})
	if tour.OperatorID != nil {
		updates["OperatorID"] = *tour.OperatorID
	}
	if tour.TourTitle != nil {
		updates["TourTitle"] = *tour.TourTitle
	}
	if tour.Location != nil {
		updates["Location"] = *tour.Location
	}
	if tour.StartTime != nil {
		updates["StartTime"] = *tour.StartTime
	}
	if tour.LanguageOffered != nil {
		updates["LanguageOffered"] = *tour.LanguageOffered
	}
	if tour.NumberOfTourist != nil {
		updates["NumberOfTourist"] = *tour.NumberOfTourist
	}
	if tour.Description != nil {
		updates["Description"] = *tour.Description
	}
	if tour.TourGuide != nil {
		updates["TourGuide"] = *tour.TourGuide
	}
	if tour.TourOperator != nil {
		updates["TourOperator"] = *tour.TourOperator
	}
	if tour.OperatorContact != nil {
		updates["OperatorContact"] = *tour.OperatorContact
	}
	if tour.Category != nil {
		updates["Category"] = *tour.Category
	}
	if tour.Activity != nil {
		updates["Activity"] = *tour.Activity
	}
	if tour.Date != nil {
		updates["Date"] = *tour.Date
	}
	if tour.Price != nil {
		updates["Price"] = *tour.Price
	}
	if tour.TouristEmail != nil {
		updates["TouristEmail"] = *tour.TouristEmail
	}
	if tour.Availability != nil {
		updates["Availability"] = *tour.Availability
	}

	return updates
}
