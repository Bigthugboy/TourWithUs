package repo

import (
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto"
)

type DBStore interface {
	InsertTourist(user dto.TouristObject) (*dto.TouristObject, int64, error)
	SearchTouristByEmail(email string) (dto.TouristObject, error)
	GetTouristByID(userID string) (dto.TouristObject, error)
	GetAllTourists() ([]dto.TouristObject, error)
	DeleteTouristByID(userID string) error
	DeleteTouristByEmail(email string) error
}
