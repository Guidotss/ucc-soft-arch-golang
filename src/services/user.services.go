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
	CreateUser(user userDomain.RegisterRequest) (userDomain.RegisterResponse, error)
	GetUserById(id uuid.UUID) (userDomain.GetUserDto, error)
	GetUserByEmail(email string) (userDomain.GetUserDto, error)
	UpdateUser(dto userDomain.UpdateRequestDto) (userDomain.UpdateResponseDto, error)
}

func NewUserService(client *users.UsersClient) IUserService {
	return &UserService{client: *client}
}

func (u *UserService) CreateUser(user userDomain.RegisterRequest) (userDomain.RegisterResponse, error) {
	hassedPassword, err := bcrypt.HasPassword(user.Password)
	if err != nil {
		return userDomain.RegisterResponse{}, err
	}

	var newUser = model.User{
		Password: hassedPassword,
		Email:    user.Email,
		Name:     user.Username,
		Avatar:   user.Avatar,
	}

	response, err := u.client.Create(newUser)
	if err != nil {
		return userDomain.RegisterResponse{}, err
	}

	return userDomain.RegisterResponse{
		Id:       response.Id,
		Email:    response.Email,
		Role:     response.Role,
		Username: response.Name,
		Avatar:   response.Avatar,
	}, nil
}

func (u *UserService) GetUserById(id uuid.UUID) (userDomain.GetUserDto, error) {
	user, err := u.client.FindById(id)
	if err != nil {
		return userDomain.GetUserDto{}, err
	}
	return userDomain.GetUserDto{
		Id:       user.Id,
		Email:    user.Email,
		Role:     user.Role,
		UserName: user.Name,
		Avatar:   user.Avatar,
	}, nil
}

func (u *UserService) GetUserByEmail(email string) (userDomain.GetUserDto, error) {
	user, err := u.client.FindByEmail(email)
	if err != nil {
		return userDomain.GetUserDto{}, err
	}

	return userDomain.GetUserDto{
		Id:       user.Id,
		Email:    user.Email,
		Role:     user.Role,
		UserName: user.Name,
		Avatar:   user.Avatar,
	}, nil
}
func (u *UserService) UpdateUser(dto userDomain.UpdateRequestDto) (userDomain.UpdateResponseDto, error) {
	var user model.User
	user.Id = dto.Id
	if dto.Password != "" {
		hassedPassword, _ := bcrypt.HasPassword(dto.Password)
		dto.Password = hassedPassword
	}
	if dto.Username != "" {
		user.Name = dto.Username
	}
	if dto.Email != "" {
		user.Email = dto.Email
	}
	if dto.Avatar != "" {
		user.Avatar = dto.Avatar
	}

	user, err := u.client.UpdateUser(user)
	if err != nil {
		return userDomain.UpdateResponseDto{}, err
	}

	return userDomain.UpdateResponseDto{
		Id:       user.Id,
		Username: user.Name,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Role:     user.Role,
	}, nil
}
