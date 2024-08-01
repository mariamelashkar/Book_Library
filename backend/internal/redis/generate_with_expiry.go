package redis

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
	"log"
	)

func GenerateTokenWithExpiry(userID, role string, minutes int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Duration(minutes) * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}

	log.Printf("Generated Token with Expiry: %s\n", tokenString)  // Debugging: Print the generated token
	return tokenString, err
}