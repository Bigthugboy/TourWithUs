package OperatorDb

import (
	"github.com/Bigthugboy/TourWithUs/internals/application.port/tourWithUs.port/output/repo/tourOperatorRepo"
	"github.com/jinzhu/gorm"
)

type OperatorDb struct {
	DB *gorm.DB
}

func NewOperatorDb(db *gorm.DB) tourOperatorRepo.TourOperatorRepo {
	return &OperatorDb{DB: db}

}
