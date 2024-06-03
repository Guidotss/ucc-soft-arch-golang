package categories

import (
	"net/http"

	categoryDomain "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/categories"
	customError "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/errors"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/services"
	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	CategoriesService services.ICategoriesService
}

type ICategoriesController interface {
	Create(g *gin.Context)
}

func NewCategoriesController(service services.ICategoriesService) *CategoriesController {
	return &CategoriesController{CategoriesService: service}
}

func (c *CategoriesController) Create(g *gin.Context) {
	var categoryDto categoryDomain.CreateCategoryRequestDto
	err := g.BindJSON(&categoryDto)
	if err != nil {
		g.Error(customError.NewError("INVALID_FIELDS", "Error when loading the inputs", http.StatusBadRequest))
		return
	}
	response, err := c.CategoriesService.CreateCategory(categoryDto)
	if err != nil {
		g.Error(err)
		return
	}
	g.JSON(201, response)
}

func (c *CategoriesController) GetAll(g *gin.Context) {
	response, err := c.CategoriesService.FindAllCategories()
	if err != nil {
		g.Error(err)
		return
	}
	g.JSON(200, response)
}
