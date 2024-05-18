package routes

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/controllers/categories"
	"github.com/gin-gonic/gin"
)

func CategoriesRoutes(engine *gin.Engine, controller *categories.CategoriesController) {
	engine.POST("/category/create", controller.Create)
	engine.GET("/categories", controller.GetAll)
}
