package users

import "github.com/google/uuid"

type UpdateRequestDto struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Avatar   string    `json:"avatar"`
}

type UpdateResponseDto struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Avatar   string    `json:"avatar"`
	Role     int       `json:"role"`
}
