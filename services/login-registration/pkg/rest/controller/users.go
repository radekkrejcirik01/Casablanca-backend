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
		Firstname: t.Firstname,
		Birthday:  t.Birthday,
		Gender:    t.Gender,
		ShowMe:    t.ShowMe,
		Email:     t.Email,
		Password:  t.Password,
	}
	tags := &users.Tags{User: t.Email, Tags: t.Tags}
	photos := &users.Photos{User: t.Email, Photos: t.Photos}

	if err := users.CreateUser(database.DB, user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	if err := users.CreatePhoto(database.DB, photos); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	if err := users.CreateTag(database.DB, tags); err != nil {
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

// AddTag POST /tags
func AddTag(c *fiber.Ctx) error {
	t := &users.Tags{}

	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	if err := users.CreateTag(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "succes",
		Message: "Tag succesfully added!",
	})
}

// AddPhoto POST /photos
func AddPhoto(c *fiber.Ctx) error {
	t := &users.Photos{}

	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	if err := users.CreatePhoto(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "succes",
		Message: "Tag succesfully added!",
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
			Email:     t.Email,
			Firstname: t.Firstname,
			Birthday:  t.Birthday,
			Photos:    *photos,
			Tags:      *tags,
			Gender:    t.Gender,
			ShowMe:    t.ShowMe,
		},
	})
}

// UserPut PUT /update
func UserPut(c *fiber.Ctx) error {
	t := &users.User{}
	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	if err := users.Update(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "succes",
		Message: "User succesfully updated!",
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
