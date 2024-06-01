package categories

import (
	"errors"
	"net/http"
	"strings"

	customError "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/errors"
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

func (c *CategoriesClient) Create(category model.Category) (model.Category, error) {
	result := c.Db.Create(&category)
	if result.Error != nil {
		var err error
		switch {
		case strings.Contains(result.Error.Error(), "duplicate key value violates unique constraint"):
			err = customError.NewError(
				"DUPLICATE_IDENTIFIER",
				"A category with the same identifier already exists. Please use a different identifier.",
				http.StatusConflict)
		case strings.Contains(result.Error.Error(), "connection"):
			err = customError.NewError(
				"DB_CONNECTION_ERROR",
				"Database connection error. Please try again later.",
				http.StatusInternalServerError)
		default:
			err = customError.NewError(
				"UNEXPECTED_ERROR",
				"An unexpected error occurred. Please try again later.",
				http.StatusInternalServerError)
		}
		return model.Category{}, err
	}
	log.Debug("Categoria creado con exito wachin, su id es: ", result)
	return category, nil
}

func (c *CategoriesClient) GetAll() (model.Categories, error) {
	var categories model.Categories
	err := c.Db.Find(&categories).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customError.NewError("NOT_FOUND", "There is no courses", http.StatusNotFound)
		}
		return nil, customError.NewError("DB_ERROR", "Error retrieving course from database", http.StatusInternalServerError)
	}
	return categories, nil
}
