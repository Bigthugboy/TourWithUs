package query

import (
	"github.com/Bigthugboy/TourWithUs/internals/application.port/tourWithUs.port/output/repo/touristRepo"
	"github.com/jinzhu/gorm"
)

type TourDB struct {
	DB *gorm.DB
}

func NewTourDB(db *gorm.DB) touristRepo.DBStore {
	return &TourDB{
		DB: db,
	}
}
