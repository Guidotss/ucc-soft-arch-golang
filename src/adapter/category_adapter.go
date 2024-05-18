package adapter

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/clients/categories"
	controller "github.com/Guidotss/ucc-soft-arch-golang.git/src/controllers/categories"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/services"
	"gorm.io/gorm"
)

func CategoryAdapter(Db *gorm.DB) *controller.CategoriesController {
	client := categories.NewCategoryClient(Db)
	service := services.NewCategoriesService(client)
	return controller.NewCategoriesController(service)
}
