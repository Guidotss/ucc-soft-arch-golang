package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uuid.UUID `sql:"type:uuid;primary_key;default:gen_random_uuid()"`
	Password string    `gorm:"password"`
	Email    string    `gorm:"email;unique"`
	Role     int       `gorm:"role;default:0"`
	Name     string    `gorm:"user_name"`
	Avatar   string    `gorm:"avatar;default:https://i.postimg.cc/wTgNFWhR/profile.png"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New()
	return
}

type Users []User
