package keycloakOutput_port

import (
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto"
)

type KeycloakOutPutPort interface {
	SaveTourist(tourist *dto.TouristDetails) (string, error)
	RetrieveTourist(details dto.RetrieveTourist) (string, error)
}
