package handlers

import (
	"DynamicQRBackend/dbconn"
	"DynamicQRBackend/models"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
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
		UserID:   userID,
		UserType: "user",
		Plan:     "free",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretkey)
}

// gets the userID via the http cookie
// (which was included in the request) contains the token
func CurrentUserID(c *fiber.Ctx) (string, error) {
	token := c.Locals("usertoken").(*jwt.Token)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("could not parse token")
	}
	return claims["user_id"].(string), nil
}

// gets the cookie and then returns user id or error if invalid
func GetUserIDFromToken(cookie string) (string, error) {
	token, err := jwt.Parse(cookie, func(t *jwt.Token) (any, error) {
		return secretkey, nil
	})

	if err != nil {
		return "", errors.New("could not parse the token")
	}
	if !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}

	user_id, ok := claims["user_id"]
	if !ok {
		return "", errors.New("could not determine user id from token")
	}

	var user models.User
	if err := dbconn.DB.First(&user, "id = ?", user_id).Error; err != nil {
		return "", err
	}

	return user.ID, nil
}
