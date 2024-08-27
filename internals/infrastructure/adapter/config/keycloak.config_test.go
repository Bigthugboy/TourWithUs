package config

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
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
	client_secret := "5uUA2bACxvF65IJkUoRADosXeb4tbtDb"
	grant_type := "client_credentials"

	payload := Payload{
		ClientId:     client_id,
		ClientSecret: client_secret,
		GrantType:    grant_type,
		Username:     "admin",
		Password:     "admin",
	}
	fmt.Println(payload)
	tokenRes, err := GenerateToken(payload)
	assert.NoError(t, err)
	fmt.Println("your token is ---->{}:", tokenRes.AccessToken)
	assert.NotNil(t, tokenRes)
	fmt.Println(tokenRes)
	//assert.Equal(t, "", tokenRes.AccessToken)
}

func TestLogin(t *testing.T) {
	server := mockServer()
	defer server.Close()

	payload := Payload{
		ClientId:     "test-client-id",
		ClientSecret: "test-client-secret",
		GrantType:    "password",
		Username:     "test-username",
		Password:     "test-password",
	}
	tokenRes, err := Login(payload)
	assert.NoError(t, err)
	assert.NotNil(t, tokenRes)
	assert.Equal(t, "mock_access_token", tokenRes.AccessToken)
}

func TestValidateToken(t *testing.T) {

	valid, err := ValidateToken("invalid_token")
	assert.False(t, valid)
	assert.Error(t, err)
	valid, err = ValidateToken("mock_access_token")
	assert.False(t, valid)
	assert.Error(t, err)
}

func TestFetchKeycloakPublicKey(t *testing.T) {
	gin.SetMode(gin.TestMode)
	server := mockServer()
	defer server.Close()
	publicKey, err := FetchKeycloakPublicKey()
	assert.NoError(t, err)
	assert.Equal(t, "mock_public_key", publicKey)
}

func TestPostData(t *testing.T) {
	server := mockServer()
	defer server.Close()

	ctx := context.Background()
	payload := Payload{
		ClientId:     "test-client-id",
		ClientSecret: "test-client-secret",
		GrantType:    "password",
		Username:     "test-username",
		Password:     "test-password",
	}

	var response TokenRes
	err := PostData(ctx, server.URL+"/realms/test/protocol/openid-connect/token", payload, &response)
	assert.NoError(t, err)
	assert.Equal(t, "mock_access_token", response.AccessToken)
}
