package users

import (
	"fmt"

	"errors"
	"net/http"
	"strings"

	customError "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/errors"
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

func (c *UsersClient) Create(user model.User) (model.User, error) {
	result := c.Db.Create(&user)
	if result.Error != nil {
		var err error
		switch {
		case strings.Contains(result.Error.Error(), "duplicate key value violates unique constraint"):
			err = customError.NewError(
				"DUPLICATE_IDENTIFIER",
				"A user with the same identifier or email already exists. Please use a different identifier or email.",
				http.StatusConflict)
		case strings.Contains(result.Error.Error(), "connection"):
			err = customError.NewError(
				"DB_CONNECTION_ERROR",
				"Database connection error. Please try again later.",
				http.StatusInternalServerError)
		default:
			err = customError.NewError(
				"UNEXPECTED_ERROR",
				"An unexpected error occurred. Please try again later.",
				http.StatusInternalServerError)
		}
		return model.User{}, err
	}
	return user, nil
}

func (c *UsersClient) FindById(id uuid.UUID) (model.User, error) {
	var user model.User
	err := c.Db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, customError.NewError("NOT_FOUND", "User not found", http.StatusNotFound)
		}
		return model.User{}, customError.NewError("DB_ERROR", "Error retrieving User from database", http.StatusInternalServerError)
	}
	return user, nil
}

func (c *UsersClient) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := c.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, customError.NewError("NOT_FOUND", "User not found", http.StatusNotFound)
		}
		return model.User{}, customError.NewError("DB_ERROR", "Error retrieving User from database", http.StatusInternalServerError)
	}
	fmt.Println("Result: ", user)
	return user, nil
}

func (c *UsersClient) UpdateUser(user model.User) (model.User, error) {
	result := c.Db.Table("users").Where("id = ?", user.Id).Updates(&user)
	if result.Error != nil {
		var err error
		switch {
		case strings.Contains(result.Error.Error(), "duplicate key value violates unique constraint"):
			err = customError.NewError(
				"DUPLICATE_IDENTIFIER",
				"A user with the same identifier or email already exists. Please use a different identifier or email.",
				http.StatusConflict)
		case strings.Contains(result.Error.Error(), "connection"):
			err = customError.NewError(
				"DB_CONNECTION_ERROR",
				"Database connection error. Please try again later.",
				http.StatusInternalServerError)
		default:
			err = customError.NewError(
				"UNEXPECTED_ERROR",
				"An unexpected error occurred. Please try again later.",
				http.StatusInternalServerError)
		}
		return model.User{}, err
	}
	return user, nil
}
