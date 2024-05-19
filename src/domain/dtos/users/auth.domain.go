package users

import "github.com/google/uuid"

type LoginRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginResponseDto struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Avatar   string    `json:"avatar"`
	Role     int       `json:"role"`
}
