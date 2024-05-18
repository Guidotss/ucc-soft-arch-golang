package services

import (
	categoriesClient "github.com/Guidotss/ucc-soft-arch-golang.git/src/clients/categories"
	categoriesDto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/categories"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
)

type ICategoriesService interface {
	CreateCategory(categoryDto categoriesDto.CreateCategoryRequestDto) categoriesDto.CreateCategoryResponseDto
	FindAllCategories() categoriesDto.GetAllCategories
}

type categoriesService struct {
	client categoriesClient.CategoriesClient
}

func NewCategoriesService(client *categoriesClient.CategoriesClient) ICategoriesService {
	return &categoriesService{client: *client}
}

func (c *categoriesService) CreateCategory(categoryDto categoriesDto.CreateCategoryRequestDto) categoriesDto.CreateCategoryResponseDto {
	var newCategory = model.Category{
		CategoryName: categoryDto.CategoryName,
	}

	createdCategory := c.client.Create(newCategory)

	return categoriesDto.CreateCategoryResponseDto{
		CategoryId:   createdCategory.Id,
		CategoryName: createdCategory.CategoryName,
	}
}
func (c *categoriesService) FindAllCategories() categoriesDto.GetAllCategories {
	var categories model.Categories = c.client.GetAll()
	var allCartegoriesDto categoriesDto.GetAllCategories
	for _, result := range categories {
		var categoryDto categoriesDto.GetAllCategoriesResponseDto
		categoryDto.CategoryId = result.Id
		categoryDto.CategoryName = result.CategoryName
		allCartegoriesDto = append(allCartegoriesDto, categoryDto)
	}
	return allCartegoriesDto
}
