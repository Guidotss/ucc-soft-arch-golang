package categories

import "github.com/google/uuid"

type CreateCategoryRequestDto struct {
	CategoryName string `json:"category_name"`
}

type CreateCategoryResponseDto struct {
	CategoryId   uuid.UUID `json:"category_id"`
	CategoryName string    `json:"category_name"`
}
