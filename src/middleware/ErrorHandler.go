package middlewares

import (
	"net/http"

	customError "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/errors"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			switch e := err.(type) {
			case *customError.Error:
				c.JSON(e.HTTPStatusCode, gin.H{"error": e.Error()})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			}
		}
	}
}
