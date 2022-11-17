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

type About struct {
	Email string
	About string
}

type Notifications struct {
	Email         string
	Notifications int
}

type Distance struct {
	Email    string
	Distance int
}

type FilterByTags struct {
	Email        string
	FilterByTags int
}

type ShowMe struct {
	Email  string
	ShowMe int
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

// Read one user from DB by email
func ReadByEmail(db *gorm.DB, t *User) error {
	return db.Where("email = ?", t.Email).First(t).Error
}

// Update about in users table in DB
func UpdateAbout(db *gorm.DB, t *About) error {
	return db.Table("users").Where("email = ?", t.Email).Update("about", t.About).Error
}

// Update notifications in users table in DB
func UpdateNotifications(db *gorm.DB, t *Notifications) error {
	return db.Table("users").Where("email = ?", t.Email).Update("notifications", t.Notifications).Error
}

// Update distance in users table in DB
func UpdateDistance(db *gorm.DB, t *Distance) error {
	return db.Table("users").Where("email = ?", t.Email).Update("distance", t.Distance).Error
}

// Update filterByTags in users table in DB
func UpdateFilterByTags(db *gorm.DB, t *FilterByTags) error {
	return db.Table("users").Where("email = ?", t.Email).Update("filter_by_tags", t.FilterByTags).Error
}

// Update showMe in users table in DB
func UpdateShowMe(db *gorm.DB, t *ShowMe) error {
	return db.Table("users").Where("email = ?", t.Email).Update("show_me", t.ShowMe).Error
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
