package tourist

import (
	"errors"
	domain2 "github.com/Bigthugboy/TourWithUs/internals/application.port/tourWithUs.port/output/keycloakOutput.port/internals/test/domain"
	"github.com/Bigthugboy/TourWithUs/internals/application.port/tourWithUs.port/output/repo/internals/test/domain"
	model "github.com/Bigthugboy/TourWithUs/internals/domain/model/touristModel"
	"github.com/Bigthugboy/TourWithUs/internals/domain/services/tourist"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/touristDto"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterTouristUseCase_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := domain.NewMockDBStore(ctrl)
	mockKeycloakAdapter := domain2.NewMockKeycloakOutPutPort(ctrl)

	mockDB.EXPECT().SearchTouristByEmail("test@example.com").Return(touristDto.TouristObject{}, errors.New("not found"))
	mockKeycloakAdapter.EXPECT().SaveTourist(gomock.Any()).Return("User created successfully", nil)
	mockDB.EXPECT().InsertTourist(gomock.Any()).Return(&touristDto.TouristObject{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "test@example.com",
		Username:  "johndoe",
	}, int64(123), nil)

	usecase := tourist.NewTourist(mockDB, mockKeycloakAdapter)

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
	mockDB.EXPECT().SearchTouristByEmail("test@example.com").Return(touristDto.TouristObject{}, nil)
	useCase := tourist.NewTourist(mockDB, mockKeycloakAdapter)

	request := &model.RegisterRequest{
		Email:     "test@example.com",
		FirstName: "John",
		Password:  "damilola",
		LastName:  "Doe",
	}
	_, err := useCase.RegisterTouristUseCase(request)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "touristModel with this email already exists")
}
