package tourDto

import "github.com/jinzhu/gorm"

type TourDto struct {
	OperatorID      string `json:"operator_id"`
	TourTitle       string `json:"tour_title"`
	MeetingPoint    string `json:"meeting_point"`
	Destination     string `json:"destination"`
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
	Availability    string `json:"availability"`
}

type TourObject struct {
	gorm.Model
	OperatorID      string `json:"operator_id"`
	TourTitle       string `json:"tour_title"`
	MeetingPoint    string `json:"meeting_point"`
	Destination     string `json:"destination"`
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
	Availability    string `json:"availability"`
}
type CreateTourResponse struct {
	Id        uint   `json:"id"`
	TourTitle string `json:"tour_title"`
	Date      string `json:"date"`
	Message   string `json:"message"`
	Status    bool   `json:"status"`
}
