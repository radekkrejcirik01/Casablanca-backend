package messages

import (
	"gorm.io/gorm"
)

type User struct {
	Email string
}

type Message struct {
	Id       uint `gorm:"primary_key;auto_increment;not_null"`
	Sender   string
	Receiver string
	Message  string
	Time     string
	IsRead   uint `gorm:"default:0"`
}

func (Message) TableName() string {
	return "messages"
}

type ConversationList struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
	Time     string `json:"time"`
	IsRead   uint   `json:"isRead"`
}

// GetMessages get messages
func GetMessages(db *gorm.DB, t *User, page string) ([]ConversationList, error) {
	offset := getOffset(page)

	messagedUsersQuery := `SELECT
								sender,
								receiver,
								message,
								time,
								is_read
							FROM
								messages
							WHERE
								id IN(
									SELECT
										MAX(id)
										FROM messages
									WHERE
										sender = '` + t.Email + `'
										OR receiver = '` + t.Email + `'
									GROUP BY
										( IF(sender = '` + t.Email + `', receiver, sender)))
							ORDER BY
								id DESC
							LIMIT 10 OFFSET ` + offset

	messagedUsers, err := GetConversationListFromQuery(db, messagedUsersQuery)
	if err != nil {
		return nil, err
	}

	var result []ConversationList
	for _, value := range messagedUsers {
		if value.Sender == t.Email {
			result = append(result, ConversationList{
				Sender:   value.Receiver,
				Receiver: value.Sender,
				Message:  value.Message,
				IsRead:   0, // Last message sent by user
			})
		} else {
			result = append(result, value)
		}
	}

	return result, nil
}

func GetConversationListFromQuery(db *gorm.DB, query string) ([]ConversationList, error) {
	rows, err := db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var array []ConversationList
	for rows.Next() {
		db.ScanRows(rows, &array)
	}

	return array, nil
}
