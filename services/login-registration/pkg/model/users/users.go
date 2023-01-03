package users

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type UserRegistration struct {
	Id                 uint `gorm:"primary_key;auto_increment;not_null"`
	Email              string
	Firstname          string
	Birthday           string
	ProfilePicture     string
	About              string
	Photos             pq.StringArray
	Tags               pq.StringArray
	Gender             int
	ShowMe             int
	DistancePreference int
	AgePreference      string
	FilterByTags       int
	Notifications      int
	LastActive         string
	Password           string
}

type User struct {
	Id                 uint
	Email              string
	Firstname          string
	Birthday           string
	ProfilePicture     string
	About              string `gorm:"size:256"`
	Gender             int
	ShowMe             int
	DistancePreference int `gorm:"default:20"`
	AgePreference      string
	FilterByTags       int `gorm:"default:0"`
	Notifications      int `gorm:"default:1"`
	LastActive         string
	Password           string
}

type About struct {
	Email string
	About string
}

type Notifications struct {
	Email         string
	Notifications int
}

type DistancePreference struct {
	Email              string
	DistancePreference int
}

type AgePreference struct {
	Email         string
	AgePreference string
}

type FilterByTags struct {
	Email        string
	FilterByTags int
}

type ShowMe struct {
	Email  string
	ShowMe int
}

type LastActive struct {
	Email      string
	LastActive string
}

type UserUpdatePassword struct {
	Email       string
	OldPassword string
	NewPassword string
}

type UserDelete struct {
	Email string
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

// Update distance preference in users table in DB
func UpdateDistancePreference(db *gorm.DB, t *DistancePreference) error {
	return db.Table("users").Where("email = ?", t.Email).Update("distance_preference", t.DistancePreference).Error
}

// Update age preference in users table in DB
func UpdateAgePreference(db *gorm.DB, t *AgePreference) error {
	return db.Table("users").Where("email = ?", t.Email).Update("age_preference", t.AgePreference).Error
}

// Update filterByTags in users table in DB
func UpdateFilterByTags(db *gorm.DB, t *FilterByTags) error {
	return db.Table("users").Where("email = ?", t.Email).Update("filter_by_tags", t.FilterByTags).Error
}

// Update showMe in users table in DB
func UpdateShowMe(db *gorm.DB, t *ShowMe) error {
	return db.Table("users").Where("email = ?", t.Email).Update("show_me", t.ShowMe).Error
}

// Update lastActive in users table in DB
func UpdateLastActive(db *gorm.DB, t *LastActive) error {
	return db.Table("users").Where("email = ?", t.Email).Update("last_active", t.LastActive).Error
}

// Update password in users table in DB
func UpdatePassword(db *gorm.DB, t *UserUpdatePassword) (string, error) {
	r := db.Table("users").
		Where("email = ? AND password = ?", t.Email, t.OldPassword).
		Update("password", t.NewPassword)

	if r.Error != nil {
		return "Sorry, something went wrong 😕", r.Error
	}

	if r.RowsAffected == 0 {
		return "Old password is incorrect", nil
	}

	return "Successfully changed password 🎉", nil
}

// Delete user from users table in DB
func DeleteUser(db *gorm.DB, t *UserDelete) error {
	return db.Table("users").Where("email = ?", t.Email).Delete(&User{}).Error
}
