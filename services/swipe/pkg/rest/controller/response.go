package controller

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type UserResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message,omitempty"`
	Data    UsersResponse `json:"data,omitempty"`
}

type UsersResponse struct {
	Id        uint     `json:"id"`
	Email     string   `json:"email"`
	Firstname string   `json:"firstname"`
	Birthday  string   `json:"birthday"`
	About     string   `json:"about"`
	Photos    []string `json:"photos"`
	Tags      []string `json:"tags"`
	Gender    int      `json:"gender"`
	Distance  int      `json:"distance"`
}
