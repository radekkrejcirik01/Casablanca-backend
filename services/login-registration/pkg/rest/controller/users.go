package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/Casblanca-backend/services/login-registration/pkg/database"
	"github.com/radekkrejcirik01/Casblanca-backend/services/login-registration/pkg/model/users"
)

// UserRegister POST /register
func UserRegister(c *fiber.Ctx) error {
	t := &users.UserRegistration{}

	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	user := &users.User{
		Id:            t.Id,
		Firstname:     t.Firstname,
		Birthday:      t.Birthday,
		About:         t.About,
		Gender:        t.Gender,
		ShowMe:        t.ShowMe,
		Email:         t.Email,
		Distance:      t.Distance,
		FilterByTags:  t.FilterByTags,
		Notifications: t.Notifications,
		Password:      t.Password,
	}
	tags := &users.Tags{User: t.Email, Tags: t.Tags}
	photos := &users.Photos{User: t.Email, Photos: t.Photos}

	if err := users.CreateUser(database.DB, user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	if err := users.AddPhotos(database.DB, photos); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	if err := users.AddTags(database.DB, tags); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	if err := users.ReadByEmail(database.DB, user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "succes",
		Message: "User succesfully registered!",
	})
}

// UserLogin AUTHENTICATE /login
func UserLogin(c *fiber.Ctx) error {
	t := &users.User{}

	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	if err := users.LoginUser(database.DB, t); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return GetUserResponse(c, t)
}

func UserGet(c *fiber.Ctx) error {
	t := &users.User{}

	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	if err := users.ReadByEmail(database.DB, t); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return GetUserResponse(c, t)
}

func GetUserResponse(c *fiber.Ctx, t *users.User) error {
	photos := &[]string{}
	tags := &[]string{}

	if err := users.GetPhotos(database.DB, photos, t.Email); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	if err := users.GetTags(database.DB, tags, t.Email); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(UserResponse{
		Status:  "succes",
		Message: "User succesfully authenticated!",
		Data: UserDataResponse{
			Email:         t.Email,
			Firstname:     t.Firstname,
			Birthday:      t.Birthday,
			About:         t.About,
			Photos:        *photos,
			Tags:          *tags,
			Gender:        t.Gender,
			ShowMe:        t.ShowMe,
			Distance:      t.Distance,
			FilterByTags:  t.FilterByTags,
			Notifications: t.Notifications,
		},
	})
}

// UpdatePhotos PUT /photos/update
func UpdatePhotos(c *fiber.Ctx) error {
	t := &users.Photos{}

	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	if err := users.UpdatePhotos(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "succes",
		Message: "Photos succesfully updated!",
	})
}

// UpdateTags PUT /tags/update
func UpdateTags(c *fiber.Ctx) error {
	t := &users.Tags{}
	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	if err := users.UpdateTags(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "succes",
		Message: "Tags succesfully updated!",
	})
}

// UpdateAbout PUT /about/update
func UpdateAbout(c *fiber.Ctx) error {
	t := &users.About{}
	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	if err := users.UpdateAbout(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "succes",
		Message: "About succesfully updated!",
	})
}

// UpdateNotifications PUT /notifications/update
func UpdateNotifications(c *fiber.Ctx) error {
	t := &users.Notifications{}
	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	if err := users.UpdateNotifications(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "succes",
		Message: "Notifications succesfully updated!",
	})
}

// UpdateDistance PUT /distance/update
func UpdateDistance(c *fiber.Ctx) error {
	t := &users.Distance{}
	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	if err := users.UpdateDistance(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "succes",
		Message: "Distance succesfully updated!",
	})
}

// UpdateFilterByTags PUT /filterByTags/update
func UpdateFilterByTags(c *fiber.Ctx) error {
	t := &users.FilterByTags{}
	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	if err := users.UpdateFilterByTags(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "succes",
		Message: "FilterByTags succesfully updated!",
	})
}

// UpdateShowMe PUT /showMe/update
func UpdateShowMe(c *fiber.Ctx) error {
	t := &users.ShowMe{}
	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	if err := users.UpdateShowMe(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "succes",
		Message: "ShowMe succesfully updated!",
	})
}

// UserDel DELETE /delete/:id
func UserDelete(c *fiber.Ctx) error {
	t := &users.User{}
	if err := users.DeleteById(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "succes",
		Message: "User succesfully deleted!",
	})
}
