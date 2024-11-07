package tourUseCaseInputPort

import "github.com/Bigthugboy/TourWithUs/internals/domain/model"

type TourUseCaseInputPort interface {
	CreateTour(request *model.TourDto) (*model.CreateTourResponse, error)
	GetTourById(id string) (*model.TourDto, error)
	GetAllTours() ([]model.TourDto, error)
	GetAvailableTours() ([]model.TourDto, error)
	GetToursByLocation(location string) ([]model.TourDto, error)
	GetToursByDateRange(startDate, endDate string) ([]model.TourDto, error)
	GetToursByPriceRange(minPrice, maxPrice float64) ([]model.TourDto, error)
	GetToursByType(tourType model.TourType) ([]model.TourDto, error)
	SearchTours(query string) ([]model.TourDto, error)
	DeleteTour(id string) (string, error)
	UpdateTour(id string) (*model.TourDto, error)
}
