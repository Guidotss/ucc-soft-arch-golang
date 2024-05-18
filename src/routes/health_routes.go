package routes

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/controllers"
	"github.com/gin-gonic/gin"
)

func HealthRoutes(engine *gin.RouterGroup) {
	healthController := controllers.NewHealthController()

	engine.GET("/health", healthController.GetHealth)
}
