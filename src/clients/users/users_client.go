package users

import (
	dto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/users"
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

func (u *UsersClient) UpdateUser(dto dto.UpdateRequestDto) model.User {
	var user model.User
	result := u.Db.First(&user, dto.Id)
	if result.Error != nil {
		//manejo de errores
		panic(result.Error)
	}
	if dto.Email != "" {
		user.Email = dto.Email
	}
	if dto.Username != "" {
		user.Name = dto.Username
	}
	if dto.Password != "" {
		user.Password = dto.Password
	}
	if dto.Avatar != "" {
		user.Avatar = dto.Avatar
	}
	result = u.Db.Save(&user)
	if result.Error != nil {
		//manejo de errores
		panic(result.Error)
	}
	return user
}
