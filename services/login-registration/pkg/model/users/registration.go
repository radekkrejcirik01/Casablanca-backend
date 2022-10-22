package users

import (
	"github.com/lib/pq"
)

type Registration struct {
	Id        uint           `gorm:"primary_key;auto_increment;not_null"`
	Firstname string         `json:""`
	Birthday  string         `json:""`
	Photos    pq.StringArray `gorm:"type:text[]"`
	Tags      pq.StringArray `gorm:"type:text[]"`
	Gender    string         `json:""`
	ShowMe    string         `json:""`
	Email     string         `json:""`
	Password  string         `json:""`
}
