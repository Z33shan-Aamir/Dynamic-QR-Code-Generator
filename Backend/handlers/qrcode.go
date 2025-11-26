package handlers

import (
	"DynamicQRBackend/dbconn"
	"DynamicQRBackend/models"
	"net/url"

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

func AddQRCode(c *fiber.Ctx) error {
	var qrcode models.QRCode
	if err := c.BodyParser(&qrcode); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request JSON",
		})
	}
	if err := validate.Struct(qrcode); err != nil {
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
	if _, err := url.Parse(qrcode.DestinationURL); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid destination url",
		})
	}
	qrcode.UserID = user_id

	if err := dbconn.DB.Create(&qrcode).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"msg": "qrcode was created",
	})

}

func UpdateQRCode() {

}

func DeleteQRCode(c *fiber.Ctx) error {
	if c.Params("nanoid") == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No qr code id was given",
		})
	}
	var qrCodeID string = c.Params("nanoid")

	var QRCode models.QRCode
	results := dbconn.DB.Where("id = ?", qrCodeID).Delete(&QRCode)

	if results.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Could Not delete QRCode",
		})
	}

	return c.JSON(fiber.Map{
		"msg": "QRCode Succefully deleted with ID: " + qrCodeID,
	})

}
