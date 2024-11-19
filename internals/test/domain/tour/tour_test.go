package tour

import (
	"errors"
	"github.com/Bigthugboy/TourWithUs/internals/application.port/tourWithUs.port/output/repo/tourRepo/internals/test/domain"
	"github.com/Bigthugboy/TourWithUs/internals/domain/domainMapper/tourMapper"
	"github.com/Bigthugboy/TourWithUs/internals/domain/exception"
	model "github.com/Bigthugboy/TourWithUs/internals/domain/model/tourModel"
	"github.com/Bigthugboy/TourWithUs/internals/domain/services/tour"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/tourDto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
	"time"
)

func tourDetails() *model.TourDto {

	return &model.TourDto{
		OperatorID:      "1",
		TourTitle:       "Lagos City Tour",
		Location:        "Lagos",
		Duration:        "3 days",
		LanguageOffered: "English",
		NumberOfTourist: "10",
		Description:     "Explore the vibrant city of Lagos...",
		TourGuide:       "John Doe",
		OperatorContact: "08012345678",
		Activity:        "Sightseeing, Shopping",
		Price:           10000.0,
		TouristEmail:    "johndoe@example.com",
		Availability:    true,
		TourType:        "Day Tour",
		StartDate:       "2024-11-15",
		EndDate:         "2024-11-20",
	}
}

func tourObjectDetails() tourDto.TourObject {
	return tourDto.TourObject{
		OperatorID:      "1",
		TourTitle:       "Lagos City Tour",
		Location:        "Lagos",
		Duration:        "3 days",
		LanguageOffered: "English",
		NumberOfTourist: "10",
		Description:     "Explore the vibrant city of Lagos...",
		TourGuide:       "John Doe",
		OperatorContact: "08012345678",
		Activity:        "Sightseeing, Shopping",
		Price:           10000.0,
		TouristEmail:    "johndoe@example.com",
		Availability:    true,
		TourType:        "Day Tour",
		StartDate:       "2024-11-15",
		EndDate:         "2024-11-20",
	}
}

func tourObjectList() []tourDto.TourObject {
	return []tourDto.TourObject{
		tourObjectDetails(),
	}
}

func tourObjectDetailsList() []tourDto.TourObject {

	return []tourDto.TourObject{
		{
			OperatorID:      "1",
			TourTitle:       "Lagos City Tour",
			Location:        "Lagos",
			Duration:        "3 days",
			LanguageOffered: "English",
			NumberOfTourist: "10",
			Description:     "Explore the vibrant city of Lagos...",
			TourGuide:       "John Doe",
			OperatorContact: "08012345678",
			Activity:        "Sightseeing, Shopping",
			Price:           10000.0,
			TouristEmail:    "johndoe@example.com",
			Availability:    false,
			TourType:        "Day Tour",
			StartDate:       "2024-11-15",
			EndDate:         "2024-11-20",
		},
		{
			OperatorID:      "2",
			TourTitle:       "Abuja Adventure",
			Location:        "Abuja",
			Duration:        "2 days",
			LanguageOffered: "French",
			NumberOfTourist: "8",
			Description:     "Discover the heart of Nigeria’s capital...",
			TourGuide:       "Jane Smith",
			OperatorContact: "08098765432",
			Activity:        "Hiking, Cultural Tour",
			Price:           15000.0,
			TouristEmail:    "janesmith@example.com",
			Availability:    true,
			TourType:        "Adventure Tour",
			StartDate:       "2024-11-15",
			EndDate:         "2024-11-20",
		},
		{
			OperatorID:      "3",
			TourTitle:       "Kano Historical Sites",
			Location:        "Kano",
			Duration:        "1 day",
			LanguageOffered: "Hausa",
			NumberOfTourist: "15",
			Description:     "Visit the ancient city walls and museum...",
			TourGuide:       "Ahmed Musa",
			OperatorContact: "08123456789",
			Activity:        "Sightseeing, Historical",
			Price:           8000.0,
			TouristEmail:    "ahmedmusa@example.com",
			Availability:    true,
			TourType:        "Historical Tour",
			StartDate:       "2024-11-15",
			EndDate:         "2024-11-20",
		},
		{
			OperatorID:      "4",
			TourTitle:       "Calabar Festival Experience",
			Location:        "Calabar",
			Duration:        "5 days",
			LanguageOffered: "English",
			NumberOfTourist: "20",
			Description:     "Experience the annual Calabar Festival...",
			TourGuide:       "Emeka Okafor",
			OperatorContact: "08055667788",
			Activity:        "Festival, Cultural Tour",
			Price:           20000.0,
			TouristEmail:    "emekaokafor@example.com",
			Availability:    true,
			TourType:        "Festival Tour",
			StartDate:       "2024-11-15",
			EndDate:         "2024-11-20",
		},
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

	t.Run("error in creating tourModel", func(t *testing.T) {
		mockDb.EXPECT().CreateTour(gomock.Eq(tourDetail)).Return(tourDto.CreateTourResponse{}, errors.New("database error"))
		result, err := service.CreateTour(tourDetails())
		assert.Nil(t, result)
		assert.Error(t, err)

		var tourErr *exception.TourWithUsError
		ok := errors.As(err, &tourErr)
		assert.True(t, ok, "Expected a TourWithUsError type")
		assert.Equal(t, exception.ErrFailToCreateTour, tourErr.Message)
		assert.Equal(t, http.StatusInternalServerError, tourErr.StatusCode)
		assert.EqualError(t, tourErr.ErrorMessage, "database error")
	})
}

func TestTourCreationRequestValidation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := domain.NewMockTourRepository(ctrl)
	tourDetail := tourDetails()
	service := tour.NewTour(mockDb)

	t.Run("empty/blank tourModel title", func(t *testing.T) {
		tourDetail.TourTitle = ""
		service := tour.NewTour(mockDb)
		res, err := service.CreateTour(tourDetail)
		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.Error(t, err)
	})

	t.Run("empty/blank operatorId", func(t *testing.T) {
		tourDetail = tourDetails()
		tourDetail.OperatorID = ""
		res, err := service.CreateTour(tourDetail)
		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.Error(t, err)
	})

	t.Run("empty/blank location", func(t *testing.T) {
		tourDetail = tourDetails()
		tourDetail.Location = ""
		service = tour.NewTour(mockDb)
		res, err := service.CreateTour(tourDetail)
		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.Error(t, err)
	})

}

func TestGetTourById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := domain.NewMockTourRepository(ctrl)
	id := "1"
	mockDb.EXPECT().GetTourById(gomock.Eq(id)).Return(tourObjectDetails(), nil)
	service := tour.NewTour(mockDb)
	res, err := service.GetTourById(id)
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, res)
	assert.Equal(t, "Lagos City Tour", res.TourTitle)
	assert.Equal(t, "08012345678", res.OperatorContact)
}

func TestGetAllTours(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := domain.NewMockTourRepository(ctrl)
	mockDb.EXPECT().GetAllTours().Return(tourObjectList(), nil)
	service := tour.NewTour(mockDb)
	res, err := service.GetAllTours()
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, res)
}

func TestGetAvailableTours(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := domain.NewMockTourRepository(ctrl)
	mockDb.EXPECT().GetAvailableTours().Return(tourObjectDetailsList(), nil)
	service := tour.NewTour(mockDb)
	res, err := service.GetAvailableTours()
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, res)
}

func TestGetAvailableTours_NoAvailableTours(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := domain.NewMockTourRepository(ctrl)
	mockDb.EXPECT().GetAvailableTours().Return([]tourDto.TourObject{}, nil)
	service := tour.NewTour(mockDb)
	res, err := service.GetAvailableTours()
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, res)
}

func TestGetTourByLocation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := domain.NewMockTourRepository(ctrl)
	location := "Lagos"
	mockDb.EXPECT().GetToursByLocation(gomock.Eq(location)).Return(tourObjectList(), nil)
	service := tour.NewTour(mockDb)
	res, err := service.GetToursByLocation(location)
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, res)
}

func TestGetToursByDateRange(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDb := domain.NewMockTourRepository(ctrl)
	service := tour.NewTour(mockDb)

	startDate := "2024-11-01"
	endDate := "2024-11-30"

	startDateStr, _ := time.Parse("2006-01-02", startDate)
	endDateStr, _ := time.Parse("2006-01-02", endDate)

	t.Run("should return tours within date range", func(t *testing.T) {
		mockDb.EXPECT().GetToursByDateRange(startDateStr, endDateStr).Return(tourObjectDetailsList(), nil)

		result, err := service.GetToursByDateRange(startDate, endDate)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 4, len(result), "Expected 4 tours in result")
	})

	t.Run("should return empty list when no tours found", func(t *testing.T) {
		mockDb.EXPECT().GetToursByDateRange(startDateStr, endDateStr).Return([]tourDto.TourObject{}, nil)

		result, err := service.GetToursByDateRange(startDate, endDate)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result, "Expected no tours in result")
	})

	t.Run("should return error when database fails", func(t *testing.T) {
		mockDb.EXPECT().GetToursByDateRange(startDateStr, endDateStr).Return(nil, errors.New("database error"))

		result, err := service.GetToursByDateRange(startDate, endDate)
		assert.Nil(t, result)
		assert.Error(t, err)

		var tourErr *exception.TourWithUsError
		ok := errors.As(err, &tourErr)
		assert.True(t, ok, "Expected a TourWithUsError type")
		assert.Equal(t, exception.ErrFailToGetTour, tourErr.Message)
		assert.Equal(t, http.StatusInternalServerError, tourErr.StatusCode)
		assert.EqualError(t, tourErr.ErrorMessage, "database error")
	})
}

func TestGetTourByPriceRange(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDb := domain.NewMockTourRepository(ctrl)
	service := tour.NewTour(mockDb)
	MinPrice := 800.0
	MaxPrice := 20000.0

	t.Run("should return tours within date range", func(t *testing.T) {
		mockDb.EXPECT().GetToursByPriceRange(MinPrice, MaxPrice).Return([]tourDto.TourObject{
			{TourTitle: "Lagos City Tour", Price: 10000.0},
			{TourTitle: "Abuja City Adventure", Price: 12000.0},
		}, nil)

		result, err := service.GetToursByPriceRange(MinPrice, MaxPrice)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 2, len(result), "Expected 2 tours in result")
	})

	t.Run("should return empty list when no tours found", func(t *testing.T) {
		mockDb.EXPECT().GetToursByPriceRange(MinPrice, MaxPrice).Return([]tourDto.TourObject{}, nil)

		result, err := service.GetToursByPriceRange(MinPrice, MaxPrice)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result, "Expected no tours in result")
	})
}

func TestSearchTours(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := domain.NewMockTourRepository(ctrl)
	query := "Lagos City Tour"
	mockDb.EXPECT().SearchTours(query).Return(tourObjectList(), nil)
	service := tour.NewTour(mockDb)
	res, err := service.SearchTours(query)
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, res)
}

func TestDeleteTour(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := domain.NewMockTourRepository(ctrl)
	id := "1"
	mockDb.EXPECT().DeleteTour(id).Return(nil)
	service := tour.NewTour(mockDb)
	deleteTour, err := service.DeleteTour(id)
	if err != nil {
		return
	}
	log.Println(deleteTour)
	assert.NotNil(t, deleteTour)
}

func TestUpdateTour_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDb := domain.NewMockTourRepository(ctrl)
	price := 15000.0
	tourTitle := "Updated Tour Title"
	location := "Updated Location"

	dto := model.UpdateTourDto{
		TourTitle: &tourTitle,
		Price:     &price,
		Location:  &location,
	}

	objectDto := tourDto.UpdateTourDto{
		TourTitle: &tourTitle,
		Price:     &price,
		Location:  &location,
	}

	expectedTour := tourDto.TourObject{
		TourTitle:       "Updated Tour Title",
		Location:        "Updated Location",
		Price:           15000.0,
		Duration:        "2 days",
		LanguageOffered: "French",
		NumberOfTourist: "8",
		Description:     "Discover the heart of Nigeria’s capital...",
		TourGuide:       "Jane Smith",
		OperatorContact: "08098765432",
		Activity:        "Hiking, Cultural Tour",
		TouristEmail:    "janesmith@example.com",
		Availability:    true,
		TourType:        "Adventure Tour",
		StartDate:       "2024-11-15",
		EndDate:         "2024-11-20",
	}
	mockDb.EXPECT().UpdateTour("1", gomock.Eq(objectDto)).Return(expectedTour, nil)

	service := tour.NewTour(mockDb)
	res, err := service.UpdateTour("1", dto)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "Updated Tour Title", res.TourTitle)
	assert.Equal(t, 15000.0, res.Price)
}

func TestUpdateTour_InvalidPriceFormat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDb := domain.NewMockTourRepository(ctrl)

	tourTitle := ""
	location := "Updated Location"
	price := 0.0

	dto := model.UpdateTourDto{
		TourTitle: &tourTitle,
		Price:     &price,
		Location:  &location,
	}
	req := tourMapper.MapUpdateTourDtoToTourObject(&dto)

	mockDb.EXPECT().UpdateTour("1", req).Return(tourDto.TourObject{}, errors.New("invalid price format"))

	service := tour.NewTour(mockDb)
	res, err := service.UpdateTour("1", dto)

	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestUpdateTour_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDb := domain.NewMockTourRepository(ctrl)

	tourTitle := ""
	location := "Updated Location"
	price := 0.0

	dto := model.UpdateTourDto{
		TourTitle: &tourTitle,
		Price:     &price,
		Location:  &location,
	}
	req := tourMapper.MapUpdateTourDtoToTourObject(&dto)

	mockDb.EXPECT().UpdateTour("1", req).Return(tourDto.TourObject{}, errors.New("database error"))

	service := tour.NewTour(mockDb)
	res, err := service.UpdateTour("1", dto)

	assert.Nil(t, res)
	assert.Error(t, err)
	var tourErr *exception.TourWithUsError
	ok := errors.As(err, &tourErr)
	assert.True(t, ok, "Expected error to be of type TourWithUsError")
	assert.Equal(t, exception.ErrFailToUpdateTour, tourErr.Message)
	assert.Equal(t, http.StatusInternalServerError, tourErr.StatusCode)

}
