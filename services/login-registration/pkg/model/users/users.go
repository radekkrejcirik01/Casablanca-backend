package users

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type UserRegistration struct {
	Firstname string         `json:""`
	Birthday  string         `json:""`
	Photos    pq.StringArray `gorm:"type:text[]"`
	Tags      pq.StringArray `gorm:"type:text[]"`
	Gender    string         `json:""`
	ShowMe    string         `json:""`
	Email     string         `json:""`
	Password  string         `json:""`
}

type User struct {
	Firstname string `json:"firstname"`
	Birthday  string `json:"birthday"`
	Gender    string `json:"gender"`
	ShowMe    string `json:"showMe"`
	Email     string `json:"email"`
	Password  string `json:"password"`
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
