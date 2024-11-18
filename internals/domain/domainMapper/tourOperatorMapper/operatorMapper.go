package tourOperatorMapper

import (
	model "github.com/Bigthugboy/TourWithUs/internals/domain/model/tourOpModel"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/operatorDto"
)

func MapperOperatorDtoTOObject(operator model.TourOperator) operatorDto.OperatorDto {
	return operatorDto.OperatorDto{
		FirstName:   operator.FirstName,
		LastName:    operator.LastName,
		Email:       operator.Email,
		PhoneNumber: operator.PhoneNumber,
		Password:    operator.Password,
	}
}
