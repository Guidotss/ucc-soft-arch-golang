package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       string `gorm:"uid"`
	Password string `gorm:"password"`
	Email    string `gorm:"email"`
	Role     int    `gorm:"role"`
	Name     string `gorm:"user_name"`
	Avatar   string `gorm:"avatar"`
}

type Users []User
