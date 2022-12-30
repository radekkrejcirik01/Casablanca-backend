package messages

import (
	"github.com/radekkrejcirik01/Casblanca-backend/services/messages/pkg/model/helpers"
	"gorm.io/gorm"
)

type MessagesBody struct {
	Email string
	User  string
}

type Messages struct {
	Id       string `json:"id"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
	Time     string `json:"time"`
}

// GetMessages get messages
func GetMessages(db *gorm.DB, t *MessagesBody, page string) ([]Messages, error) {
	offset := helpers.GetOffset(page)

	messagesQuery := `SELECT
							id,
							sender,
							receiver,
							message,
							time
						FROM
							messages
						WHERE (sender = '` + t.Email + `'
							AND receiver = '` + t.User + `')
							OR(sender = '` + t.User + `'
								AND receiver = '` + t.Email + `')
						ORDER BY
							id DESC
						LIMIT 10 OFFSET ` + offset

	messages, err := GetMessagesFromQuery(db, messagesQuery)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func GetMessagesFromQuery(db *gorm.DB, query string) ([]Messages, error) {
	rows, err := db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var array []Messages
	for rows.Next() {
		db.ScanRows(rows, &array)
	}

	return array, nil
}
