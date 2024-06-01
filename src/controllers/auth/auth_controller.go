package auth

import (
	"fmt"

	"github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/users"
	customError "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/errors"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service services.IAuthService
}

type IAuthController interface {
	RefreshToken(c *gin.Context)
	Login(c *gin.Context)
}

func NewAuthController(service *services.IAuthService) *AuthController {
	return &AuthController{service: *service}
}

func (a *AuthController) RefreshToken(c *gin.Context) {
	authorization := c.GetHeader("Authorization")

	if authorization == "" {
		err := customError.NewError("AUTHORIZATION_REQUIRED", "Authorization header is required", 400)
		c.Error(err)
		return
	}

	authorization = authorization[7:]

	user, token, err := a.service.RefreshToken(authorization)
	if err != nil {
		c.Error(err)
		return
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Content-Type", "application/json")
	c.JSON(200, gin.H{
		"ok":      true,
		"message": "Token refreshed",
		"token":   token,
		"user":    user,
	})
}

func (a *AuthController) Login(c *gin.Context) {
	var loginDto users.LoginRequestDto

	if err := c.ShouldBindJSON(&loginDto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("LoginDTO: ", loginDto)

	user, token, err := a.service.Login(loginDto)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{
		"ok":      true,
		"message": "User logged in",
		"user":    user,
		"token":   token,
	})
}
