package messages

import "gorm.io/gorm"

type User struct {
	Email string
}

type Message struct {
	Id      uint `gorm:"primary_key;auto_increment;not_null"`
	Email   string
	User    string
	Message string
}

func (Message) TableName() string {
	return "messages"
}

// GetMessages get messages
func GetMessages(db *gorm.DB, t *User, offset string) error {
	return nil
}
