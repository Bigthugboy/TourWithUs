package operatorDto

import "github.com/jinzhu/gorm"

type OperatorDto struct {
	gorm.Model
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
	Rating      int
	Password    string
}

type SavedOperatorRes struct {
	Id      uint   `json:"id"`
	Message string `json:"message"`
	Email   string `json:"email"`
}
