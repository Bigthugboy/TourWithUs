package tourOperatorMapper

import (
	model2 "github.com/Bigthugboy/TourWithUs/internals/domain/model/tourModel"
	model "github.com/Bigthugboy/TourWithUs/internals/domain/model/tourOpModel"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/operatorDto"
)

func MapperOperatorDtoTOObject(operator model.TourOperator) operatorDto.OperatorDto {
	return operatorDto.OperatorDto{
		FirstName:   operator.FirstName,
		LastName:    operator.LastName,
		Email:       operator.Email,
		PhoneNumber: operator.PhoneNumber,
		Password:    operator.Password,
	}
}

func MapperTourDtoToUpdateTourDto(tour model2.TourDto) model2.UpdateTourDto {
	return model2.UpdateTourDto{
		OperatorID:      &tour.OperatorID,
		TourTitle:       &tour.TourTitle,
		Location:        &tour.Location,
		Duration:        &tour.Duration,
		LanguageOffered: &tour.LanguageOffered,
		NumberOfTourist: &tour.NumberOfTourist,
		Description:     &tour.Description,
		TourGuide:       &tour.TourGuide,
		OperatorContact: &tour.OperatorContact,
		Activity:        &tour.Activity,
		Price:           &tour.Price,
		TouristEmail:    &tour.TouristEmail,
		Availability:    &tour.Availability,
		StartDate:       &tour.StartDate,
		EndDate:         &tour.EndDate,
	}
}
