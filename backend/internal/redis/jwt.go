package redis

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
	"log"
	)

	var JwtKey = []byte("my_secret_key")

func GenerateToken(userID, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Minute * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}

	log.Printf("Generated Token: %s\n", tokenString)  // Debugging: Print the generated token
	return tokenString, err
}
