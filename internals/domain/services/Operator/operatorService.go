package Operator

import (
	"fmt"
	"github.com/Bigthugboy/TourWithUs/internals/application.port/tourWithUs.port/input/tourOperatorUseCase"
	"github.com/Bigthugboy/TourWithUs/internals/application.port/tourWithUs.port/input/tourUseCase"
	"github.com/Bigthugboy/TourWithUs/internals/application.port/tourWithUs.port/output/repo/tourOperatorRepo"
	"github.com/Bigthugboy/TourWithUs/internals/domain/configs"
	"github.com/Bigthugboy/TourWithUs/internals/domain/domainMapper/tourOperatorMapper"
	"github.com/Bigthugboy/TourWithUs/internals/domain/exception"
	"github.com/Bigthugboy/TourWithUs/internals/domain/model/tourModel"
	model "github.com/Bigthugboy/TourWithUs/internals/domain/model/tourOpModel"
	"github.com/Bigthugboy/TourWithUs/internals/domain/services"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/output/OperatorSecurity"
	"log"
	"net/http"
)

type Service struct {
	Db      tourOperatorRepo.TourOperatorRepo
	UseCase tourUseCase.TourUseCaseInputPort
}

func NewOperatorService(DB tourOperatorRepo.TourOperatorRepo, usecase tourUseCase.TourUseCaseInputPort) tourOperatorUseCase.TourOperatorUseCase {
	return &Service{
		Db:      DB,
		UseCase: usecase,
	}
}

func (o *Service) RegisterTourOperator(operator model.TourOperator) (*model.CreateTourOperatorResponse, error) {
	if err := services.ValidateRequest(operator); err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrValidatingRequest,
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: err,
		}
	}
	_, err := o.Db.GetTourOperatorByEmail(operator.Email)
	if err == nil {
		return nil, &exception.TourWithUsError{
			Message:      "User with email already exists",
			StatusCode:   http.StatusConflict,
			ErrorMessage: err,
		}
	}
	hashedPassword, err := configs.Encrypt(operator.Password)
	if err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrEncryptingPassword,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err,
		}
	}
	operator.Password = hashedPassword
	touroperator := tourOperatorMapper.MapperOperatorDtoTOObject(operator)
	res, err := o.Db.SaveTourOperator(touroperator)
	if err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrFailToSaveUserToDatabase,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err,
		}
	}
	return &model.CreateTourOperatorResponse{
		Id:      res.Id,
		Message: "created successfully",
		Email:   res.Email,
	}, nil
}

func (o *Service) Login(req model.LoginRequest) (*model.LoginRes, error) {
	if err := services.ValidateRequest(req); err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrValidatingRequest,
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: err,
		}
	}
	fetchedOperator, err := o.Db.GetTourOperatorByEmail(req.Email)
	if err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrUserNotFound,
			StatusCode:   http.StatusNotFound,
			ErrorMessage: fmt.Errorf("no user found with email %s: %w", req.Email, err),
		}
	}
	if err := configs.ComparePasswords(fetchedOperator.Password, req.Password); err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrInvalidCredentials,
			StatusCode:   http.StatusUnauthorized,
			ErrorMessage: fmt.Errorf("invalid credentials: %w", err),
		}
	}

	generatedToken, refreshedToken, err := OperatorSecurity.Generate(fetchedOperator.Email, int64(fetchedOperator.ID))
	if err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrFailToGenerateToken,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err,
		}
	}
	return &model.LoginRes{
		Message:      "logged in Successfully",
		AccessToken:  generatedToken,
		RefreshToken: refreshedToken,
	}, nil
}

func (o *Service) CreateTour(tour tourModel.TourDto) (*tourModel.CreateTourResponse, error) {
	if err := services.ValidateRequest(tour); err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrValidatingRequest,
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: err,
		}
	}
	res, err := o.UseCase.CreateTour(&tour)
	if err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrFailToCreateTour,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err,
		}
	}
	log.Printf("Tour created successfully: %v", res)
	return res, nil
}

func (o *Service) UpdateTour(tourID string, tour tourModel.UpdateTourDto) (*tourModel.TourDto, error) {
	if err := services.ValidateRequest(tour); err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrValidatingRequest,
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: err,
		}
	}
	res, err := o.UseCase.UpdateTour(tourID, tour)
	if err != nil {
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrFailToUpdateTour,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err,
		}
	}
	log.Printf("Tour updated successfully: %v", res)
	return res, nil
}

func (o *Service) DeleteTour(tourID string) (string, error) {
	res, err := o.UseCase.DeleteTour(tourID)
	if err != nil {
		log.Printf("Failed to delete tour with ID %s: %v", tourID, err)
		return "", &exception.TourWithUsError{
			Message:      exception.ErrFailToDeleteTour,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err,
		}
	}
	log.Printf("Tour with ID %s deleted successfully", tourID)

	return res.Message, nil
}

func (o *Service) ViewTourDetails(tourID string) (*tourModel.TourDto, error) {
	res, err := o.UseCase.GetTourById(tourID)
	if err != nil {
		log.Printf("Failed to get tour with ID %s: %v", tourID, err)
		return nil, &exception.TourWithUsError{
			Message:      exception.ErrFailToGetTour,
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: err,
		}
	}
	log.Printf("Tour with ID %s retrieved successfully: %+v", tourID, res)
	return res, nil
}

func (o *Service) ListTours(operatorID string) ([]tourModel.TourDto, error) {
	//TODO implement me
	panic("implement me")
}

func (o *Service) ConfirmBooking(bookingID string) error {
	//TODO implement me
	panic("implement me")
}

func (o *Service) CancelBooking(bookingID string) error {
	//TODO implement me
	panic("implement me")
}

func (o *Service) ManageAvailability(tourID string, availability bool) error {
	//TODO implement me
	panic("implement me")
}

func (o *Service) UpdateProfile(operator model.TourOperator) (*model.CreateTourOperatorResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (o *Service) ChangePassword(req model.ChangePasswordRequest) error {
	//TODO implement me
	panic("implement me")
}

func (o *Service) Logout(operatorID string) error {
	//TODO implement me
	panic("implement me")
}
