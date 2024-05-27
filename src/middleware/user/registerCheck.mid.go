package user

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/users"

	//"github.com/Guidotss/ucc-soft-arch-golang.git/src/services"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware verifica el token JWT
func RegisterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user users.RegisterRequest
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(400, gin.H{
				"Ok":    false,
				"error": "Invalid request",
			})
			c.Abort()
			return
		}
		if user.Password == "" || user.Email == "" || user.Username == "" {
			c.JSON(400, gin.H{
				"Ok":    false,
				"error": "All fields are required",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
