package tourDb

import (
	"github.com/Bigthugboy/TourWithUs/internals/application.port/tourist.port/output/tour/tourRepo"
	"github.com/jinzhu/gorm"
)

type TourRepositories struct {
	DB *gorm.DB
}

func NewTourDB(db *gorm.DB) tourRepo.TourRepository {
	return &TourRepositories{
		DB: db,
	}
}
