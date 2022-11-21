package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/Casblanca-backend/services/swipe/pkg/database"
	users "github.com/radekkrejcirik01/Casblanca-backend/services/swipe/pkg/model"
)

// GetUsers GET /get
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
	return c.Status(fiber.StatusOK).JSON(UserResponse{
		Status:  "succes",
		Message: "User succesfully get",
		Data: UsersResponse{
			Id:        t.Id,
			Email:     t.Email,
			Firstname: t.Firstname,
			Birthday:  t.Birthday,
			About:     t.About,
			Photos:    t.Photos,
			Tags:      t.Tags,
			Gender:    t.Gender,
			Distance:  t.Distance,
		},
	})
}
