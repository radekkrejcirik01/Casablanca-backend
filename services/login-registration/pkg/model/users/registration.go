package users

import (
	"github.com/lib/pq"
)

type Registration struct {
	Firstname string         `json:""`
	Birthday  string         `json:""`
	Photos    pq.StringArray `gorm:"type:text[]"`
	Tags      pq.StringArray `gorm:"type:text[]"`
	Gender    string         `json:""`
	ShowMe    string         `json:""`
	Email     string         `json:""`
	Password  string         `json:""`
}

type RegistrationDataResponse struct {
	Email string `json:"email"`
}
