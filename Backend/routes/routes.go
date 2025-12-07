package routes

import (
	"DynamicQRBackend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api/v1")

	qrcode := api.Group("/q")
	qrcode.Get("/id/:nanoid", handlers.GetRedirectUrl)
	qrcode.Post("/create", handlers.CreateQRCode)
	qrcode.Post("/update", handlers.UpdateQRCode)

	// qrcode.Post("/Create", handlers.AddQRCode)

	auth := api.Group("/auth")
	auth.Get("/get-user", handlers.User)
	auth.Post("/signin", handlers.Signup)
	auth.Post("/login", handlers.Login)
}
