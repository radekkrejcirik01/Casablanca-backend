package controller

import (
	"github.com/radekkrejcirik01/Casblanca-backend/services/messages/pkg/model/matches"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseMatches struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    []matches.Matched
}
