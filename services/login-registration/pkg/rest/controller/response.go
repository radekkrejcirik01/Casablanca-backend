package controller

import "github.com/radekkrejcirik01/Casblanca-backend/services/login-registration/pkg/model/users"

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type RegistrationResponse struct {
	Status  string                          `json:"status"`
	Message string                          `json:"message,omitempty"`
	Data    *users.RegistrationDataResponse `json:"data,omitempty"`
}

type UserGetResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message,omitempty"`
	Data    *[]users.User `json:"data,omitempty"`
}
