package tourModel

import (
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/tourDto"
	"time"
)

type TourDto struct {
	OperatorID      string    `json:"operator_id"validate:"required"`
	TourTitle       string    `json:"tour_title"validate:"required"`
	Location        string    `json:"location"validate:"required"`
	Duration        string    `json:"duration"validate:"required"`
	LanguageOffered string    `json:"language_offered"`
	NumberOfTourist string    `json:"number_of_tourist"`
	Description     string    `json:"description"`
	TourGuide       string    `json:"tour_guide"validate:"required"`
	OperatorContact string    `json:"operator_contact"validate:"required"`
	Activity        string    `json:"activity"`
	Date            time.Time `json:"date"`
	Price           float64   `json:"price"validate:"required"`
	TouristEmail    string    `json:"tourist_email"`
	Availability    bool      `json:"availability"`
	TourType        TourType  `json:"tour_type"`
	StartDate       time.Time `json:"start_date"validate:"required"`
	EndDate         time.Time `json:"end_date"validate:"required"`
}

type CreateTourResponse struct {
	Message         string  `json:"message"`
	TourTitle       string  `json:"tour_title"`
	TourId          uint    `json:"tour_id"`
	OperatorContact string  `json:"operator_Contact"`
	Price           float64 `json:"price"`
}
type TourType string

const (
	Adventure  TourType = "Adventure"
	Cultural   TourType = "Cultural"
	Historical TourType = "Historical"
	Nature     TourType = "Nature"
	Relaxation TourType = "Relaxation"
	// Add more types as needed
)

func (t TourType) ToDto() tourDto.TourType {
	return tourDto.TourType(t)
}

type DeleteResponse struct {
	Success bool
	Message string
}

type UpdateTourDto struct {
	OperatorID      *string    `json:"operator_id"`
	TourTitle       *string    `json:"tour_title"`
	Location        *string    `json:"location"`
	Duration        *string    `json:"duration"`
	LanguageOffered *string    `json:"language_offered"`
	NumberOfTourist *string    `json:"number_of_tourist"`
	Description     *string    `json:"description"`
	TourGuide       *string    `json:"tour_guide"`
	OperatorContact *string    `json:"operator_contact"`
	Activity        *string    `json:"activity"`
	Price           *float64   `json:"price"`
	TouristEmail    *string    `json:"tourist_email"`
	Availability    *bool      `json:"availability"`
	StartDate       *time.Time `json:"start_date"`
	EndDate         *time.Time `json:"end_date"`
}
