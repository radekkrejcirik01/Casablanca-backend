package users

import "gorm.io/gorm"

type Login struct {
	Id       uint   `json:""`
	Email    string `json:""`
	Password string `json:""`
}

type LoginDataResponse struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
}

// Login authenticate user
func LoginUser(db *gorm.DB, t *Login) error {
	return db.Table("users").Where("email = ? AND password = ?", t.Email, t.Password).First(t).Error
}
