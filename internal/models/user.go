package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserId    uint
	ProductId uint `json:"product_id" validate:"required"`
}

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Cart     []Cart `gorm:"foreignKey:UserID"`
}
