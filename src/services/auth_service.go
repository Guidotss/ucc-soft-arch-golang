package services

import (
	"fmt"

	client "github.com/Guidotss/ucc-soft-arch-golang.git/src/clients/users"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/users"
	customError "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/errors"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/utils/bcrypt"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/utils/jwt"
	"github.com/google/uuid"
)

type AuthService struct {
	userService IUserService
	client      client.UsersClient
}

type IAuthService interface {
	RefreshToken(token string) (users.GetUserDto, string, error)
	Login(loginDto users.LoginRequestDto) (users.GetUserDto, string, error)
}

func NewAuthService(userService *IUserService, client *client.UsersClient) IAuthService {
	return &AuthService{
		userService: *userService,
		client:      *client,
	}
}

func (a *AuthService) RefreshToken(token string) (users.GetUserDto, string, error) {
	claims, err := jwt.VerifyToken(token)
	if err != nil {
		return users.GetUserDto{}, "", customError.NewError("INVALID TOKEN", "Invalid token", 401)
	}

	fmt.Println(claims)
	id, err := uuid.Parse(claims["id"].(string))
	if err != nil {
		return users.GetUserDto{}, "", customError.NewError("INVALID ID", "Invalid ID", 401)
	}

	roleInterface := claims["role"].(float64)
	role := int(roleInterface)

	checkUser, err := a.userService.GetUserById(id)
	if err != nil {
		return users.GetUserDto{}, "", err
	}

	if checkUser.Id == uuid.Nil {
		return users.GetUserDto{}, "", customError.NewError("USER NOT FOUND", "User not found", 404)
	}

	newToken := jwt.SignDocument(id, role)

	return checkUser, newToken, nil
}

func (a *AuthService) Login(loginDto users.LoginRequestDto) (users.GetUserDto, string, error) {
	var userDto users.GetUserDto
	user, err := a.client.FindByEmail(loginDto.Email)
	if err != nil {
		return users.GetUserDto{}, "", err
	}

	if !bcrypt.ComparePassword(loginDto.Password, user.Password) {
		return users.GetUserDto{}, "", customError.NewError("INVALID CREDENTIALS", "Invalid credentials", 401)
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
