package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/Casblanca-backend/services/login-registration/pkg/rest/controller"
)

// Create new REST API serveer
func Create() *fiber.App {
	app := fiber.New()

	app.Get("/", controller.Index)

	app.Post("/register", controller.UserRegister)
	app.Post("/login", controller.UserLogin)

	app.Put("/photos/update", controller.UpdatePhotos)
	app.Put("/tags/update", controller.UpdateTags)
	app.Put("/about/update", controller.UpdateAbout)
	app.Put("/notifications/update", controller.UpdateNotifications)
	app.Put("/distancePreference/update", controller.UpdateDistancePreference)
	app.Put("/agePreference/update", controller.UpdateAgePreference)
	app.Put("/filterByTags/update", controller.UpdateFilterByTags)
	app.Put("/showMe/update", controller.UpdateShowMe)
	app.Put("/lastActive/update", controller.UpdateLastActive)

	app.Post("/get", controller.UserGet)

	app.Delete("/delete", controller.UserDelete)

	return app
}
