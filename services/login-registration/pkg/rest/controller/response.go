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
	Email              string   `json:"email"`
	Firstname          string   `json:"firstname"`
	Birthday           string   `json:"birthday"`
	ProfilePicture     string   `json:"profilePicture"`
	About              string   `json:"about"`
	Photos             []string `json:"photos"`
	Tags               []string `json:"tags"`
	Gender             int      `json:"gender"`
	ShowMe             int      `json:"showMe"`
	DistancePreference int      `json:"distancePreference"`
	AgePreference      string   `json:"agePreference"`
	FilterByTags       int      `json:"filterByTags"`
	Notifications      int      `json:"notifications"`
}
