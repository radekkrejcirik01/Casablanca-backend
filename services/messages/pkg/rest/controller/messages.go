package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/Casblanca-backend/services/messages/pkg/database"
	messages "github.com/radekkrejcirik01/Casblanca-backend/services/messages/pkg/model/messages"
)

// GetConversations POST /get/conversations/:page
func GetConversations(c *fiber.Ctx) error {
	page := c.Params("page")

	t := &messages.Email{}

	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	conversationList, err := messages.GetConversationsList(database.DB, t, page)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(ResponseConversationList{
		Status:  "succes",
		Message: "Conversation list succesfully get",
		Data:    conversationList,
	})
}

// GetMessages POST /get/messages/:page
func GetMessages(c *fiber.Ctx) error {
	page := c.Params("page")

	t := &messages.MessagesBody{}

	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	messages, err := messages.GetMessages(database.DB, t, page)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(ResponseMessages{
		Status:  "succes",
		Message: "Messages succesfully get",
		Data:    messages,
	})
}

// UpdateRead POST /update/read
func UpdateRead(c *fiber.Ctx) error {
	t := &messages.MessagesBody{}

	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	if err := messages.UpdateRead(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "succes",
		Message: "Read succesfully updated",
	})
}

// SendMessage POST /send/messages
func SendMessage(c *fiber.Ctx) error {
	t := &messages.SentMessage{}

	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	if err := messages.SendMessage(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "succes",
		Message: "Message succesfully sent",
	})
}

// GetUser POST /get/user
func GetUser(c *fiber.Ctx) error {
	t := &messages.UserBody{}

	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	user, err := messages.GetUser(database.DB, t)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(ResponseUser{
		Status:  "succes",
		Message: "User succesfully get",
		Data:    user,
	})
}
