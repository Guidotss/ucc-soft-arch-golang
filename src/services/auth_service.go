package services

import (
	"errors"
	"fmt"

	client "github.com/Guidotss/ucc-soft-arch-golang.git/src/clients/users"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/users"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/utils/bcrypt"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/utils/jwt"
	"github.com/google/uuid"
)

type AuthService struct {
	userService IUserService
	client      client.UsersClient
}

type IAuthService interface {
	RefreshToken(token string) (users.GetUserDto, string)
	Login(loginDto users.LoginRequestDto) (users.GetUserDto, string, error)
}

func NewAuthService(userService *IUserService, client *client.UsersClient) IAuthService {
	return &AuthService{
		userService: *userService,
		client:      *client,
	}
}

func (a *AuthService) RefreshToken(token string) (users.GetUserDto, string) {
	claims, err := jwt.VerifyToken(token)

	if err != nil {
		panic(err)
	}
	fmt.Println(claims)
	id, err := uuid.Parse(claims["id"].(string))
	roleInterface := claims["role"].(float64)
	role := int(roleInterface)
	if err != nil {
		panic(err)
	}

	checkUser := a.userService.GetUserById(id)

	if checkUser.Id == uuid.Nil {
		panic("User not found")
	}

	newToken := jwt.SignDocument(id, role)

	return checkUser, newToken
}

func (a *AuthService) Login(loginDto users.LoginRequestDto) (users.GetUserDto, string, error) {
	user := a.client.FindByEmail(loginDto.Email)
	var userDto users.GetUserDto
	if user.Id == uuid.Nil {
		err := errors.New("user not found")
		return users.GetUserDto{}, "", err
	}

	if !bcrypt.ComparePassword(loginDto.Password, user.Password) {
		err := errors.New("invalid Password")
		return users.GetUserDto{}, "", err
	}

	newToken := jwt.SignDocument(user.Id, user.Role)

	userDto = users.GetUserDto{
		Id:       user.Id,
		Email:    user.Email,
		Role:     user.Role,
		UserName: user.Name,
		Avatar:   user.Avatar,
	}

	return userDto, newToken, nil
}
