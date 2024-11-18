package OperatorSecurity

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type OperatorClaims struct {
	jwt.RegisteredClaims
	Email string
	ID    int64
}

var secretKey = "404E635266556A586E3272357538782F413F4428472B4B6250645367566B5970"

func Generate(email string, id int64) (string, string, error) {
	operatorClaims := OperatorClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Tour Operator Admin",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(24 * time.Hour)},
		},
		Email: email,
		ID:    id,
	}
	refOperatorCliams := jwt.RegisteredClaims{
		Issuer:    "Tour Operator Admin",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(48 * time.Hour)},
	}

	// Generate JWT tokens
	walletToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, operatorClaims).SignedString([]byte(secretKey))
	if err != nil {
		return "", "", err
	}
	refWalletToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refOperatorCliams).SignedString([]byte(secretKey))
	if err != nil {
		return "", "", err
	}

	return walletToken, refWalletToken, nil
}

func Parse(tokenString string) (*OperatorClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &OperatorClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(*OperatorClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
