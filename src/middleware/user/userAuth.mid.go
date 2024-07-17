package user

import (
	"fmt"
	"strings"

	"github.com/Guidotss/ucc-soft-arch-golang.git/src/config"
	customError "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/errors"
	utilsJWT "github.com/Guidotss/ucc-soft-arch-golang.git/src/utils/jwt"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt"
)

// AuthMiddleware verifica el token JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		fmt.Println("auth:", authHeader)
		if authHeader == "" {
			err := customError.NewError("AUTHORIZATION_REQUIRED", "Authorization header is required", 400)
			c.Error(err)
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, "Bearer ")
		if len(tokenParts) != 2 {
			err := customError.NewError("INVALID_TOKEN", "Invalid token", 400)
			c.Error(err)
			c.Abort()
			return
		}

		tokenString := tokenParts[1]

		claims := &utilsJWT.CustomClaims{}
		envs := config.LoadEnvs(".env")
		secret := []byte(envs.Get("JWT_SECRET"))

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})

		if err != nil {
			err := customError.NewError("INVALID_TOKEN", "Invalid token", 400)
			c.Error(err)
			c.Abort()
			return
		}

		if !token.Valid {
			err := customError.NewError("INVALID_TOKEN", "Invalid token", 400)
			c.Error(err)
			c.Abort()
			return
		}

		c.Set("userID", claims.Id)
		c.Next()
	}
}
