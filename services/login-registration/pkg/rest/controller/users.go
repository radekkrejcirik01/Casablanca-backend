package controller

import (
	"login-registration/pkg/database"
	"login-registration/pkg/model/users"

	"github.com/gofiber/fiber/v2"
)

// UserRegister POST /register
func UserRegister(c *fiber.Ctx) error {
	t := &users.USERS{}
	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	if err := users.Create(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(resp{
		Status:  "succes",
		Message: "User succesfully registered!",
	})
}

// UserLogin AUTHENTICATE /login
func UserLogin(c *fiber.Ctx) error {
	t := &users.USERS{}
	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	if err := users.Login(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(resp{
		Status:  "succes",
		Message: "User succesfully authenticated!",
	})
}

// UserGet GET /:id
func UserGet(c *fiber.Ctx) error {
	id := c.Params("id")
	t := &users.USERS{}
	if err := users.Read(database.DB, t, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(resp{
		Status: "succes",
		Data:   &[]users.USERS{*t},
	})
}

// UserPut PUT /update
func UserPut(c *fiber.Ctx) error {
	t := &users.USERS{}
	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	if err := users.Update(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(resp{
		Status:  "succes",
		Message: "User succesfully updated!",
	})
}

// UserDel DELETE /delete/:id
func UserDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := users.DeleteById(database.DB, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(resp{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(resp{
		Status:  "succes",
		Message: "User succesfully deleted!",
	})
}
