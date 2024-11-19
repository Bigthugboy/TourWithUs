package services

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

func ValidateRequest(request interface{}) error {
	if request == nil || request == "" {
		return errors.New("invalid request")
	} else {
		validate := validator.New()
		err := validate.Struct(request)
		if err != nil {
			return err
		}
	}

	return nil
}
