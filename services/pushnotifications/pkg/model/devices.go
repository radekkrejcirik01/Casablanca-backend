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
	entries := GetNumberOfEntries(db, t)
	if entries > 0 {
		return db.Model(&t).Where("email = ?", t.Email).Update("device_token", t.DeviceToken).Error
	} else {
		return db.Create(t).Error
	}
}

func GetNumberOfEntries(db *gorm.DB, t *Device) int {
	query := `SELECT COUNT(*) FROM devices WHERE email = '` + t.Email + `'`
	var result int
	db.Raw(query).Scan(&result)
	return result
}
