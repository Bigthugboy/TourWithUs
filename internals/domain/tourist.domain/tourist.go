package tourist_domain

import (
	"github.com/Bigthugboy/TourWithUs/internals/application.port/tourist.port/input/touristUseCaseIntputPort"
	"github.com/Bigthugboy/TourWithUs/internals/application.port/tourist.port/output/keycloakOutput.port"
	"github.com/Bigthugboy/TourWithUs/internals/domain/model"
)

type Tourist struct {
	useCase touristUseCaseIntputPort.TouristUseCase
	output  keycloakOutput_port.TouristOutPutPort
}

func NewTourist(useCase touristUseCaseIntputPort.TouristUseCase, output keycloakOutput_port.TouristOutPutPort) *Tourist {
	return &Tourist{
		useCase: useCase,
		output:  output,
	}
}

func (t *Tourist) RegisterTouristUseCase(tourist *model.TouristDetails) (*model.TouristDetails, error) {
	regTourist, err := t.output.CreateTourist(tourist)
	if err != nil {
		return nil, err
	}
	return regTourist, nil
}
