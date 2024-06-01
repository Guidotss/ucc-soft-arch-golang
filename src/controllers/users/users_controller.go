package users

import (
	"fmt"
	"net/http"

	customError "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/errors"

	"github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/users"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/services"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/utils/jwt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UsersController struct {
	service services.IUserService
}

func NewUserController(service services.IUserService) *UsersController {
	return &UsersController{service: service}
}
func (u *UsersController) FindByEmail(g *gin.Context) {
	email, exists := g.Get("email")
	if !exists {
		err := customError.NewError("EMAIL_NOT_FOUND", "Email not found", http.StatusBadRequest)
		g.Error(err)
		return
	}
	response, err := u.service.GetUserByEmail(email.(string))
	if err != nil {
		g.Error(err)
		return
	}
	g.JSON(200, gin.H{
		"ok":   true,
		"user": response,
	})
}

func (u *UsersController) CreateUser(g *gin.Context) {
	userEmail, _ := g.Get("Email")

	userName, _ := g.Get("Username")

	userPassword, _ := g.Get("Password")

	var user users.RegisterRequest
	user.Email = userEmail.(string)
	user.Username = userName.(string)
	user.Password = userPassword.(string)
	fmt.Println("UserRequest: ", user)

	response, err := u.service.CreateUser(user)
	if err != nil {
		g.Error(err)
		return
	}
	token := jwt.SignDocument(response.Id, response.Role)
	g.JSON(201, gin.H{
		"ok":      true,
		"message": "User created successfully",
		"user":    response,
		"token":   token,
	})
}

func (u *UsersController) UpdateUser(g *gin.Context) {
	var user users.UpdateRequestDto
	userID, _ := g.Get("userID")

	user.Id = userID.(uuid.UUID)

	err := g.BindJSON(&user)
	if err != nil {
		err := customError.NewError("INVALID_REQUEST", "Invalid request", http.StatusBadRequest)
		g.Error(err)
		return
	}

	response, err := u.service.UpdateUser(user)
	if err != nil {
		g.Error(err)
		return
	}

	g.JSON(201, gin.H{
		"ok":      true,
		"message": "User updated successfully",
		"data":    response,
	})
}
