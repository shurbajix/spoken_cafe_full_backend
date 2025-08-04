package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey = []byte("your-very-secret-key") // ideally loead form config

type Claims struct {
	UserID string `json:"user_id`
	jwt.RegisteredClaims
}

// GnerateRoken creaters a signed JWT token with expiration

func GenerateToken(UserID string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

// ValidateToken pareses and validates JWT token string and returns claims
func ValidateToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.token) (interface{}, error) {
		return jwtSecretKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
