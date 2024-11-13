package tourRepo

import (
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/tourDto"
	"time"
)

type TourRepository interface {
	CreateTour(object tourDto.TourObject) (tourDto.CreateTourResponse, error)
	GetAllTours() ([]tourDto.TourObject, error)
	GetTourById(id string) (tourDto.TourObject, error)
	GetAvailableTours() ([]tourDto.TourObject, error)
	GetToursByLocation(location string) ([]tourDto.TourObject, error)
	GetToursByDateRange(startDate, endDate time.Time) ([]tourDto.TourObject, error)
	GetToursByPriceRange(minPrice, maxPrice float64) ([]tourDto.TourObject, error)
	GetToursByType(tourType tourDto.TourType) ([]tourDto.TourObject, error)
	SearchTours(query string) ([]tourDto.TourObject, error)
	DeleteTour(id string) error
	UpdateTour(id string, updatedFields tourDto.UpdateTourDto) (tourDto.TourObject, error)
}
