package tourOperatorRepo

import dto "github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/operatorDto"

type TourOperatorRepo interface {
	SaveTourOperator(object dto.OperatorDto) (dto.SavedOperatorRes, error)
	GetTourOperatorById(id string) (dto.OperatorDto, error)
	GetAllTourOperator() ([]dto.OperatorDto, error)
	GetTourOperatorByRating(limit int) ([]dto.OperatorDto, error)
	DeleteTourOperator(id string) (string, error)
	GetTourOperatorByEmail(email string) (dto.OperatorDto, error)
}
