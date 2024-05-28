package services

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/clients/users"
	userDomain "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/users"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/utils/bcrypt"
	"github.com/google/uuid"
)

type UserService struct {
	client users.UsersClient
}

type IUserService interface {
	CreateUser(user userDomain.RegisterRequest) userDomain.RegisterResponse
	GetUserById(id uuid.UUID) userDomain.GetUserDto
	GetUserByEmail(email string) userDomain.GetUserDto
	UpdateUser(dto userDomain.UpdateRequestDto) userDomain.UpdateResponseDto
}

func NewUserService(client *users.UsersClient) IUserService {
	return &UserService{client: *client}
}

func (u *UserService) CreateUser(user userDomain.RegisterRequest) userDomain.RegisterResponse {
	hassedPassword, err := bcrypt.HasPassword(user.Password)
	if err != nil {
		panic(err)
	}
	var newUser = model.User{
		Password: hassedPassword,
		Email:    user.Email,
		Name:     user.Username,
		Avatar:   user.Avatar,
	}

	response := u.client.Create(newUser)

	return userDomain.RegisterResponse{
		Id:       response.Id,
		Email:    response.Email,
		Role:     response.Role,
		Username: response.Name,
		Avatar:   response.Avatar,
	}
}

func (u *UserService) GetUserById(id uuid.UUID) userDomain.GetUserDto {
	user := u.client.FindById(id)

	return userDomain.GetUserDto{
		Id:       user.Id,
		Email:    user.Email,
		Role:     user.Role,
		UserName: user.Name,
		Avatar:   user.Avatar,
	}
}

func (u *UserService) GetUserByEmail(email string) userDomain.GetUserDto {
	user := u.client.FindByEmail(email)

	return userDomain.GetUserDto{
		Id:       user.Id,
		Email:    user.Email,
		Role:     user.Role,
		UserName: user.Name,
		Avatar:   user.Avatar,
	}
}
func (u *UserService) UpdateUser(dto userDomain.UpdateRequestDto) userDomain.UpdateResponseDto {
	if dto.Password != "" {
		hassedPassword, err := bcrypt.HasPassword(dto.Password)
		if err != nil {
			panic(err)
		}
		dto.Password = hassedPassword
	}
	user := u.client.UpdateUser(dto)

	return userDomain.UpdateResponseDto{
		Id:       user.Id,
		Username: user.Name,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Role:     user.Role,
	}
}
