package main

import (
	"DynamicQRBackend/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "dynamic-qr-code-backend",
	})
	// dbconn.DBconn()
	routes.SetupRoutes(app)
	app.Listen(":3080")
}
