package controllers

import "github.com/gin-gonic/gin"

type HealthController interface {
	GetHealth(c *gin.Context)
}

type healthControllerImpl struct{}

func NewHealthController() HealthController {
	return &healthControllerImpl{}
}

func (h *healthControllerImpl) GetHealth(c *gin.Context) {
	c.JSON(200, gin.H{
		"ok":      true,
		"status":  "UP",
		"message": "The service is running",
	})
}
