package routes

import (
	controller "github.com/Guidotss/ucc-soft-arch-golang.git/src/controllers/auth"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(engine *gin.Engine, controller *controller.AuthController) {
	engine.POST("/auth/refresh-token", controller.RefreshToken)
	engine.POST("/auth/login", controller.Login)
}
