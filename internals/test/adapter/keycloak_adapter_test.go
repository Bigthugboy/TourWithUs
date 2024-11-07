package adapter

import (
	config2 "github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/config"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/touristDto"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/output/keycloakAdapter"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var config = config2.Keycloak{}

func touristDetails() touristDto.TouristDetails {
	return touristDto.TouristDetails{
		FirstName: "paul",
		LastName:  "scott",
		Email:     "scott12345679@test.com",
		Username:  "dejiboy",
		Password:  "damilola",
	}
}
func TestSaveTourist(t *testing.T) {
	gin.SetMode(gin.TestMode)
	service := keycloakAdapter.NewKeycloakAdapter()
	details := touristDetails()
	res, err := service.SaveTourist(&details)
	if err != nil {
		t.Error(err)
	}
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestSaveTourist_CantSaveExistingUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	service := keycloakAdapter.NewKeycloakAdapter()
	details := touristDetails()
	res, err := service.SaveTourist(&details)
	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestSaveTourist_ValidateTouristDetails(t *testing.T) {
	gin.SetMode(gin.TestMode)
	service := keycloakAdapter.NewKeycloakAdapter()
	details := touristDetails()
	details.Email = ""
	res, err := service.SaveTourist(&details)
	log.Println("---", err)
	assert.Error(t, err)
	assert.Empty(t, res)

	details = touristDetails()
	details.Username = ""
	res, err = service.SaveTourist(&details)
	log.Println("---", err)
	assert.Error(t, err)
	assert.Empty(t, res)

	details = touristDetails()
	details.FirstName = ""
	res, err = service.SaveTourist(&details)
	log.Println("---", err)
	assert.Error(t, err)
	assert.Empty(t, res)

	details = touristDetails()
	details.LastName = ""
	res, err = service.SaveTourist(&details)
	log.Println("---", err)
	assert.Error(t, err)
	assert.Empty(t, res)

	details = touristDetails()
	details.Password = ""
	res, err = service.SaveTourist(&details)
	log.Println("---", err)
	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestSaveTourist_ValidateTouristEmail_InvalidEmail(t *testing.T) {
	gin.SetMode(gin.TestMode)
	service := keycloakAdapter.NewKeycloakAdapter()
	details := touristDetails()
	details.Email = "invalid@"
	res, err := service.SaveTourist(&details)
	log.Println("---", err)
	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestSaveTourist_PasswordLength(t *testing.T) {
	gin.SetMode(gin.TestMode)
	service := keycloakAdapter.NewKeycloakAdapter()
	details := touristDetails()
	details.Password = "dami"
	res, err := service.SaveTourist(&details)
	log.Println("---", err)
	assert.Error(t, err)
	assert.Empty(t, res)

	details = touristDetails()
	details.Password = "damilol"
	res, err = service.SaveTourist(&details)
	log.Println("---", err)
	assert.Error(t, err)
	assert.Empty(t, res)
}
