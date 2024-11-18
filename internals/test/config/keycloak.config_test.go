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

func details() config2.RegisterTouristPayload {
	return config2.RegisterTouristPayload{
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
}
func Login() config2.LoginCredentials {
	return config2.LoginCredentials{
		Username: "deji2444444@gmail.com",
		Password: "damilola",
	}
}
func TestGenerateToken(t *testing.T) {
	client_id := "tourModel"
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
	res, err := config2.SaveTouristOnKeycloak(details())
	if err != nil {
		t.Error(err)
		logrus.Info(err.Error())
	}
	fmt.Println(res)
}

func TestSaveTourist_shouldThrowError_DuplicatedUser(t *testing.T) {
	res, err := config2.SaveTouristOnKeycloak(details())
	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestSaveTourist_shouldThrowError_InvalidUser(t *testing.T) {
	req := details()
	req.Email = ""
	logrus.Info("------->", req)
	res, err := config2.SaveTouristOnKeycloak(req)
	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestSaveTourist_shouldThrowError_invalidDetails(t *testing.T) {
	req := details()
	req.Username = ""
	_, err := config2.SaveTouristOnKeycloak(req)
	logrus.Info("err is ", err)
	assert.Error(t, err)

	req = details()
	req.Email = ""
	_, err = config2.SaveTouristOnKeycloak(req)
	logrus.Info("err is ", err)
	assert.Error(t, err)

	req = details()
	req.Email = "hgfnknspidgjg"
	_, err = config2.SaveTouristOnKeycloak(req)
	logrus.Info("err is ", err)
	assert.Error(t, err)
}

func TestSaveTourist_shouldThrowError_InvalidPassword(t *testing.T) {
	req := details()
	req.Email = "thugboy@gmail.com"
	req.Credentials = []config2.Credentials{
		{
			Type:      "password",
			Value:     "",
			Temporary: false,
		},
	}
	_, err := config2.SaveTouristOnKeycloak(req)
	assert.Error(t, err)
}

func TestSaveTourist_shouldThrowError_InvalidPasswordLength(t *testing.T) {
	req := details()
	req.Email = "thugboy@gmail.com"
	req.Credentials = []config2.Credentials{
		{
			Type:      "password",
			Value:     "passwor",
			Temporary: false,
		},
	}
	_, err := config2.SaveTouristOnKeycloak(req)
	assert.Error(t, err)
}

func TestLoginTourist(t *testing.T) {
	details := Login()
	res, err := config2.LoginUser(details)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestLoginTourist_shouldThrowError_InvalidCredentials(t *testing.T) {
	details := Login()
	details.Username = ""
	res, err := config2.LoginUser(details)
	assert.Error(t, err)
	assert.Empty(t, res)

	details.Username = "fhdmmdvjrg"
	res, err = config2.LoginUser(details)
	assert.Error(t, err)
	assert.Empty(t, res)

	details.Password = ""
	res, err = config2.LoginUser(details)
	assert.Error(t, err)
	assert.Empty(t, res)

	details.Password = "7459njf"
	res, err = config2.LoginUser(details)
	assert.Error(t, err)
	assert.Empty(t, res)
}
