package users

import (
	"gorm.io/gorm"
)

type User struct {
	Id        uint   `gorm:"primary_key;auto_increment;not_null"`
	Firstname string `json:""`
	Birthday  string `json:""`
	Gender    string `json:""`
	ShowMe    string `json:""`
	Email     string `json:""`
	Password  string `json:""`
}

func (User) TableName() string {
	return "users"
}

// Create new User in DB
func CreateUser(db *gorm.DB, t *User) error {
	return db.Create(t).Error
}

// Read one User from DB by ID
func Read(db *gorm.DB, t *User, id string) error {
	return db.Where("id = ?", id).First(t).Error
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
func DeleteById(db *gorm.DB, id string) error {
	users := &User{}
	if err := Read(db, users, id); err != nil {
		return err
	}
	return db.Where("id = ?", id).Delete(users).Error

}

// Login authenticate user
func Login(db *gorm.DB, t *User) error {
	return db.Where("username = ? AND password = ?", t.Email, t.Password).First(t).Error
}
