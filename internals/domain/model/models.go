package model

import "github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/tourDto"

type RegisterRequest struct {
	FirstName  string `json:"firstName" validate:"required"`
	LastName   string `json:"lastName" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=8"`
	ProfilePic string `json:"profilePic"`
	Username   string `json:"username"`
}

type RegisterResponse struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName" `
	Email     string `json:"email" `
	Username  string `json:"username"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
type LoginResponse struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type TourDto struct {
	OperatorID      string   `json:"operator_id"`
	TourTitle       string   `json:"tour_title"`
	Location        string   `json:"location"`
	Destination     string   `json:"destination"`
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
	Message         string `json:"message"`
	TourId          uint   `json:"tour_id"`
	OperatorContact string `json:"operator_Contact"`
	Price           string `json:"price"`
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
	OperatorID      *string `json:"operator_id"`
	TourTitle       *string `json:"tour_title"`
	Location        *string `json:"location"`
	Destination     *string `json:"destination"`
	StartTime       *string `json:"start_time"`
	LanguageOffered *string `json:"language_offered"`
	NumberOfTourist *string `json:"number_of_tourist"`
	Description     *string `json:"description"`
	TourGuide       *string `json:"tour_guide"`
	TourOperator    *string `json:"tour_operator"`
	OperatorContact *string `json:"operator_contact"`
	Category        *string `json:"category"`
	Activity        *string `json:"activity"`
	Date            *string `json:"date"`
	Price           *string `json:"price"`
	TouristEmail    *string `json:"tourist_email"`
	Availability    *bool   `json:"availability"`
}
