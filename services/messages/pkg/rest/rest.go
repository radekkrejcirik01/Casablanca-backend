package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/Casblanca-backend/services/messages/pkg/rest/controller"
)

// Create new REST API serveer
func Create() *fiber.App {
	app := fiber.New()

	app.Get("/", controller.Index)

	app.Post("/get/matches/:page", controller.GetMatches)
	app.Post("/update/seen", controller.UpdateSeen)

	app.Post("/get/conversations/:page", controller.GetConversations)
	app.Post("/get/messages/:page", controller.GetMessages)
	app.Post("/update/read", controller.UpdateRead)
	app.Post("/send/message", controller.SendMessage)
	app.Post("/get/user", controller.GetUser)

	return app
}
