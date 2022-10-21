package users

import (
	"gorm.io/gorm"
)

// USERS db model
type USERS struct {
	Id       uint   `gorm:"primary_key;auto_increment;not_null"`
	Username string `json:""`
	Password string `json:""`
}

func (USERS) TableName() string {
	return "users"
}

// Create new USERS in DB
func Create(db *gorm.DB, t *USERS) error {
	return db.Create(t).Error
}

// Read one USERS from DB by ID
func Read(db *gorm.DB, t *USERS, id string) error {
	return db.Where("id = ?", id).First(t).Error
}

// ReadAll USERS from DB
func ReadAll(db *gorm.DB, t *[]USERS) error {
	return db.Find(t).Error
}

// Update USERS in DB
func Update(db *gorm.DB, t *USERS) error {
	return db.Save(t).Error
}

// Delete USERS from DB
func Delete(db *gorm.DB, t *USERS) error {
	return db.Delete(t).Error
}

// DeleteByID one USERS by ID
func DeleteById(db *gorm.DB, id string) error {
	users := &USERS{}
	if err := Read(db, users, id); err != nil {
		return err
	}
	return db.Where("id = ?", id).Delete(users).Error

}

// Login authenticate user
func Login(db *gorm.DB, t *USERS) error {
	return db.Where("username = ? AND password = ?", t.Username, t.Password).First(t).Error
}
