package users

import (
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
	var users users.RegisterRequest
	err := g.BindJSON(&users)
	if err != nil {
		g.JSON(400, gin.H{
			"Ok":    false,
			"error": err.Error(),
		})
		return
	}

	response := u.service.CreateUser(users)
	token := jwt.SignDocument(response.Id, response.Role)
	g.JSON(201, gin.H{
		"ok":      true,
		"message": "User created successfully",
		"data":    response,
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
