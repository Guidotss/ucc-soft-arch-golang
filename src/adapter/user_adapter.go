package adapter

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/clients/users"
	controllers "github.com/Guidotss/ucc-soft-arch-golang.git/src/controllers/users"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/services"
	"gorm.io/gorm"
)

func UserAdapter(db *gorm.DB) (*controllers.UsersController, services.IUserService) {
	client := users.NewUsersClient(db)
	service := services.NewUserService(client)
	return controllers.NewUserController(service), service
}
