package response

import "time"

type RegisterRespone struct {
	User_id  int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

type UserToken struct {
	Token string `json:"token"`
}

type GetUserById struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

type UserUpdate struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	UpdatedAt time.Time	`json:"updated_at"`
}

type UserDelete struct {
	Message	string	`json:"message"`
}