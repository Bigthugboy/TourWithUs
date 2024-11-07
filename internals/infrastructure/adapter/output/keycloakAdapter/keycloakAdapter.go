package keycloakAdapter

import (
	"errors"
	"fmt"
	keycloakoutputport "github.com/Bigthugboy/TourWithUs/internals/application.port/tourWithUs.port/output/keycloakOutput.port"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/config"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/touristDto"
	"github.com/go-playground/validator/v10"
	"strings"
)

type KeycloakAdapter struct {
	KeycloakConfig *config.Keycloak
}

func NewKeycloakAdapter() keycloakoutputport.KeycloakOutPutPort {

	return &KeycloakAdapter{}
}

func (s *KeycloakAdapter) SaveTourist(details *touristDto.TouristDetails) (string, error) {
	if err := ValidateRequest(details); err != nil {
		return "", fmt.Errorf("validation exception: %w", err)
	}
	regPayload := config.RegisterTouristPayload{
		Username:      details.Username,
		FirstName:     details.FirstName,
		LastName:      details.LastName,
		Email:         details.Email,
		Enabled:       true,
		EmailVerified: true,
		Credentials: []config.Credentials{
			{
				Type:      "password",
				Value:     details.Password,
				Temporary: false,
			},
		},
	}
	result, err := config.SaveTouristOnKeycloak(regPayload)
	if err != nil {
		return "", fmt.Errorf("failed to save tourist in KeycloakAdapter: %w", err)
	}
	return result, nil
}

func (s *KeycloakAdapter) RetrieveTourist(details touristDto.RetrieveTourist) (string, error) {
	Username := details.Email
	Password := details.Password

	result, err := config.LoginUser(Username, Password)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve tourist in KeycloakAdapter: %w", err)
	}
	return result, nil
}

func ValidateRequest(details *touristDto.TouristDetails) error {
	if details == nil {
		return errors.New("invalid request: details cannot be nil")
	}
	validate := validator.New()
	if err := validate.Struct(details); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("validation failed on field '%s' with condition '%s'", err.Field(), err.Tag()))
		}
		return fmt.Errorf("validation errors: %s", strings.Join(validationErrors, ", "))
	}
	return nil
}
