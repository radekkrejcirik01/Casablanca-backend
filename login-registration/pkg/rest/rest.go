package rest

import (
	"login-registration/pkg/rest/controller"

	"github.com/gofiber/fiber/v2"
)

// Create new REST API serveer
func Create() *fiber.App {
	app := fiber.New()

	app.Get("/", controller.Index)

	app.Post("/register", controller.UserRegister)
	app.Post("/login", controller.UserLogin)
	app.Get("/:id", controller.UserGet)
	app.Put("/update", controller.UserPut)
	app.Delete("/delete/:id", controller.UserDelete)

	return app
}
