package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/Casblanca-backend/services/messages/pkg/database"
	"github.com/radekkrejcirik01/Casblanca-backend/services/messages/pkg/model/matches"
)

// GetMatches POST /get/matches
func GetMatches(c *fiber.Ctx) error {
	t := &matches.User{}

	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	users, err := matches.GetMatches(database.DB, t)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(ResponseMatches{
		Status:  "succes",
		Message: "Matches succesfully get",
		Data:    users,
	})
}
