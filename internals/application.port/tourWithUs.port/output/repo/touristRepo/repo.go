package touristRepo

import (
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/touristDto"
)

type DBStore interface {
	InsertTourist(user touristDto.TouristObject) (*touristDto.TouristObject, int64, error)
	SearchTouristByEmail(email string) (touristDto.TouristObject, error)
	GetTouristByID(userID string) (touristDto.TouristObject, error)
	GetAllTourists() ([]touristDto.TouristObject, error)
	DeleteTouristByID(userID string) error
	DeleteTouristByEmail(email string) error
}
