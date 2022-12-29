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
