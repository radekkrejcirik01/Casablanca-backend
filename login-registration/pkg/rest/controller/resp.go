package controller

import "login-registration/pkg/model/users"

type resp struct {
	Status  string         `json:""`
	Message string         `json:",omitempty"`
	Data    *[]users.USERS `json:",omitempty"`
}
