package categories

import "github.com/google/uuid"

type GetAllCategoriesResponseDto struct {
	CategoryId   uuid.UUID `json:"category_id"`
	CategoryName string    `json:"category_name"`
}
type GetAllCategories []GetAllCategoriesResponseDto
