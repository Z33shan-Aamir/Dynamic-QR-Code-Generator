package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	qrcode := app.Group("/q")
	qrcode.Get("/id/:uuid?", func(c *fiber.Ctx) error {
		if c.Params("uuid") != "" {
			return c.SendString("UUID of qr code is  " + c.Params("uuid"))
		}
		return c.SendString("Look bro you ned to give me the uuids") // this will only return when there is a trailing URL direcotory like /id/:uuid? if there was only uuid I wouldn't work

	})
}
