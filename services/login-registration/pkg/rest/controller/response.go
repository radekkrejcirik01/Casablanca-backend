package controller

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type UserResponse struct {
	Status  string           `json:"status"`
	Message string           `json:"message,omitempty"`
	Data    UserDataResponse `json:"data,omitempty"`
}

type UserDataResponse struct {
	Email     string   `json:"email"`
	Firstname string   `json:"firstname"`
	Birthday  string   `json:"birthday"`
	Photos    []string `json:"photos"`
	Tags      []string `json:"tags"`
	Gender    string   `json:"gender"`
	ShowMe    string   `json:"showMe"`
}
