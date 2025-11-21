package handlers

import (
	"DynamicQRBackend/dbconn"
	"DynamicQRBackend/models"

	"github.com/gofiber/fiber/v2"
)

func GetRedirectUrl(c *fiber.Ctx) error {
	if c.Params("nanoid") == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No QR Code of id" + c.Params("nanoid") + "was found",
		})
	}
	// TODO
	// Also create A new analytic which has frelevant functions in another file
	// Like getting the location of the request. Client info

	return c.Redirect("https://github.com/supabase-community/supabase-go")
}

func AddQRCode() {

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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "QRCode Succefully deleted with ID: " + qrCodeID,
	})

}
