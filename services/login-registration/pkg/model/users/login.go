package users

import "gorm.io/gorm"

type Login struct {
	Email     string `json:""`
	Firstname string `json:""`
	Birthday  string `json:""`
	Gender    string `json:""`
	ShowMe    string `json:""`
	Password  string `json:""`
}

type LoginDataResponse struct {
	Email     string   `json:"email"`
	Firstname string   `json:"firstname"`
	Birthday  string   `json:"birthday"`
	Photos    []string `json:"photos"`
	Tags      []string `json:"tags"`
	Gender    string   `json:"gender"`
	ShowMe    string   `json:"showMe"`
}

// Login authenticate user
func LoginUser(db *gorm.DB, t *Login) error {
	return db.Table("users").Where("email = ? AND password = ?", t.Email, t.Password).First(t).Error
}
