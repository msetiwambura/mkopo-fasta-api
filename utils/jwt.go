package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	ChannelID string `json:"ChannelID"`
	IPAddress string `json:"IPAddress"`
	jwt.StandardClaims
}

func GenerateJWT(ChannelID, IPAddress string) (string, error) {

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		ChannelID: ChannelID,
		IPAddress: IPAddress,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil // Replace with your actual secret key
	})
	if err != nil {
		return nil, err // Return parsing error
	}

	if !token.Valid {
		return nil, errors.New("token is not valid") // Return specific error for invalid token
	}

	// At this point, token is valid, return claims
	return claims, nil
}
