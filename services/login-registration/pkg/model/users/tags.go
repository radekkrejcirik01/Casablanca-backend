package users

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Tags struct {
	Email string
	Tags  pq.StringArray `gorm:"type:text[]"`
}

type Tag struct {
	Id    uint `gorm:"primary_key;auto_increment;not_null"`
	Email string
	Tag   string
}

func (Tag) TableName() string {
	return "tags"
}

// Add new tags to DB
func AddTags(db *gorm.DB, t *Tags) error {
	result := make([]Tag, 0)
	for _, tag := range t.Tags {
		result = append(result, Tag{Email: t.Email, Tag: tag})
	}
	return db.Create(result).Error
}

// Read tags from DB by user
func GetTags(db *gorm.DB, t *[]string, user string) error {
	return db.Table("tags").Select("tag").Where("email = ?", user).Find(t).Error
}

// Update tags in DB by user
func UpdateTags(db *gorm.DB, t *Tags) error {
	tags := make([]Tag, 0)
	for _, tag := range t.Tags {
		tags = append(tags, Tag{Email: t.Email, Tag: tag})
	}

	if err := db.Table("tags").Where("email = ?", t.Email).Delete(t).Error; err != nil {
		return err
	}
	if err := db.Create(tags).Error; err != nil {
		return err
	}

	return nil
}
