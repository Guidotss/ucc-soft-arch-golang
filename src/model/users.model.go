package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uuid.UUID `sql:"type:uuid;primary_key;default:gen_random_uuid()"`
	Password string    `gorm:"password"`
	Email    string    `gorm:"email"`
	Role     int       `gorm:"role"`
	Name     string    `gorm:"user_name"`
	Avatar   string    `gorm:"avatar"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New()
	return
}

type Users []User
