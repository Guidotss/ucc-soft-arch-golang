package users

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UsersClient struct {
	Db *gorm.DB
}

func NewUsersClient(db *gorm.DB) *UsersClient {
	return &UsersClient{Db: db}
}

func (u *UsersClient) Create(user model.User) model.User {
	result := u.Db.Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	return user
}

func (u *UsersClient) FindById(id uuid.UUID) model.User {
	var user model.User
	result := u.Db.First(&user, id)
	if result.Error != nil {
		panic(result.Error)
	}
	return user
}

func (u *UsersClient) FindByEmail(email string) model.User {
	var user model.User
	result := u.Db.First(&user, "email = ?", email)
	if result.Error != nil {
		panic(result.Error)
	}
	return user
}
