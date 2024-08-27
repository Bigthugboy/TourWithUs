package repo

import (
	"github.com/Bigthugboy/TourWithUs/internals/domain/model"
)

type DBStore interface {
	InsertTourist(user model.TouristDetails) (int64, error)
	SearchTouristByEmail(email string) (model.TouristDetails, error)
	GetTouristByID(userID string) (model.TouristDetails, error)
	GetAllTourists() ([]model.TouristDetails, error)
	DeleteTouristByID(userID string) error
	DeleteTouristByEmail(email string) error
}
