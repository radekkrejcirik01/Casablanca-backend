package users

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Photos struct {
	Email  string
	Photos pq.StringArray `gorm:"type:text[]"`
}

type Photo struct {
	Id    uint `gorm:"primary_key;auto_increment;not_null"`
	Email string
	Photo string
}

func (Photo) TableName() string {
	return "photos"
}

// Add new photos to DB
func AddPhotos(db *gorm.DB, t *Photos) error {
	result := make([]Photo, 0)
	for _, photo := range t.Photos {
		result = append(result, Photo{Email: t.Email, Photo: photo})
	}
	return db.Create(result).Error
}

// Read Photos from DB by user
func GetPhotos(db *gorm.DB, t *[]string, user string) error {
	return db.Table("photos").Select("photo").Where("email = ?", user).Find(t).Error
}

// Update photos in DB by user
func UpdatePhotos(db *gorm.DB, t *Photos) error {
	photos := make([]Photo, 0)
	for _, photo := range t.Photos {
		photos = append(photos, Photo{Email: t.Email, Photo: photo})
	}

	profilePicture := photos[0].Photo

	if err := db.Table("photos").Where("email = ?", t.Email).Delete(t).Error; err != nil {
		return err
	}
	if err := db.Create(photos).Error; err != nil {
		return err
	}
	if err := db.Table("users").Where("email = ?", t.Email).
		Update("profile_picture", profilePicture).Error; err != nil {
		return err
	}

	return nil
}
