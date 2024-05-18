package services

import (
	categoriesClient "github.com/Guidotss/ucc-soft-arch-golang.git/src/clients/categories"
	categoriesDto "github.com/Guidotss/ucc-soft-arch-golang.git/src/domain/dtos/categories"
	"github.com/Guidotss/ucc-soft-arch-golang.git/src/model"
)

type ICategoriesService interface {
	CreateCategory(categoryDto categoriesDto.CreateCategoryRequestDto) categoriesDto.CreateCategoryResponseDto
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
		CategoryName: createdCategory.CategoryName,
		CategoryId:   createdCategory.Id,
	}
}
