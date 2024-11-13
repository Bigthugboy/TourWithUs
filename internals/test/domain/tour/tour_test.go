package tour

import (
	"github.com/Bigthugboy/TourWithUs/internals/application.port/tourWithUs.port/output/repo/tourRepo/internals/test/domain"
	"github.com/Bigthugboy/TourWithUs/internals/domain/model"
	"github.com/Bigthugboy/TourWithUs/internals/domain/services/tour"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/tourDto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func tourDetails() *model.TourDto {
	return &model.TourDto{
		OperatorID:      "1",
		TourTitle:       "Lagos City Tour",
		Location:        "Lagos",
		StartTime:       "09:00 AM",
		EndTime:         "10:00 PM",
		LanguageOffered: "English",
		NumberOfTourist: "10",
		Description:     "Explore the vibrant city of Lagos...",
		TourGuide:       "John Doe",
		TourOperator:    "TourWiz",
		OperatorContact: "08012345678",
		Category:        "City Tour",
		Activity:        "Sightseeing, Shopping",
		Date:            "2024-11-08",
		Price:           "10000",
		TouristEmail:    "johndoe@example.com",
		Availability:    true,
		TourType:        "Day Tour",
	}
}

func tourObjectDetails() tourDto.TourObject {
	return tourDto.TourObject{
		OperatorID:      "1",
		TourTitle:       "Lagos City Tour",
		Location:        "Lagos",
		StartTime:       "09:00 AM",
		EndTime:         "10:00 PM",
		LanguageOffered: "English",
		NumberOfTourist: "10",
		Description:     "Explore the vibrant city of Lagos...",
		TourGuide:       "John Doe",
		TourOperator:    "TourWiz",
		OperatorContact: "08012345678",
		Category:        "City Tour",
		Activity:        "Sightseeing, Shopping",
		Date:            "2024-11-08",
		Price:           "10000",
		TouristEmail:    "johndoe@example.com",
		Availability:    true,
		TourType:        "Day Tour",
	}
}

func TestTourCanBeCreated(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tourDetail := tourObjectDetails()
	mockDb := domain.NewMockTourRepository(ctrl)

	expectedResponse := tourDto.CreateTourResponse{
		TourTitle:       tourDetail.TourTitle,
		Message:         "Tour successfully created",
		Date:            tourDetail.Date,
		Price:           tourDetail.Price,
		OperatorContact: tourDetail.OperatorContact,
		Status:          true,
	}

	mockDb.EXPECT().CreateTour(gomock.Eq(tourDetail)).Return(expectedResponse, nil)

	service := tour.NewTour(mockDb)
	res, err := service.CreateTour(tourDetails())
	if err != nil {
		t.Error(err)
	}

	assert.NoError(t, err)
	assert.Equal(t, "Tour Created", res.Message)
	assert.Equal(t, "Lagos City Tour", res.TourTitle)
	assert.Equal(t, "08012345678", res.OperatorContact)
}
