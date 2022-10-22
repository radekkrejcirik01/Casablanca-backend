package controller

import "github.com/radekkrejcirik01/Casblanca-backend/services/login-registration/pkg/model/users"

type resp struct {
	Status  string        `json:""`
	Message string        `json:",omitempty"`
	Data    *[]users.User `json:",omitempty"`
}
