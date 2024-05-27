package user

import (
	"fmt"

	"github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/users"

	//"github.com/Guidotss/ucc-soft-arch-golang.git/src/services"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware verifica el token JWT
func RegisterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Entro al middleware")
		var user users.RegisterRequest
		err := c.BindJSON(&user)
		fmt.Println("UserRequest: ", user)
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
		fmt.Println("Paso register middleware")
		c.Set("Username", user.Username)
		c.Set("Email", user.Email)
		c.Set("Password", user.Password)
		c.Next()
	}
}
