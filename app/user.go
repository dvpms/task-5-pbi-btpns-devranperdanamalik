package app

type User struct {
	UserID   int    `json:"user_id"`
	USERNAME string `json:"username" `
	EMAIL    string `json:"email"`
	PASSWORD string `json:"password"`
}
