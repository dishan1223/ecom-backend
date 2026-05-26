package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model

	CoverImage  string `json:"cover" validate:"required, url"`
	BookName    string `json:"book_name" validate:"required"`
	Writer      string `json:"write_name" validate:"required"`
	Description string `json:"description" validate:"required"`

	BestSeller bool   `json:"best_seller" validate:"boolean"`
	Category   string `json:"category" validate:"required"`

	Like      uint   `json:"like" gorm:"default:0" validate:"lte=0|gte=0`
	TotalSold uint64 `json:"total_sold" gorm:"default:0" validate:"lte=0|gte=0"`

	OriginalPrice uint `json:"original_price" validate:"numeric"`
	CurrentPrice  uint `json:"current_price" validate:"numeric"`

	Stock uint `json:"stock" validate:"numeric"`

	ProductType string `json:"product_type" validate:"required"`
}
