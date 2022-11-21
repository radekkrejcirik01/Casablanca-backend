package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/Casblanca-backend/services/swipe/pkg/database"
)

// UpdatePhotos PUT /photos/update
func GetUsers(c *fiber.Ctx) error {
	t := &users.Users{}

	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	if err := users.GetUsers(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "succes",
		Message: "Users succesfully get!",
	})
}
