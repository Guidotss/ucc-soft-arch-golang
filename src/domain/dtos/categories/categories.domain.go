package categories

type CreateCategoryRequestDto struct {
	CategoryName string `json:"category_name"`
}

type CreateCategoryResponseDto struct {
	CategoryName string `json:"category_name"`
	CategoryId   string `json:"category_id"`
}
