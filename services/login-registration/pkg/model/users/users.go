package users

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type UserRegistration struct {
	Id            uint `gorm:"primary_key;auto_increment;not_null"`
	Firstname     string
	Birthday      string
	About         string
	Photos        pq.StringArray
	Tags          pq.StringArray
	Gender        int
	ShowMe        int
	Email         string
	Distance      int
	FilterByTags  int
	Notifications int
	Password      string
}

type User struct {
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

func (User) TableName() string {
	return "users"
}

// Create new User in DB
func CreateUser(db *gorm.DB, t *User) error {
	return db.Create(t).Error
}

// Login authenticate user
func LoginUser(db *gorm.DB, t *User) error {
	return db.Where("email = ? AND password = ?", t.Email, t.Password).First(t).Error
}

// Read one User from DB by ID
func ReadByEmail(db *gorm.DB, t *User) error {
	return db.Where("email = ?", t.Email).First(t).Error
}

// ReadAll User from DB
func ReadAll(db *gorm.DB, t *[]User) error {
	return db.Find(t).Error
}

// Update User in DB
func Update(db *gorm.DB, t *User) error {
	return db.Save(t).Error
}

// Delete User from DB
func Delete(db *gorm.DB, t *User) error {
	return db.Delete(t).Error
}

// DeleteByID one User by ID
func DeleteById(db *gorm.DB, t *User) error {
	users := &User{}
	if err := ReadByEmail(db, t); err != nil {
		return err
	}
	return db.Where("email = ?", t.Email).Delete(users).Error

}
