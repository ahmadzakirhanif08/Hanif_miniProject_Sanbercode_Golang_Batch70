package utils
package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// SecretKey
var SecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// GenerateToken
func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token valid for 24 hours
		"iss":      "quiz-golang-app",                    
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign the token with the secret key
	t, err := token.SignedString(SecretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return t, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		// Pastikan algoritma yang digunakan sama
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token is not valid")
	}

	return token, nil
}