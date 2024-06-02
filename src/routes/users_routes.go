package routes

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/controllers/users"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/middleware/user"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/services"
	"github.com/gin-gonic/gin"
)

func UsersRoutes(engine *gin.Engine, controller *users.UsersController, service services.IUserService) {
	engine.POST("/users/register",
		user.RegisterInputCheckMiddleware(),
		user.IsEmailAvailable(service),
		controller.CreateUser)

	engine.PUT("/users/update", user.AuthMiddleware(), controller.UpdateUser)
}
