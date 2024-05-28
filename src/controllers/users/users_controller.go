package users

import (
	"fmt"
	"net/http"

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
func (u *UsersController) FindByEmail(email string) users.GetUserDto {
	response := u.service.GetUserByEmail(email)
	return response
}

func (u *UsersController) CreateUser(g *gin.Context) {
	userEmail, exist := g.Get("Email")
	if !exist {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Email not found"})
		return
	}
	userName, exist := g.Get("Username")
	if !exist {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Email not found"})
		return
	}
	userPassword, exist := g.Get("Password")
	if !exist {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Email not found"})
		return
	}
	var user users.RegisterRequest
	user.Email = userEmail.(string)
	user.Username = userName.(string)
	user.Password = userPassword.(string)
	fmt.Println("UserRequest: ", user)

	response := u.service.CreateUser(user)
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
	userID, exists := g.Get("userID")
	if !exists {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found"})
		return
	}

	user.Id = userID.(uuid.UUID)

	err := g.BindJSON(&user)
	if err != nil {
		g.JSON(400, gin.H{
			"Ok":    false,
			"error": err.Error(),
		})
		return
	}

	response := u.service.UpdateUser(user)
	g.JSON(201, gin.H{
		"ok":      true,
		"message": "User updated successfully",
		"data":    response,
	})
}
