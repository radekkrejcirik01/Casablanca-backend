package users

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Tags struct {
	Id   uint           `gorm:"primary_key;auto_increment;not_null"`
	User string         `json:""`
	Tags pq.StringArray `gorm:"type:text[]"`
}

type Tag struct {
	Id   uint   `gorm:"primary_key;auto_increment;not_null"`
	User string `json:""`
	Tag  string `json:""`
}

func (Tag) TableName() string {
	return "tags"
}

// Create new Tag in DB
func CreateTag(db *gorm.DB, t *Tags) error {
	result := make([]Tag, 0)
	for _, tag := range t.Tags {
		result = append(result, Tag{Id: t.Id, User: t.User, Tag: tag})
	}
	return db.Create(result).Error
}
