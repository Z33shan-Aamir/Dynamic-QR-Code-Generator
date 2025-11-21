package handlers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// contains helper functions to get the IP of the client and other
// client data as well as some location data
func GeneratePasswordHash(p string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), 16)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CompareHash(hashed_password, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password))

	return err == nil // returns true if there is no error

}

func GenerateToken(userID string) (string, error) {
	claims := MyClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretkey)
}
