package user

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/users"
	customError "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/errors"

	"github.com/gin-gonic/gin"
)

func RegisterInputCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user users.RegisterRequest
		err := c.BindJSON(&user)
		if err != nil {
			err := customError.NewError("INVALID_REQUEST", "Invalid request", 400)
			c.Error(err)
			c.Abort()
			return
		}
		if user.Password == "" || user.Email == "" || user.Username == "" {
			err := customError.NewError("INVALID_REQUEST", "Invalid request", 400)
			c.Error(err)
			c.Abort()
			return
		}
		c.Set("Username", user.Username)
		c.Set("Email", user.Email)
		c.Set("Password", user.Password)
		c.Next()
	}
}
