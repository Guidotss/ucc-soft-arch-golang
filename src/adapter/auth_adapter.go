package adapter

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/clients/users"
	controller "github.com/Guidotss/ucc-soft-arch-golang.git/src/controllers/auth"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/services"
	"gorm.io/gorm"
)

func AuthAdapter(Db *gorm.DB) *controller.AuthController {
	client := users.NewUsersClient(Db)
	userService := services.NewUserService(client)
	authService := services.NewAuthService(&userService, client)
	return controller.NewAuthController(&authService)
}
