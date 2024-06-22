package model

type UserData struct {
	User_id   int    `json:"user_id"`
	User_Name string `json:"user_name"`
}

type UserProfile struct {
	User_id       int     `json:"user_id"`
	User_Name     string  `json:"user_name"`
	StatusMessage *string `json:"status_message"`
}
