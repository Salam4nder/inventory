package auth

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// NewJWT creates a new JWT token and signs it with the given secret.
// The signed token is returned as a string.
func NewJWT(secret string) (string, error) {
	expiration := time.Now().Add(1 * time.Hour)
	claims := jwt.StandardClaims{
		ExpiresAt: expiration.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateJWT validates the given JWT token.
// Returns an error if the token is invalid or expired.
func ValidateJWT(tokenString, secret string) error {
	token, err := jwt.Parse(
		tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf(
					"unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(secret), nil
		})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Printf("claims: %v", claims)
	} else {
		return err
	}

	return nil
}
