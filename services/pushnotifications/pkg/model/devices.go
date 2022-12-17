package devices

import "gorm.io/gorm"

type Device struct {
	Id          uint `gorm:"primary_key;auto_increment;not_null"`
	Email       string
	DeviceToken string
}

func (Device) TableName() string {
	return "devices"
}

func SaveDevice(db *gorm.DB, t *Device) error {
	exists := GetEntryByToken(db, t.DeviceToken)
	if !exists {
		return db.Create(t).Error
	}
	return nil
}

func GetEntryByToken(db *gorm.DB, token string) bool {
	var exists bool
	db.Table("devices").Select("count(*) > 0").Where("device_token = ?", token).Find(&exists)

	return exists
}
