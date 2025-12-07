package handlers

import (
	"DynamicQRBackend/dbconn"
	"DynamicQRBackend/models"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetRedirectUrl(c *fiber.Ctx) error {
	if c.Params("nanoid") == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No QR Code of id" + c.Params("nanoid") + "was found",
		})
	}
	qrcodeid := c.Params("nanoid")
	var qrcode models.QRCode
	if err := dbconn.DB.Where("id = ?", qrcodeid).First(&qrcode).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Redirect(qrcode.DestinationURL)

}

func CreateQRCode(c *fiber.Ctx) error {
	var reqQRCode models.CreateQRCodeDTO
	if err := c.BodyParser(&reqQRCode); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request JSON",
		})
	}
	if err := validate.Struct(reqQRCode); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "missing data in request JSON",
		})
	}

	cookie := c.Cookies("jwt")

	// also checks if it exists in DB
	user_id, err := GetUserIDFromToken(cookie)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if _, err := url.Parse(reqQRCode.DestinationURL); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid destination url",
		})
	}
	qrcode := models.QRCode{
		ID:             reqQRCode.ID,
		UserID:         user_id,
		Name:           reqQRCode.Name,
		Description:    reqQRCode.Description,
		DestinationURL: reqQRCode.DestinationURL,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	if err := dbconn.DB.Create(&qrcode).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"msg": "qrcode was created",
	})

}

func UpdateQRCode(c *fiber.Ctx) error {

	var updatedQRCode models.UpdateQRCodeDTO
	if err := c.BodyParser(&updatedQRCode); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request JSON",
		})
	}
	if err := validate.Struct(updatedQRCode); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "missing data in request JSON",
		})
	}

	cookie := c.Cookies("jwt")

	// also checks if it exists in DB
	user_id, err := GetUserIDFromToken(cookie)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var qrCode models.QRCode
	if err := dbconn.DB.Where("id = ? AND user_id = ?", updatedQRCode.ID, user_id).First(&qrCode).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	qrCode.Name = updatedQRCode.Name
	qrCode.Description = updatedQRCode.Description
	qrCode.DestinationURL = updatedQRCode.DestinationURL
	qrCode.UpdatedAt = time.Now()
	qrCode.Active = updatedQRCode.Active

	if err := dbconn.DB.Omit("id", "user_id").Save(&qrCode).Error; err == nil {
		return c.JSON(fiber.Map{
			"msg": "qrcode was updated succesfully",
		})
	} else {
		return c.Status(500).JSON(fiber.Map{
			"error": "Could not save qrcode: " + err.Error(),
		})
	}

}

func DeleteQRCode(c *fiber.Ctx) error {
	if c.Params("nanoid") == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No qr code id was given",
		})
	}
	var qrCodeID string = c.Params("nanoid")

	cookie := c.Cookies("jwt")

	user_id, err := GetUserIDFromToken(cookie)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var QRCode models.QRCode
	results := dbconn.DB.Where("id = ? AND user_id = ?", qrCodeID, user_id).Delete(&QRCode)

	if results.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Could Not delete QRCode: " + results.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"msg": "QRCode Succefully deleted with ID: " + qrCodeID,
	})

}
