package routes

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/controllers/users"
	"github.com/gin-gonic/gin"
)

// AuthRoutes es la función que se encarga de definir las rutas de autenticación
func AuthRoutes(g *gin.RouterGroup /* controller *AuthController */) {
	// TODO: Implementar las rutas de autenticación
	g.Group("/auth")
	{
		g.POST("/login", users.Login)
	}
}
