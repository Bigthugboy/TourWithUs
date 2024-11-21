package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

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

type Credentials struct {
	Type      string `json:"type"`
	Value     string `json:"value" validate:"required,min=8"`
	Temporary bool   `json:"temporary"`
}

type RegisterTouristPayload struct {
	Username      string            `json:"username" validate:"required"`
	FirstName     string            `json:"firstName" validate:"required"`
	LastName      string            `json:"lastName" validate:"required"`
	Email         string            `json:"email" validate:"required,email"`
	Enabled       bool              `json:"enabled"`
	EmailVerified bool              `json:"emailVerified"`
	Credentials   []Credentials     `json:"credentials"`
	Attributes    map[string]string `json:"attributes,omitempty"`
}

type LoginCredentials struct {
	Username string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (k *Keycloak) GenerateToken(payload Payload) (*TokenRes, error) {
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
	k.AuthKeyMutex.Lock()
	k.BearerToken = tokenResponse.AccessToken
	k.AuthKeyMutex.Unlock()

	return &tokenResponse, nil

}

func SaveTouristOnKeycloak(regPayload RegisterTouristPayload) (string, error) {
	validate := validator.New()
	for _, credential := range regPayload.Credentials {
		if err := validate.Struct(credential); err != nil {
			return "", fmt.Errorf("validation failed on credentials: %w", err)
		}
	}

	if err := validate.Struct(regPayload); err != nil {
		return "", fmt.Errorf("validation failed: %w", err)
	}
	k := Keycloak{}
	if err := k.ensureValidToken(); err != nil {
		return "", err
	}
	registerReq := RegisterTouristPayload{
		Username:      regPayload.Username,
		FirstName:     regPayload.FirstName,
		LastName:      regPayload.LastName,
		Email:         regPayload.Email,
		Enabled:       true,
		EmailVerified: true,
		Credentials: []Credentials{
			{
				Type:      "password",
				Value:     regPayload.Credentials[0].Value,
				Temporary: false,
			},
		},
	}
	log.Println("payload", regPayload)
	jsonData, err := json.Marshal(registerReq)
	if err != nil {
		logrus.WithError(err).Error("Error marshaling JSON")
		return "", err
	}
	req, err := http.NewRequest("POST", "http://localhost:8080/admin/realms/TourWithUs/users", bytes.NewBuffer(jsonData))
	if err != nil {
		logrus.WithError(err).Error("Error creating HTTP request")
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+k.BearerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.WithError(err).Error("Error performing HTTP request")
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.WithError(err).Error("Error reading response body")
		return "", err
	}
	logrus.Errorf("OK HTTP status: %d, response: %s", resp.StatusCode, body)

	log.Println("response value{}", resp)

	if resp.StatusCode != http.StatusCreated {
		logrus.Errorf("Non-OK HTTP status: %d", resp.StatusCode)
		return "", errors.New("something went wrong while connecting to Keycloak")
	}
	return "User created successfully", nil
}

func LoginUser(credentials LoginCredentials) (string, error) {
	validate := validator.New()
	if err := validate.Struct(credentials); err != nil {
		return "", fmt.Errorf("validation failed on credentials: %w", err)
	}
	endpoint := "http://localhost:8080/realms/TourWithUs/protocol/openid-connect/token"
	payload := map[string]string{
		"client_id":     "tourModel",
		"grant_type":    "password",
		"username":      credentials.Username,
		"password":      credentials.Password,
		"client_secret": "SjrSFWLqOzRVa36FC5SI6sdBDfc7AjJk",
	}

	form := url.Values{}
	for key, value := range payload {
		form.Add(key, value)
	}

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create login request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("login request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("login failed: %s", body)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("failed to parse response: %v", err)
	}

	token, ok := result["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("access token not found in response")
	}
	log.Println("token", token)
	return token, nil
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
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		log.Println("non-200 response: ", resp.StatusCode)
		return "", errors.New("something went wrong while fetching your key from keycloak")
	}
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var keycloakPublicKey string
	return keycloakPublicKey, nil
}

func isTokenExpired(tokenString string) (bool, error) {
	const bufferMinutes = 10
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		logrus.WithError(err).Error("Failed to parse token")
		return false, fmt.Errorf("failed to parse token: %w", err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, errors.New("could not parse claims")
	}
	exp, ok := claims["exp"].(float64)
	if !ok {
		return false, errors.New("'exp' claim missing or invalid in token")
	}
	expirationTime := time.Unix(int64(exp), 0)
	bufferDuration := time.Duration(bufferMinutes) * time.Minute
	isExpiringSoon := time.Now().After(expirationTime.Add(-bufferDuration))
	return isExpiringSoon, nil
}

func (k *Keycloak) ensureValidToken() error {
	if k.BearerToken == "" {
		logrus.Info("BearerToken is empty; refreshing token.")
		return k.refreshToken()
	}
	isExpired, err := isTokenExpired(k.BearerToken)
	if err != nil {
		logrus.WithError(err).Error("Failed to verify token expiration")
		return err
	}
	if isExpired {
		logrus.Info("BearerToken is expired; refreshing token.")
		return k.refreshToken()
	}
	logrus.Info("BearerToken is valid.")
	return nil
}

func (k *Keycloak) refreshToken() error {
	payload := Payload{
		ClientId:     "tourModel",
		ClientSecret: "SjrSFWLqOzRVa36FC5SI6sdBDfc7AjJk",
		GrantType:    "client_credentials",
		Username:     "admin",
		Password:     "admin",
	}
	_, err := k.GenerateToken(payload)
	if err != nil {
		logrus.WithError(err).Error("Error refreshing token")
		return fmt.Errorf("exception refreshing token: %w", err)
	}
	logrus.Info("BearerToken refreshed successfully.")
	return nil
}
