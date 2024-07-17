package user

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/services"

	"github.com/gin-gonic/gin"
)

func IsEmailAvailable(service services.IUserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.GetString("Email")
		_, err := service.GetUserByEmail(email)
		if err == nil {
			c.Error(err)
			c.Abort()
			return
		}
		c.Next()
	}
}
