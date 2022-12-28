package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/Casblanca-backend/services/messages/pkg/database"
	messages "github.com/radekkrejcirik01/Casblanca-backend/services/messages/pkg/model/messages"
)

// GetMessages POST /get/messages/:offset
func GetMessages(c *fiber.Ctx) error {
	offset := c.Params("offset")

	t := &messages.User{}

	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	err := messages.GetMessages(database.DB, t, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "succes",
		Message: "Messages succesfully get",
	})
}
