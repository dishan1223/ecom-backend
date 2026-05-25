package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model

	CoverImage  string `json:"cover"`
	BookName    string `json:"book_name"`
	Writer      string `json:"write_name"`
	Description string `json:"description"`

	BestSeller bool   `json:"best_seller"`
	Category   string `json:"category"`

	Like      uint   `json:"like" gorm:"default:0"`
	TotalSold uint64 `json:"total_sold" gorm:"default:0"`

	OriginalPrice uint `json:"original_price"`
	CurrentPrice  uint `json:"current_price"`

	Stock uint `json:"stock"`

	ProductType string `json:"product_type"`
}
