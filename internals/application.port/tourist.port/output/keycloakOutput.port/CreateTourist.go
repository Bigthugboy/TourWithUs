package keycloakOutput_port

import (
	"github.com/Bigthugboy/TourWithUs/internals/domain/model"
)

type TouristOutPutPort interface {
	CreateTourist(tourist *model.TouristDetails) (*model.TouristDetails, error)
}
