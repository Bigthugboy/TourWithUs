package config

import (
	"fmt"
	config2 "github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/config"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockServer() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/realms/TourWithUS/protocol/openid-connect/token", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"accessToken": "mock_access_token"}`))
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	handler.HandleFunc("/realms/TourWithUs/protocol/openid-connect/certs", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"publicKey": "mock_public_key"}`))
	})

	return httptest.NewServer(handler)
}

func TestGenerateToken(t *testing.T) {
	client_id := "tour"
	client_secret := "SjrSFWLqOzRVa36FC5SI6sdBDfc7AjJk"
	grant_type := "client_credentials"
	service := config2.Keycloak{}

	payload := config2.Payload{
		ClientId:     client_id,
		ClientSecret: client_secret,
		GrantType:    grant_type,
		Username:     "admin",
		Password:     "admin",
	}
	fmt.Println(payload)
	tokenRes, err := service.GenerateToken(payload)
	assert.NoError(t, err)
	fmt.Println("your token is ---->{}:", tokenRes.AccessToken)
	assert.NotNil(t, tokenRes)
	fmt.Println(tokenRes)

}

func TestSavedTourist(t *testing.T) {
	kCreatePayload := config2.RegisterTouristPayload{
		Username:      "shakur",
		FirstName:     "deji",
		LastName:      "ayo",
		Email:         "deji2444444@gmail.com",
		Enabled:       true,
		EmailVerified: false,
		Credentials: []config2.Credentials{
			{
				Type:      "password",
				Value:     "damilola",
				Temporary: false,
			},
		},
		Attributes: map[string]string{
			"customAttribute": "test_value",
		},
	}
	res, err := config2.SaveTouristOnKeycloak(kCreatePayload)
	if err != nil {
		t.Error(err)
		logrus.Info(err.Error())
	}
	fmt.Println(res)
}
