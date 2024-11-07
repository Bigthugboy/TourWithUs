package keycloakOutput_port

import (
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/touristDto"
)

type KeycloakOutPutPort interface {
	SaveTourist(tourist *touristDto.TouristDetails) (string, error)
	RetrieveTourist(details touristDto.RetrieveTourist) (string, error)
}
