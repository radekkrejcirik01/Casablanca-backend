package users

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Photos struct {
	User   string         `json:""`
	Photos pq.StringArray `gorm:"type:text[]"`
}

type Photo struct {
	Id    uint   `gorm:"primary_key;auto_increment;not_null"`
	User  string `json:""`
	Photo string `json:""`
}

func (Photo) TableName() string {
	return "photos"
}

// Create new Photo in DB
func CreatePhoto(db *gorm.DB, t *Photos) error {
	result := make([]Photo, 0)
	for _, photo := range t.Photos {
		result = append(result, Photo{User: t.User, Photo: photo})
	}
	return db.Create(result).Error
}

// Read Photos from DB by user
func GetPhotos(db *gorm.DB, t *[]string, user string) error {
	return db.Table("photos").Select("photo").Where("user = ?", user).Find(t).Error
}
