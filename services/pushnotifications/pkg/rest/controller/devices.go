package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/Casblanca-backend/services/pushnotifications/pkg/database"
	devices "github.com/radekkrejcirik01/Casblanca-backend/services/pushnotifications/pkg/model"
)

func SaveDevice(c *fiber.Ctx) error {
	t := &devices.Device{}

	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	if err := devices.SaveDevice(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "succes",
		Message: "Device succesfully saved",
	})
}
