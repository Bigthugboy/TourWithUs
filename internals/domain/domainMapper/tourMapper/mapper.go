package tourMapper

import (
	"github.com/Bigthugboy/TourWithUs/internals/domain/model/tourModel"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/tourDto"
)

func MapModelTourDtoToTourDBObject(tour *tourModel.TourDto) tourDto.TourObject {
	return tourDto.TourObject{
		OperatorID:      tour.OperatorID,
		TourTitle:       tour.TourTitle,
		Location:        tour.Location,
		Duration:        tour.Duration,
		LanguageOffered: tour.LanguageOffered,
		NumberOfTourist: tour.NumberOfTourist,
		Description:     tour.Description,
		TourGuide:       tour.TourGuide,
		OperatorContact: tour.OperatorContact,
		Activity:        tour.Activity,
		Date:            tour.Date,
		Price:           tour.Price,
		TouristEmail:    tour.TouristEmail,
		Availability:    tour.Availability,
		StartDate:       tour.StartDate,
		EndDate:         tour.EndDate,
		TourType:        tourDto.TourType(tour.TourType),
	}
}
func MapObjectDtoToModelDto(tour *tourDto.TourObject) tourModel.TourDto {
	return tourModel.TourDto{
		OperatorID:      tour.OperatorID,
		TourTitle:       tour.TourTitle,
		Location:        tour.Location,
		Duration:        tour.Duration,
		LanguageOffered: tour.LanguageOffered,
		NumberOfTourist: tour.NumberOfTourist,
		Description:     tour.Description,
		TourGuide:       tour.TourGuide,
		OperatorContact: tour.OperatorContact,
		Activity:        tour.Activity,
		Date:            tour.Date,
		Price:           tour.Price,
		TouristEmail:    tour.TouristEmail,
		Availability:    tour.Availability,
		StartDate:       tour.StartDate,
		EndDate:         tour.EndDate,
		TourType:        tourModel.TourType(tour.TourType),
	}
}
func MapToursToDto(tours []tourDto.TourObject) []tourModel.TourDto {
	var dtoTours []tourModel.TourDto
	for _, tour := range tours {
		dtoTours = append(dtoTours, MapObjectDtoToModelDto(&tour))
	}
	return dtoTours
}

func MapUpdateTourDtoToTourObject(tour *tourModel.UpdateTourDto) tourDto.UpdateTourDto {
	return tourDto.UpdateTourDto{
		OperatorID:      tour.OperatorID,
		TourTitle:       tour.TourTitle,
		Location:        tour.Location,
		Duration:        tour.Duration,
		LanguageOffered: tour.LanguageOffered,
		NumberOfTourist: tour.NumberOfTourist,
		Description:     tour.Description,
		TourGuide:       tour.TourGuide,
		OperatorContact: tour.OperatorContact,
		Activity:        tour.Activity,
		Price:           tour.Price,
		TouristEmail:    tour.TouristEmail,
		Availability:    tour.Availability,
		StartDate:       tour.StartDate,
		EndDate:         tour.EndDate,
	}
}
