package query

import (
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/output/persistence/Db/repo"
	"github.com/jinzhu/gorm"
)

type TourDB struct {
	DB *gorm.DB
}

func NewTourDB(db *gorm.DB) repo.DBStore {
	return &TourDB{
		DB: db,
	}
}
