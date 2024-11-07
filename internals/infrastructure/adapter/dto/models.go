package dto

import (
	"github.com/jinzhu/gorm"
)

type TouristDetails struct {
	FirstName  string `json:"firstName" validate:"required"`
	LastName   string `json:"lastName" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=8"`
	ProfilePic string `json:"profilePic"`
	Username   string `json:"username"`
}

type TouristObject struct {
	gorm.Model
	ID               uint `gorm:"primaryKey"`
	FirstName        string
	LastName         string
	Email            string `gorm:"unique"`
	Password         string
	ProfilePic       string
	Username         string `gorm:"unique"`
	MotherMaidenName string
}

type RetrieveTourist struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
