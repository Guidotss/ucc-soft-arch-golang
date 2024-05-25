package auth

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/users"
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
		c.JSON(400, gin.H{"error": "Authorization header is required"})
		return
	}

	authorization = authorization[7:]

	user, token := a.service.RefreshToken(authorization)

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

	user, token := a.service.Login(loginDto)

	c.JSON(200, gin.H{
		"ok":      true,
		"message": "User logged in",
		"user":    user,
		"token":   token,
	})
}
