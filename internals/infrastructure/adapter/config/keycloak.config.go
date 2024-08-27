package config

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

var secretKey = []byte("404E635266556A586E3272357538782F413F4428472B4B6250645367566B5970")

type Keycloak struct {
	AuthKeyMutex sync.Mutex
	BearerToken  string
}

type Payload struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}
type TokenRes struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	Scope            string `json:"scope"`
}

type RegisterTouristPayload struct {
	Username      string `json:"username"`
	Enabled       bool   `json:"enabled"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"emailVerified"`
	Password      string `json:"password"`
	Credentials   []struct {
		Type      string `json:"type"`
		Value     string `json:"value"`
		Temporary bool   `json:"temporary"`
	} `json:"credentials"`
}

func GenerateToken(payload Payload) (*TokenRes, error) {
	form := url.Values{
		"client_id":     {payload.ClientId},
		"client_secret": {payload.ClientSecret},
		"grant_type":    {payload.GrantType},
		"username":      {payload.Username},
		"password":      {payload.Password},
	}
	encodedData := form.Encode()
	req, err := http.NewRequest("POST", "http://localhost:8080/realms/TourWithUs/protocol/openid-connect/token", strings.NewReader(encodedData))
	if err != nil {
		log.Println("Error creating request: ", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error performing request: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Non-OK HTTP status: ", resp.StatusCode)
		return nil, errors.New("something went wrong while connecting to Keycloak")
	}

	var tokenResponse TokenRes
	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)
	if err != nil {
		log.Println("Error decoding response: ", err)
		return nil, err
	}

	k := Keycloak{}
	k.AuthKeyMutex.Lock()
	k.BearerToken = tokenResponse.AccessToken
	k.AuthKeyMutex.Unlock()

	return &tokenResponse, nil

}

func Login(payload Payload) (*TokenRes, error) {
	ctx := context.Background()
	logrus.Info("login user ")
	var response TokenRes
	err := PostData(ctx, "http://localhost:8080/realms/TourWithUs/protocol/openid-connect/token", payload, &response)
	if err != nil {
		logrus.WithError(err).Error("Error creating request")
		return nil, err
	}
	return &response, nil
}

func SaveTourist(regPayload RegisterTouristPayload) (string, error) {

	kCreatePayload := RegisterTouristPayload{
		Username:  regPayload.Username,
		FirstName: regPayload.FirstName,
		LastName:  regPayload.LastName,
		Email:     regPayload.Email,
		Enabled:   true,
		Credentials: []struct {
			Type      string `json:"type"`
			Value     string `json:"value"`
			Temporary bool   `json:"temporary"`
		}{
			{
				Type:      "password",
				Value:     regPayload.Password,
				Temporary: false,
			},
		},
	}

	jsonData, err := json.Marshal(kCreatePayload)
	if err != nil {
		logrus.WithError(err).Error("Error marshaling JSON")
		return "", err
	}
	req, err := http.NewRequest("POST", "http://localhost:8080/admin/realms/TourWithUs/users", bytes.NewBuffer(jsonData))
	if err != nil {
		logrus.WithError(err).Error("Error creating HTTP request")
		return "", err
	}
	k := Keycloak{}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+k.BearerToken)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		logrus.WithError(err).Error("Error performing HTTP request")
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		logrus.Errorf("Non-OK HTTP status: %d", resp.StatusCode)
		return "", errors.New("something went wrong while connecting to Keycloak")
	}
	return "User created successfully", nil
}

func ValidateToken(token string) (bool, error) {
	keycloakPublicKey, err := FetchKeycloakPublicKey()
	if err != nil {
		return false, err
	}
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(keycloakPublicKey))
		if err != nil {
			return nil, err
		}
		return publicKey, nil
	})
	if err != nil {
		return false, err
	}
	return parsedToken.Valid, nil
}

func FetchKeycloakPublicKey() (string, error) {
	resp, err := http.Get("http://localhost:8080/realms/TourWithUs/protocol/openid-connect/certs")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Println("non-200 response: ", resp.StatusCode)
		return "", errors.New("something went wrong while fetching your key from keycloak")
	}
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var keycloakPublicKey string
	return keycloakPublicKey, nil
}

func PostData(ctx context.Context, url string, requestData interface{}, response interface{}) error {
	data, err := json.Marshal(requestData)
	if err != nil {
		logrus.WithError(err).Error("failed to marshal request")
		return err
	}
	request, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(string(data)))
	if err != nil {
		logrus.WithError(err).Error("failed to create request")
		return err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		logrus.WithError(err).Error("failed to send request")
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		logrus.WithFields(logrus.Fields{
			"status_code": resp.StatusCode,
			"response":    string(bodyBytes),
		}).Error("Received non-200 response")
		return err
	}
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to decode response")
		return err
	}
	return nil
}
