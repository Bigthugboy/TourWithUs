package tourDto

import "github.com/jinzhu/gorm"

type TourDto struct {
	OperatorID      string `json:"operator_id"`
	TourTitle       string `json:"tour_title"`
	Location        string `json:"location"`
	StartTime       string `json:"start_time"`
	LanguageOffered string `json:"language_offered"`
	NumberOfTourist string `json:"number_of_tourist"`
	Description     string `json:"description"`
	TourGuide       string `json:"tour_guide"`
	TourOperator    string `json:"tour_operator"`
	OperatorContact string `json:"operator_contact"`
	Category        string `json:"category"`
	Activity        string `json:"activity"`
	Date            string `json:"date"`
	Price           string `json:"price"`
	TouristEmail    string `json:"tourist_email"`
	Availability    bool   `json:"availability"`
}

type TourObject struct {
	gorm.Model
	OperatorID      string   `json:"operator_id"`
	TourTitle       string   `json:"tour_title"`
	Location        string   `json:"location"`
	StartTime       string   `json:"start_time"`
	LanguageOffered string   `json:"language_offered"`
	NumberOfTourist string   `json:"number_of_tourist"`
	Description     string   `json:"description"`
	TourGuide       string   `json:"tour_guide"`
	TourOperator    string   `json:"tour_operator"`
	OperatorContact string   `json:"operator_contact"`
	Category        string   `json:"category"`
	Activity        string   `json:"activity"`
	Date            string   `json:"date"`
	Price           string   `json:"price"`
	TouristEmail    string   `json:"tourist_email"`
	Availability    bool     `json:"availability"`
	TourType        TourType `json:"tour_type"`
}
type CreateTourResponse struct {
	TourId          uint   `json:"id"`
	TourTitle       string `json:"tour_title"`
	Date            string `json:"date"`
	Message         string `json:"message"`
	Status          bool   `json:"status"`
	Price           string `json:"price"`
	OperatorContact string `json:"operator_contact"`
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
