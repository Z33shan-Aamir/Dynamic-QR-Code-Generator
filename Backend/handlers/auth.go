package handlers

import (
	"DynamicQRBackend/dbconn"
	"DynamicQRBackend/models"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var secretkey = []byte("j4hb23jh4bjhb234hjb1khg3v4132123!@#!@#hbvm1cv23b!M@#Vnm")
var validate = validator.New()

type MyClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func Signup(c *fiber.Ctx) error {

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Could not parse request json",
		})
	}
	var existingUser models.User
	results := dbconn.DB.Where("email = ?", user.Email).First(&existingUser)

	if results.Error == nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "User with email already exists",
		})
	}

	if results.Error != nil && !errors.Is(results.Error, gorm.ErrRecordNotFound) {
		// executes when there is an "unexpected" out of the blue error
		return c.Status(500).JSON(fiber.Map{
			"error": "Database error",
		})
	}

	hashed, err := GeneratePasswordHash(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Could not hash password",
		})
	}

	// user.CreatedAt = time.Now()
	user.Password = string(hashed)
	user.ID = uuid.NewString()
	user.Role = "user"

	if err := validate.Struct(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Missing JSON data",
		})
	}

	dbconn.DB.Create(&user)

	return c.Status(201).JSON(fiber.Map{
		"message": "user created successfully",
		"id":      user.ID,
	})
}

func Login(c *fiber.Ctx) error {

	type user_login struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	var body user_login

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid JSON",
		})
	}

	if err := validate.Struct(body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Missing JSON data",
		})
	}

	var user models.User
	err := dbconn.DB.Where("email = ?", body.Email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(400).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	if !CompareHash(user.Password, body.Password) {
		return c.Status(400).JSON(fiber.Map{
			"error": "incorrect password",
		})
	}

	token, err := GenerateToken(user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "could not create token",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	cookie.SameSite = "Lax"

	// creates a cookie with the above data
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "login successful",
		"token":   token,
	})

}

// checks if user is loged in

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.Parse(cookie, func(t *jwt.Token) (any, error) {
		return secretkey, nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{
			"error": "invalid or expired token",
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	userID, ok := claims["user_id"]
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	var user models.User
	if err := dbconn.DB.First(&user, "id = ?", userID).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "authenticated",
		"user":    user,
	})
}
