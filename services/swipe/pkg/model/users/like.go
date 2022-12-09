package users

import "gorm.io/gorm"

type Like struct {
	Email string
	User  string
	Value int
}

func (Like) TableName() string {
	return "likes"
}

// LikeUser like user
func LikeUser(db *gorm.DB, t *Like) error {
	return db.Save(t).Error
}
