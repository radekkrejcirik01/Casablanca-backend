package users

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	ShowMe       int
	Email        string
	Distance     int
	FilterByTags int
	Tags         pq.StringArray
}

type Users struct {
	Id        uint
	Email     string
	Firstname string
	Birthday  string
	About     string
	Photos    pq.StringArray
	Tags      pq.StringArray
	Gender    int
	Distance  int
}

// Get users from DB for swiper
func GetUsers(db *gorm.DB, t *Users) error {
	return db.Where("email = ?", t.Email).First(t).Error
}
