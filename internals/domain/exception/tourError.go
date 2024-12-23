package exception

import (
	"fmt"
)

const (
	ErrMakingRequest            = "Error making request"
	ErrSavingUser               = "Error saving user"
	ErrFetchingUser             = "Error fetching user"
	ErrCreatingUser             = "Error creating user"
	ErrInvalidRequestData       = "Invalid request data"
	ErrUserAlreadyExists        = "User already exists"
	ErrFailToSaveUser           = "Failed to save user"
	ErrValidatingRequest        = "Error validating request"
	ErrFailToGetUser            = "Error getting user"
	ErrFailToSaveUserToKeycloak = "Failed to save user to keycloak"
	ErrFailToSaveUserToDatabase = "Failed to save user to database"
	ErrFailTOLoginUser          = "Failed to login user"
	ErrUserNotFound             = "User not found"
	ErrFailToLoginUser          = "Failed to login user"
	ErrFailToCreateTour         = "Failed to create tourModel"
	ErrFailToGetTour            = "Failed to get tourModel"
	ErrFailToDeleteTour         = "Failed to delete tourModel"
	ErrFailToUpdateTour         = "Failed to update tourModel"
	ErrInvalidTourID            = "Invalid Tour ID"
	ErrInvalidRequest           = "Invalid request"
	ErrFailToGenerateToken      = "Failed to generate token"
	ErrInvalidCredentials       = "Invalid credentials"
	ErrEncryptingPassword       = "Error encrypting password"
)

type TourWithUsError struct {
	Message      string
	StatusCode   int
	ErrorMessage error
}

func (e *TourWithUsError) Error() string {
	return fmt.Sprintf("%s (Status: %d), %v", e.Message, e.StatusCode, e.ErrorMessage)
}
