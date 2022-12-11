package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/Casblanca-backend/services/swipe/pkg/database"
	users "github.com/radekkrejcirik01/Casblanca-backend/services/swipe/pkg/model/users"
)

// GetUsers POST /get
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

// LikeUser POST /like
func LikeUser(c *fiber.Ctx) error {
	t := &users.Like{}

	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	if err := users.LikeUser(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(LikeResponse{
		Status:  "succes",
		Message: "Like succesfully perfomed",
		Value:   strconv.Itoa(t.Value),
	})
}
