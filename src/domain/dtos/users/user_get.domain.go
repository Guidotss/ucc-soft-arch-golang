package users

import "github.com/google/uuid"

type GetUserDto struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Role     int       `json:"role"`
	UserName string    `json:"username"`
	Avatar   string    `json:"avatar"`
}

type GetAllUsersDto []GetUserDto
