package tourUseCase

import "github.com/Bigthugboy/TourWithUs/internals/domain/model/tourModel"

type TourUseCaseInputPort interface {
	CreateTour(request *tourModel.TourDto) (*tourModel.CreateTourResponse, error)
	GetTourById(id string) (*tourModel.TourDto, error)
	GetAllTours() ([]tourModel.TourDto, error)
	GetAvailableTours() ([]tourModel.TourDto, error)
	GetToursByLocation(location string) ([]tourModel.TourDto, error)
	GetToursByDateRange(startDate, endDate string) ([]tourModel.TourDto, error)
	GetToursByPriceRange(minPrice, maxPrice float64) ([]tourModel.TourDto, error)
	GetToursByType(tourType tourModel.TourType) ([]tourModel.TourDto, error)
	SearchTours(query string) ([]tourModel.TourDto, error)
	DeleteTour(id string) (*tourModel.DeleteResponse, error)
	UpdateTour(id string, dto tourModel.UpdateTourDto) (*tourModel.TourDto, error)
	GetTourByTourOperator(tour tourModel.TourDto) (tourModel.TourDto, error)
	GetToursByTourOperator(tour tourModel.TourDto) ([]tourModel.TourDto, error)
}
