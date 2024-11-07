package tourDb

import (
	"github.com/Bigthugboy/TourWithUs/internals/application.port/tourWithUs.port/output/repo/tourRepo"
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
