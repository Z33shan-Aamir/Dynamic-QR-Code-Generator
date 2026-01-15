package routes

import (
	"DynamicQRBackend/handlers"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

var secretkey = []byte("sesyt7yZE5ImLLVCgDp3sefdSc4aYXQWhwgllMOd8Mo5xDep2U4U1EElHI1u+0c0gjZ+CG0w/COne+dHX5AArQ==")

func SetupRoutes(app *fiber.App) {

	jwtMiddleWare := jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: secretkey},
		TokenLookup: "cookie:jwt",
		ContextKey:  "usertoken",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized or invalid token",
			})
		},
	})

	api := app.Group("/api/v1")

	// ------ PUBLIC ROUTES ------
	api.Get("/id/:nanoid", handlers.GetRedirectUrl)
	auth := api.Group("/auth")
	// auth.Get("/get-user", handlers.User) // checks if user is authenticated and returns user data
	auth.Post("/signin", handlers.Signup)
	auth.Post("/login", handlers.Login)

	// --- PROTECTED ROUTES-------
	qrcode := api.Group("/q")
	qrcode.Post("/create", jwtMiddleWare, handlers.CreateQRCode)
	qrcode.Post("/update", jwtMiddleWare, handlers.UpdateQRCode)
	qrcode.Delete("/delete", jwtMiddleWare, handlers.DeleteQRCode)
}
