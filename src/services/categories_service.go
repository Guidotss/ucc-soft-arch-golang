package services

import (
	categoriesClient "github.com/Guidotss/ucc-soft-arch-golang.git/src/clients/categories"
	categoriesDto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/categories"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
)

type ICategoriesService interface {
	CreateCategory(categoryDto categoriesDto.CreateCategoryRequestDto) (categoriesDto.CreateCategoryResponseDto, error)
	FindAllCategories() (categoriesDto.GetAllCategories, error)
}

type categoriesService struct {
	client categoriesClient.CategoriesClient
}

func NewCategoriesService(client *categoriesClient.CategoriesClient) ICategoriesService {
	return &categoriesService{client: *client}
}

func (c *categoriesService) CreateCategory(categoryDto categoriesDto.CreateCategoryRequestDto) (categoriesDto.CreateCategoryResponseDto, error) {
	var newCategory = model.Category{
		CategoryName: categoryDto.CategoryName,
	}

	createdCategory, err := c.client.Create(newCategory)
	if err != nil {
		return categoriesDto.CreateCategoryResponseDto{}, err
	}

	return categoriesDto.CreateCategoryResponseDto{
		CategoryId:   createdCategory.Id,
		CategoryName: createdCategory.CategoryName,
	}, nil
}
func (c *categoriesService) FindAllCategories() (categoriesDto.GetAllCategories, error) {
	categories, err := c.client.GetAll()
	if err != nil {
		return nil, err
	}
	var allCartegoriesDto categoriesDto.GetAllCategories
	for _, result := range categories {
		var categoryDto categoriesDto.GetAllCategoriesResponseDto
		categoryDto.CategoryId = result.Id
		categoryDto.CategoryName = result.CategoryName
		allCartegoriesDto = append(allCartegoriesDto, categoryDto)
	}
	return allCartegoriesDto, nil
}
