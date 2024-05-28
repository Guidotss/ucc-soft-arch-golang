package routes

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/controllers/users"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/middleware/user"
	"github.com/gin-gonic/gin"
)

func UsersRoutes(engine *gin.Engine, controller *users.UsersController) {
	engine.POST("/users/register", user.RegisterMiddleware(), controller.CreateUser)
	engine.PUT("/users/update", user.AuthMiddleware(), controller.UpdateUser)
}
