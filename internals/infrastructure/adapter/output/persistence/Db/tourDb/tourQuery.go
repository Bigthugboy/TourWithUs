package tourDb

import (
	"errors"
	"fmt"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/tourDto"
)

func (t *TourRepositories) CreateTour(tourObject tourDto.TourObject) (tourDto.CreateTourResponse, error) {
	if t.DB == nil {
		return tourDto.CreateTourResponse{}, errors.New("tour DB Not Initialized")
	}
	if err := t.DB.Create(&tourObject).Error; err != nil {
		return tourDto.CreateTourResponse{}, fmt.Errorf("exception inserting new tourist: %w", err)
	}
	var response tourDto.CreateTourResponse
	response.Id = tourObject.ID
	response.TourTitle = tourObject.TourTitle
	response.Message = "tour successfully created"
	response.Date = tourObject.Date
	response.Status = true
	return response, nil
}
func (t *TourRepositories) GetAllTours() ([]tourDto.TourObject, error) {
	if t.DB == nil {
		return nil, errors.New("tour DB Not Initialized")
	}
	var tours []tourDto.TourObject
	if err := t.DB.Find(&tours).Error; err != nil {
		return nil, fmt.Errorf("exception getting all tours: %w", err)
	}
	return tours, nil

}
func (t *TourRepositories) GetTourById(id string) (tourDto.TourObject, error) {
	if t.DB == nil {
		return tourDto.TourObject{}, errors.New("tour DB Not Initialized")
	}
	var tour tourDto.TourObject
	if err := t.DB.Where("id = ?", id).First(&tour).Error; err != nil {
		return tourDto.TourObject{}, fmt.Errorf("exception getting tour: %w", err)
	}
	return tour, nil
}

func (t *TourRepositories) GetAvailableTours() ([]tourDto.TourObject, error) {
	if t.DB == nil {
		return nil, errors.New("tour DB Not Initialized")
	}
	var tours []tourDto.TourObject
	if err := t.DB.Where("Availability = ?", true).Find(&tours).Error; err != nil {
		return []tourDto.TourObject{}, fmt.Errorf("exception getting tour: %w", err)
	}
	return tours, nil

}

func (t *TourRepositories) GetToursByLocation(location string) ([]tourDto.TourObject, error) {
	if t.DB == nil {
		return nil, errors.New("tour DB Not Initialized")
	}
	var tours []tourDto.TourObject
	if err := t.DB.Where("location = ?", location).Find(&tours).Error; err != nil {
		return nil, fmt.Errorf("exception getting tours by location: %w", err)
	}
	return tours, nil
}

func (t *TourRepositories) GetToursByDateRange(startDate, endDate string) ([]tourDto.TourObject, error) {
	if t.DB == nil {
		return nil, errors.New("tour DB Not Initialized")
	}
	var tours []tourDto.TourObject
	if err := t.DB.Where("date BETWEEN ? AND ?", startDate, endDate).Find(&tours).Error; err != nil {
		return nil, fmt.Errorf("exception getting tours by date range: %w", err)
	}
	return tours, nil
}

func (t *TourRepositories) GetToursByPriceRange(minPrice, maxPrice float64) ([]tourDto.TourObject, error) {
	if t.DB == nil {
		return nil, errors.New("tour DB Not Initialized")
	}
	var tours []tourDto.TourObject
	if err := t.DB.Where("price BETWEEN ? AND ?", minPrice, maxPrice).Find(&tours).Error; err != nil {
		return nil, fmt.Errorf("exception getting tours by price range: %w", err)
	}
	return tours, nil
}
func (t *TourRepositories) SearchTours(query string) ([]tourDto.TourObject, error) {
	if t.DB == nil {
		return nil, errors.New("tour DB Not Initialized")
	}

	var tours []tourDto.TourObject
	if err := t.DB.Where("tour_title LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%").Find(&tours).Error; err != nil {
		return nil, fmt.Errorf("exception searching tours: %w", err)
	}

	return tours, nil
}

func (t *TourRepositories) DeleteTour(id string) error {
	if t.DB == nil {
		return errors.New("tour DB Not Initialized")
	}

	if err := t.DB.Delete(&tourDto.TourObject{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("exception deleting tour: %w", err)
	}

	return nil
}

func (t *TourRepositories) UpdateTour(id string, updatedObject tourDto.TourObject) (tourDto.TourObject, error) {
	if t.DB == nil {
		return tourDto.TourObject{}, errors.New("tour DB Not Initialized")
	}

	var tour tourDto.TourObject
	result := t.DB.Model(&tour).Where("id = ?", id).Updates(updatedObject)
	if err := result.Error; err != nil {
		return tourDto.TourObject{}, fmt.Errorf("exception updating tour: %w", err)
	}

	return tour, nil
}

func (t *TourRepositories) GetToursByType(tourType string) ([]tourDto.TourObject, error) {
	if t.DB == nil {
		return nil, errors.New("tour DB Not Initialized")
	}
	var tours []tourDto.TourObject
	if err := t.DB.Where("category = ?", tourType).Find(&tours).Error; err != nil {
		return nil, fmt.Errorf("exception getting tours by type: %w", err)
	}

	return tours, nil
}
