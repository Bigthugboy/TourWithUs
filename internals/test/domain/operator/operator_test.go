package operator

import (
	"github.com/Bigthugboy/TourWithUs/internals/application.port/tourWithUs.port/input/tourUseCase"
	"github.com/Bigthugboy/TourWithUs/internals/application.port/tourWithUs.port/output/repo/tourOperatorRepo/internals/test/domain"
	"github.com/Bigthugboy/TourWithUs/internals/domain/model/tourOpModel"
	"github.com/Bigthugboy/TourWithUs/internals/domain/services/Operator"
	dto "github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/operatorDto"
	"github.com/golang/mock/gomock"
	"log"
	"testing"
)

var usecase tourUseCase.TourUseCaseInputPort

func operatorObject() dto.OperatorDto {
	return dto.OperatorDto{
		FirstName:   "john",
		LastName:    "smith",
		Email:       "john@gmail.com",
		PhoneNumber: "080123456789",
		Password:    "damilola",
	}

}

func TestOperatorCanBeCreated(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDb := domain.NewMockTourOperatorRepo(ctrl)
	service := Operator.NewOperatorService(mockDb, usecase)

	mockDb.EXPECT().GetTourOperatorByEmail("john@gmail.com").Return(dto.OperatorDto{}, nil)

	//details := operatorObject()

	operator := tourOpModel.TourOperator{
		FirstName:   "john",
		LastName:    "smith",
		Email:       "john@gmail.com",
		Password:    "damilola",
		PhoneNumber: "080123456789",
	}

	mockDb.EXPECT().SaveTourOperator(operatorObject()).Return(dto.SavedOperatorRes{
		Id:      1,
		Message: "saved tour operator",
		Email:   "john@gmail.com",
	}, nil)

	res, err := service.RegisterTourOperator(operator)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	log.Println(res)
}
