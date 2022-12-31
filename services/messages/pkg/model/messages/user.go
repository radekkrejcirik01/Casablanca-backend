package messages

import (
	"gorm.io/gorm"
)

type UserBody struct {
	User string
}

type ChatUser struct {
	Firstname  string   `json:"firstname"`
	Birthday   string   `json:"birthday"`
	About      string   `json:"about"`
	Photos     []string `json:"photos"`
	Tags       []string `json:"tags"`
	LastActive string   `json:"lastActive"`
}

type UserData struct {
	Firstname  string
	Birthday   string
	About      string
	LastActive string
}

// GetUser get chat user
func GetUser(db *gorm.DB, t *UserBody) (ChatUser, error) {
	user := UserData{}
	photos := &[]string{}
	tags := &[]string{}

	if err := GetUserByUser(db, &user, t.User); err != nil {
		return ChatUser{}, err
	}

	if err := GetPhotos(db, photos, t.User); err != nil {
		return ChatUser{}, err
	}

	if err := GetTags(db, tags, t.User); err != nil {
		return ChatUser{}, err
	}

	result := ChatUser{
		Firstname:  user.Firstname,
		Birthday:   user.Birthday,
		About:      user.About,
		Photos:     *photos,
		Tags:       *tags,
		LastActive: user.LastActive,
	}

	return result, nil
}

func GetUserByUser(db *gorm.DB, t *UserData, user string) error {
	return db.Table("users").Where("email = ?", user).First(t).Error
}

func GetPhotos(db *gorm.DB, t *[]string, user string) error {
	return db.Table("photos").Select("photo").Where("user = ?", user).Find(t).Error
}

func GetTags(db *gorm.DB, t *[]string, user string) error {
	return db.Table("tags").Select("tag").Where("user = ?", user).Find(t).Error
}
