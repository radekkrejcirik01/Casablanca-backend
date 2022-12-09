package controller

import (
	users "github.com/radekkrejcirik01/Casblanca-backend/services/swipe/pkg/model/users"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type UserResponse struct {
	Status  string           `json:"status"`
	Message string           `json:"message,omitempty"`
	Data    []users.UserData `json:"data,omitempty"`
}
