package services

import (
	"fmt"
	"github.com/Bigthugboy/TourWithUs/internals/application.port/tourist.port/input/touristUseCaseInputPort"
	port "github.com/Bigthugboy/TourWithUs/internals/application.port/tourist.port/output/keycloakOutput.port"
	"github.com/Bigthugboy/TourWithUs/internals/application.port/tourist.port/output/repo"
	"github.com/Bigthugboy/TourWithUs/internals/domain/domainMapper"
	"github.com/Bigthugboy/TourWithUs/internals/domain/exception"
	"github.com/Bigthugboy/TourWithUs/internals/domain/model"
	"net/http"
)

type TouristUseCase struct {
	Db                 repo.DBStore
	keycloakOutPutPort port.KeycloakOutPutPort
}

func NewTourist(db repo.DBStore, keycloakAdapter port.KeycloakOutPutPort) touristUseCaseInputPort.TouristUseCase {
	return &TouristUseCase{
		Db:                 db,
		keycloakOutPutPort: keycloakAdapter,
	}
}

func (t *TouristUseCase) RegisterTouristUseCase(request *model.RegisterRequest) (*model.RegisterResponse, error) {
	if err := ValidateRequest(request); err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrValidatingRequest,
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: err,
		}
	}
	_, err := t.Db.SearchTouristByEmail(request.Email)
	if err == nil {
		return nil, fmt.Errorf("tourist with this email already exists")
	}
	Ktourist := domainMapper.MapRegisterRequestToTouristDetails(request)

	response, err := t.keycloakOutPutPort.SaveTourist(&Ktourist)
	if err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrSavingUser,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err,
		}
	}

	if response != "User created successfully" {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrFailToSaveUserToKeycloak,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: fmt.Errorf("unexpected response: %s", response),
		}
	}
	DTourist := domainMapper.MapRegisterRequestToTouristObject(request)
	savedTourist, id, err := t.Db.InsertTourist(DTourist)
	if err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrFailToSaveUserToDatabase,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err,
		}
	}
	return &model.RegisterResponse{
		ID:        id,
		FirstName: savedTourist.FirstName,
		LastName:  savedTourist.LastName,
		Email:     savedTourist.Email,
		Username:  savedTourist.Username,
	}, nil
}

func (t *TouristUseCase) Login(details model.LoginRequest) (*model.LoginResponse, error) {
	if err := ValidateRequest(details); err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrValidatingRequest,
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: err,
		}
	}
	fetchedUser, err := t.Db.SearchTouristByEmail(details.Email)
	if err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrUserNotFound,
			StatusCode:   http.StatusNotFound,
			ErrorMessage: fmt.Errorf("no user found with email %s: %w", details.Email, err),
		}
	}
	touristDetails := domainMapper.MapLoginRequestToTouristDetails(&details)

	_, err = t.keycloakOutPutPort.RetrieveTourist(touristDetails)
	if err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrFailToLoginUser,
			StatusCode:   http.StatusUnauthorized,
			ErrorMessage: fmt.Errorf("failed to authenticate user: %w", err),
		}
	}
	return &model.LoginResponse{
		Username: fetchedUser.Username,
		Message:  "logged in successfully",
	}, nil
}
