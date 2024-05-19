package users

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/users"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/services"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/utils/jwt"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	service services.IUserService
}

type IUserController interface {
	CreateUser(g *gin.Context)
}

func NewUserController(service services.IUserService) *UsersController {
	return &UsersController{service: service}
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
	token := jwt.SignDocument(response.Id)
	g.JSON(201, gin.H{
		"ok":      true,
		"message": "User created successfully",
		"data":    response,
		"token":   token,
	})
}
