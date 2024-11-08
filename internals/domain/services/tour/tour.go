package tour

import (
	usecase "github.com/Bigthugboy/TourWithUs/internals/application.port/tourWithUs.port/input/tourUseCaseInputPort"
	database "github.com/Bigthugboy/TourWithUs/internals/application.port/tourWithUs.port/output/repo/tourRepo"
	"github.com/Bigthugboy/TourWithUs/internals/domain/domainMapper"
	"github.com/Bigthugboy/TourWithUs/internals/domain/exception"
	"github.com/Bigthugboy/TourWithUs/internals/domain/model"
	"github.com/Bigthugboy/TourWithUs/internals/domain/services"
	"net/http"
)

type Tour struct {
	DB database.TourRepository
	UC usecase.TourUseCaseInputPort
}

func NewTour(db database.TourRepository, uc usecase.TourUseCaseInputPort) usecase.TourUseCaseInputPort {
	return &Tour{
		DB: db,
		UC: uc,
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
	req := domainMapper.MapModelTourDtoToTourDBObject(request)
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
		TourId:          res.TourId,
		OperatorContact: res.OperatorContact,
		Price:           res.Price,
	}
	return response, nil
}

func (t *Tour) GetTourById(id string) (*model.TourDto, error) {
	res, err := t.DB.GetTourById(id)
	if err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrFailToGetTour,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err,
		}
	}
	resp := domainMapper.MapObjectDtoToModelDto(&res)
	return &resp, nil
}

func (t Tour) GetAllTours() ([]model.TourDto, error) {
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
	return domainMapper.MapToursToDto(res), nil
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
	return domainMapper.MapToursToDto(res), nil
}

func (t *Tour) GetToursByLocation(location string) ([]model.TourDto, error) {
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
	return domainMapper.MapToursToDto(res), nil
}

func (t Tour) GetToursByDateRange(startDate, endDate string) ([]model.TourDto, error) {
	res, err := t.DB.GetToursByDateRange(startDate, endDate)
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
	return domainMapper.MapToursToDto(res), nil
}

func (t Tour) GetToursByPriceRange(minPrice, maxPrice float64) ([]model.TourDto, error) {
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
	return domainMapper.MapToursToDto(res), nil
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
	return domainMapper.MapToursToDto(res), nil
}

func (t *Tour) SearchTours(query string) ([]model.TourDto, error) {
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
	return domainMapper.MapToursToDto(res), nil
}

func (t *Tour) DeleteTour(id string) (*model.DeleteResponse, error) {
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

func (t *Tour) UpdateTour(id string, dto model.UpdateTourDto) (*model.TourDto, error) {
	req := domainMapper.MapUpdateTourDtoToTourObject(&dto)
	tour, err := t.DB.UpdateTour(id, req)
	if err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrFailToUpdateTour,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err,
		}
	}
	res := domainMapper.MapObjectDtoToModelDto(&tour)
	return &res, nil
}
