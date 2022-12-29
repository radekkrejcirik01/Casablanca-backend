package users

import "gorm.io/gorm"

type Like struct {
	Id    uint `gorm:"primary_key;auto_increment;not_null"`
	Email string
	User  string
	Value int
}

func (Like) TableName() string {
	return "likes"
}

// LikeUser like user
func LikeUser(db *gorm.DB, t *Like) error {
	if db.Model(&t).Where("email = ? AND user = ?", t.Email, t.User).Update("value", t.Value).RowsAffected == 0 {
		return db.FirstOrCreate(t, t).Error
	}
	return nil
}
