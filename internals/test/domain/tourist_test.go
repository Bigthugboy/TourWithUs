package domain

import (
	"errors"
	domain2 "github.com/Bigthugboy/TourWithUs/internals/application.port/tourist.port/output/keycloakOutput.port/internals/test/domain"
	"github.com/Bigthugboy/TourWithUs/internals/application.port/tourist.port/output/repo/internals/test/domain"
	"github.com/Bigthugboy/TourWithUs/internals/domain/model"
	"github.com/Bigthugboy/TourWithUs/internals/domain/services"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterTouristUseCase_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := domain.NewMockDBStore(ctrl)
	mockKeycloakAdapter := domain2.NewMockKeycloakOutPutPort(ctrl)

	mockDB.EXPECT().SearchTouristByEmail("test@example.com").Return(dto.TouristObject{}, errors.New("not found"))
	mockKeycloakAdapter.EXPECT().SaveTourist(gomock.Any()).Return("User created successfully", nil)
	mockDB.EXPECT().InsertTourist(gomock.Any()).Return(&dto.TouristObject{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "test@example.com",
		Username:  "johndoe",
	}, int64(123), nil)

	usecase := services.NewTourist(mockDB, mockKeycloakAdapter)

	request := &model.RegisterRequest{
		Email:     "test@example.com",
		FirstName: "John",
		Password:  "damilola",
		LastName:  "Doe",
	}

	response, err := usecase.RegisterTouristUseCase(request)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, int64(123), response.ID)
	assert.Equal(t, "John", response.FirstName)
	assert.Equal(t, "Doe", response.LastName)
	assert.Equal(t, "test@example.com", response.Email)
}
func TestRegisterTouristUseCase_UserAlreadyExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := domain.NewMockDBStore(ctrl)
	mockKeycloakAdapter := domain2.NewMockKeycloakOutPutPort(ctrl)
	mockDB.EXPECT().SearchTouristByEmail("test@example.com").Return(dto.TouristObject{}, nil)
	useCase := services.NewTourist(mockDB, mockKeycloakAdapter)

	request := &model.RegisterRequest{
		Email:     "test@example.com",
		FirstName: "John",
		Password:  "damilola",
		LastName:  "Doe",
	}
	_, err := useCase.RegisterTouristUseCase(request)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "tourist with this email already exists")
}
