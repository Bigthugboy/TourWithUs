package tour

import (
	"errors"
	usecase "github.com/Bigthugboy/TourWithUs/internals/application.port/tourWithUs.port/input/tourUseCase"
	database "github.com/Bigthugboy/TourWithUs/internals/application.port/tourWithUs.port/output/repo/tourRepo"
	"github.com/Bigthugboy/TourWithUs/internals/domain/domainMapper/tourMapper"
	"github.com/Bigthugboy/TourWithUs/internals/domain/exception"
	model "github.com/Bigthugboy/TourWithUs/internals/domain/model/tourModel"
	"github.com/Bigthugboy/TourWithUs/internals/domain/services"
	"net/http"
	"strings"
	"time"
)

type Tour struct {
	DB database.TourRepository
}

func NewTour(db database.TourRepository) usecase.TourUseCaseInputPort {
	return &Tour{
		DB: db,
	}
}

func (t *Tour) CreateTour(request *model.TourDto) (*model.CreateTourResponse, error) {
	if err := services.ValidateRequest(request); err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrValidatingRequest,
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: err,
		}
	}
	req := tourMapper.MapModelTourDtoToTourDBObject(request)
	res, err := t.DB.CreateTour(req)
	if err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrFailToCreateTour,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err,
		}
	}
	response := &model.CreateTourResponse{
		Message:         "Tour Created",
		TourTitle:       res.TourTitle,
		TourId:          res.TourId,
		OperatorContact: res.OperatorContact,
		Price:           res.Price,
	}
	return response, nil
}

func (t *Tour) GetTourById(id uint) (*model.TourDto, error) {
	res, err := t.DB.GetTourById(id)
	if err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrFailToGetTour,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err,
		}
	}
	resp := tourMapper.MapObjectDtoToModelDto(&res)
	return &resp, nil
}

func (t *Tour) GetAllTours() ([]model.TourDto, error) {
	res, err := t.DB.GetAllTours()
	if err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrFailToGetTour,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err,
		}
	}
	if len(res) == 0 {
		return []model.TourDto{}, nil
	}
	return tourMapper.MapToursToDto(res), nil
}

func (t *Tour) GetAvailableTours() ([]model.TourDto, error) {
	res, err := t.DB.GetAvailableTours()
	if err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrFailToGetTour,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err,
		}
	}
	if len(res) == 0 {
		return []model.TourDto{}, nil
	}
	return tourMapper.MapToursToDto(res), nil
}

func (t *Tour) GetToursByLocation(location string) ([]model.TourDto, error) {
	if location == "" || strings.TrimSpace(location) == "" {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrInvalidRequest,
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: errors.New("invalid tourModel location"),
		}
	}
	res, err := t.DB.GetToursByLocation(location)
	if err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrFailToGetTour,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err,
		}
	}
	if len(res) == 0 {
		return []model.TourDto{}, nil
	}
	return tourMapper.MapToursToDto(res), nil
}

func (t *Tour) GetToursByDateRange(startDate, endDate string) ([]model.TourDto, error) {
	if startDate == "" || endDate == "" {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrInvalidRequest,
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: errors.New("invalid tourModel date"),
		}
	}
	startDateStr, _ := time.Parse("2006-01-02", startDate)
	endDateStr, _ := time.Parse("2006-01-02", endDate)

	res, err := t.DB.GetToursByDateRange(startDateStr, endDateStr)
	if err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrFailToGetTour,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err,
		}
	}
	if len(res) == 0 {
		return []model.TourDto{}, nil
	}
	return tourMapper.MapToursToDto(res), nil
}

func (t *Tour) GetToursByPriceRange(minPrice, maxPrice float64) ([]model.TourDto, error) {
	if minPrice == 0.0 || maxPrice == 0.0 {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrInvalidRequest,
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: errors.New("invalid price"),
		}
	}
	res, err := t.DB.GetToursByPriceRange(minPrice, maxPrice)
	if err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrFailToGetTour,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err,
		}
	}
	if len(res) == 0 {
		return []model.TourDto{}, nil
	}
	return tourMapper.MapToursToDto(res), nil
}

func (t *Tour) GetToursByType(tourType model.TourType) ([]model.TourDto, error) {
	res, err := t.DB.GetToursByType(tourType.ToDto())
	if err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrFailToGetTour,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err,
		}
	}
	if len(res) == 0 {
		return []model.TourDto{}, nil
	}
	return tourMapper.MapToursToDto(res), nil
}

func (t *Tour) SearchTours(query string) ([]model.TourDto, error) {
	if query == "" {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrInvalidRequest,
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: errors.New("invalid query "),
		}
	}
	res, err := t.DB.SearchTours(query)
	if err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrFailToGetTour,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err,
		}
	}
	if len(res) == 0 {
		return []model.TourDto{}, nil
	}
	return tourMapper.MapToursToDto(res), nil
}

func (t *Tour) DeleteTour(id string) (*model.DeleteResponse, error) {
	if id == "" || strings.TrimSpace(id) == "" {
		return &model.DeleteResponse{}, &exception.TourWithUsError{
			Message:      exception.ErrInvalidTourID,
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: errors.New("invalid id"),
		}
	}
	if err := t.DB.DeleteTour(id); err != nil {
		return &model.DeleteResponse{}, &exception.TourWithUsError{
			Message:      exception.ErrFailToDeleteTour,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err,
		}
	}
	return &model.DeleteResponse{
		Success: true,
		Message: "Tour Deleted",
	}, nil
}

func (t *Tour) UpdateTour(id uint, dto model.UpdateTourDto) (*model.TourDto, error) {
	req := tourMapper.MapUpdateTourDtoToTourObject(&dto)
	tour, err := t.DB.UpdateTour(id, req)
	if err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrFailToUpdateTour,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err,
		}
	}
	res := tourMapper.MapObjectDtoToModelDto(&tour)
	return &res, nil
}

func (t *Tour) GetTourByTourOperator(operatorId string, tourId string) (model.TourDto, error) {
	res, err := t.DB.GetTourByTourOperator(operatorId, tourId)
	if err != nil {
		return model.TourDto{}, &exception.TourWithUsError{
			Message:      exception.ErrFailToGetTour,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err,
		}
	}
	resp := tourMapper.MapObjectDtoToModelDto(&res)
	return resp, err
}

func (t *Tour) GetToursByTourOperator(operatorId string) ([]model.TourDto, error) {
	res, err := t.DB.GetListOfToursByOperator(operatorId)
	if err != nil {
		return []model.TourDto{}, &exception.TourWithUsError{
			Message:      exception.ErrFailToGetTour,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err,
		}
	}
	if len(res) == 0 {
		return []model.TourDto{}, nil
	}
	return tourMapper.MapToursToDto(res), nil
}
