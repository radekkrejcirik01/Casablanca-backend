package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/Casblanca-backend/services/login-registration/pkg/database"
	"github.com/radekkrejcirik01/Casblanca-backend/services/login-registration/pkg/model/users"
)

// UserRegister POST /register
func UserRegister(c *fiber.Ctx) error {
	t := &users.Registration{}
	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	user := &users.User{
		Id:        t.Id,
		Firstname: t.Firstname,
		Birthday:  t.Birthday,
		Gender:    t.Gender,
		ShowMe:    t.ShowMe,
		Email:     t.Email,
		Password:  t.Password,
	}
	tags := &users.Tags{Id: t.Id, User: t.Email, Tags: t.Tags}
	photos := &users.Photos{Id: t.Id, User: t.Email, Photos: t.Photos}

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

	return c.Status(fiber.StatusOK).JSON(RegistrationResponse{
		Status:  "succes",
		Message: "User succesfully registered!",
		Data:    &users.RegistrationDataResponse{Id: user.Id, Email: user.Email},
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
	if err := users.Login(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "succes",
		Message: "User succesfully authenticated!",
	})
}

// UserGet GET /:id
func UserGet(c *fiber.Ctx) error {
	id := c.Params("id")
	t := &users.User{}
	if err := users.ReadById(database.DB, t, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(UserGetResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(UserGetResponse{
		Status: "succes",
		Data:   &[]users.User{*t},
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
	id := c.Params("id")
	if err := users.DeleteById(database.DB, id); err != nil {
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
