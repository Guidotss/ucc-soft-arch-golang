package users

import "github.com/google/uuid"

type RegisterRequest struct {
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Avatar   string `json:"avatar"`
}

type RegisterResponse struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Avatar   string    `json:"avatar"`
	Role     int       `json:"role"`
}
