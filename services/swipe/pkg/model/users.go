package users

import (
	"gorm.io/gorm"
)

type Users struct {
	Id            uint
	Firstname     string
	Birthday      string
	About         string `gorm:"size:256"`
	Gender        int
	ShowMe        int
	Email         string
	Distance      int `gorm:"default:20"`
	FilterByTags  int `gorm:"default:0"`
	Notifications int `gorm:"default:1"`
	Password      string
}

// Get users from DB for swiper
func GetUsers(db *gorm.DB, t *Users) error {
	return db.Where("email = ?", t.Email).First(t).Error
}
