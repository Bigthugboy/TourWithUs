package keycloakAdapter

import "github.com/Bigthugboy/TourWithUs/internals/domain/model"

type Keycloak struct {
}

type newKeycloak struct {
}

func (k *Keycloak) saveTourist(details model.TouristDetails) (model.TouristDetails, error) {

	return model.TouristDetails{}, nil
}
