package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/Casblanca-backend/services/swipe/pkg/database"
	users "github.com/radekkrejcirik01/Casblanca-backend/services/swipe/pkg/model"
)

// GetUsers GET /get
func GetUsers(c *fiber.Ctx) error {
	t := &users.User{}

	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	users, err := users.GetUsers(database.DB, t)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(UserResponse{
		Status:  "succes",
		Message: "Users succesfully get",
		Data:    users,
	})
}
