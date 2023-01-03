package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/Casblanca-backend/services/messages/pkg/database"
	"github.com/radekkrejcirik01/Casblanca-backend/services/messages/pkg/model/matches"
)

// GetMatches POST /get/matches
func GetMatches(c *fiber.Ctx) error {
	page := c.Params("page")
	t := &matches.Email{}

	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	users, err := matches.GetMatches(database.DB, t, page)
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

// UpdateSeen POST /update/seen
func UpdateSeen(c *fiber.Ctx) error {
	t := &matches.Match{}

	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	if err := matches.UpdateSeen(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "succes",
		Message: "Seen succesfully updated",
	})
}
