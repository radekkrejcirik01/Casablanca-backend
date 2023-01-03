package users

import (
	"log"

	"github.com/appleboy/go-fcm"
	"github.com/radekkrejcirik01/Casblanca-backend/services/swipe/pkg/database"
	"gorm.io/gorm"
)

type Like struct {
	Id     uint `gorm:"primary_key;auto_increment;not_null"`
	Email  string
	User   string
	Value  int `gorm:"size:1"`
	IsSeen int `gorm:"size:1;default:0"`
}

func (Like) TableName() string {
	return "likes"
}

type Notification struct {
	Title   string
	Body    string
	Devices []string
}

// LikeUser like user
func LikeUser(db *gorm.DB, t *Like) error {
	var err error
	if db.Model(&t).Where("email = ? AND user = ?", t.Email, t.User).Update("value", t.Value).RowsAffected == 0 {
		err = db.FirstOrCreate(t, t).Error
	}
	if err == nil {
		if CheckIfMatch(db, t) {
			NotifyAboutMatch(db, t)
		}
	}
	return err
}

func CheckIfMatch(db *gorm.DB, t *Like) bool {
	var exists bool
	db.Table("likes").Select("count(*) > 1").
		Where("(email = ? AND user = ? AND value = 1) OR (email = ? AND user = ? AND value = 1)",
			t.Email, t.User, t.User, t.Email).
		Find(&exists)

	return exists
}

func NotifyAboutMatch(db *gorm.DB, t *Like) error {
	tokens := &[]string{}
	users := []string{t.Email, t.User}
	if err := GetUserTokensByUsers(db, tokens, users); err != nil {
		return err
	}
	notification := Notification{
		Title:   "You got a new match! 🥳",
		Body:    "Let's see!",
		Devices: *tokens,
	}

	return SendNotification(&notification)
}

func GetUserTokensByUsers(db *gorm.DB, t *[]string, users []string) error {
	return db.Table("devices").Select("device_token").
		Where("email = ? OR email = ?", users[0], users[1]).
		Find(t).Error
}

func SendNotification(t *Notification) error {
	fcmClient := database.GetFcmClient()
	tokens := t.Devices

	for _, token := range tokens {
		msg := &fcm.Message{
			To: token,
			Data: map[string]interface{}{
				"type": "match",
			},
			Notification: &fcm.Notification{
				Title: t.Title,
				Body:  t.Body,
				Sound: "default",
			},
		}

		client, err := fcm.NewClient(fcmClient)
		if err != nil {
			log.Fatalln(err)
			return err
		}

		response, err := client.Send(msg)
		if err != nil {
			log.Fatalln(err)
			return err
		}

		log.Printf("%#v\n", response)
	}

	return nil
}
