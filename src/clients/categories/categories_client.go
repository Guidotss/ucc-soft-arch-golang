package categories

import (
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CategoriesClient struct {
	Db *gorm.DB
}

func NewCategoryClient(db *gorm.DB) *CategoriesClient {
	return &CategoriesClient{Db: db}
}

func (c *CategoriesClient) Create(category model.Category) model.Category {
	result := c.Db.Create(&category)

	if result.Error != nil {
		log.Error()
	}
	log.Debug("Categoria creado con exito wachin, su id es: ", result)
	return category
}

func (c *CategoriesClient) GetAll() model.Categories {
	var categories model.Categories
	c.Db.Find(&categories)
	return categories
}
