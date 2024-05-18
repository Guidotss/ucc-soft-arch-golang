package model

type Category struct {
	/* 	gorm.Model
	 */Id        string `gorm:"id_category"`
	CategoryName string `gorm:"category_name"`
}

type Categories []Category
