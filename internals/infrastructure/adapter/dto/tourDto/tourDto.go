package tourDto

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TourDto struct {
	OperatorID      string    `json:"operator_id"`
	TourTitle       string    `json:"tour_title"`
	Location        string    `json:"location"`
	Duration        string    `json:"duration"`
	LanguageOffered string    `json:"language_offered"`
	NumberOfTourist string    `json:"number_of_tourist"`
	Description     string    `json:"description"`
	TourGuide       string    `json:"tour_guide"`
	OperatorContact string    `json:"operator_contact"`
	Activity        string    `json:"activity"`
	Date            time.Time `json:"date"`
	Price           float64   `json:"price"`
	TouristEmail    string    `json:"tourist_email"`
	Availability    bool      `json:"availability"`
	StartDate       string    `json:"start_date"`
	EndDate         string    `json:"end_date"`
}

type TourObject struct {
	gorm.Model
	OperatorID      string    `json:"operator_id"`
	TourTitle       string    `json:"tour_title"`
	Location        string    `json:"location"`
	Duration        string    `json:"duration"`
	LanguageOffered string    `json:"language_offered"`
	NumberOfTourist string    `json:"number_of_tourist"`
	Description     string    `json:"description"`
	TourGuide       string    `json:"tour_guide"`
	OperatorContact string    `json:"operator_contact"`
	Activity        string    `json:"activity"`
	Date            time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Price           float64   `json:"price"`
	TouristEmail    string    `json:"tourist_email"`
	Availability    bool      `json:"availability"`
	TourType        TourType  `json:"tour_type"`
	StartDate       string    `json:"start_date"`
	EndDate         string    `json:"end_date"`
}
type CreateTourResponse struct {
	TourId          uint    `json:"id"`
	TourTitle       string  `json:"tour_title"`
	Date            string  `json:"date"`
	Message         string  `json:"message"`
	Status          bool    `json:"status"`
	Price           float64 `json:"price"`
	OperatorContact string  `json:"operator_contact"`
}
type TourType string

const (
	Adventure  TourType = "Adventure"
	Cultural   TourType = "Cultural"
	Historical TourType = "Historical"
	Nature     TourType = "Nature"
	Relaxation TourType = "Relaxation"
)

type UpdateTourDto struct {
	OperatorID      *string  `json:"operator_id"`
	TourTitle       *string  `json:"tour_title"`
	Location        *string  `json:"location"`
	Duration        *string  `json:"duration"`
	LanguageOffered *string  `json:"language_offered"`
	NumberOfTourist *string  `json:"number_of_tourist"`
	Description     *string  `json:"description"`
	TourGuide       *string  `json:"tour_guide"`
	OperatorContact *string  `json:"operator_contact"`
	Activity        *string  `json:"activity"`
	Price           *float64 `json:"price"`
	TouristEmail    *string  `json:"tourist_email"`
	Availability    *bool    `json:"availability"`
	StartDate       *string  `json:"start_date"`
	EndDate         *string  `json:"end_date"`
}
