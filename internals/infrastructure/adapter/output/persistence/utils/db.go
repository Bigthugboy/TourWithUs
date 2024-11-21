package utils

import (
	_ "database/sql"
	"fmt"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/tourDto"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

func NewDatabaseConnection() (*gorm.DB, error) {
	dbUser := "root"
	dbPassword := "damilola"
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "TourWithUs"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	db.AutoMigrate(tourDto.TourObject{})

	log.Println("Database connection established.")
	return db, nil
}

//
//func TourDatabaseConnection() (*gorm.DB, error) {
//	dbUser := "root"
//	dbPassword := "damilola"
//	dbHost := "localhost"
//	dbPort := "3307"
//	dbName := "TourDb"
//
//	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
//	db, err := gorm.Open("mysql", dsn)
//	if err != nil {
//		return nil, fmt.Errorf("failed to open database connection: %w", err)
//	}
//	db.AutoMigrate(&touristDto.TouristDetails{})
//
//	log.Println("Database connection established.")
//	return db, nil
//}
