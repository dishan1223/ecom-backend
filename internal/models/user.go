package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RegisterRequest struct{
    Username string `json:"user_name" validate:"required,min=3,max=32"`
    Email string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=6,max=20"`
    Role string `json:"role" validate:"required,oneof=buyer seller"`
    PhoneNumber string `json:"phone_number" validate:"required,e164"`
}

type LoginRequest struct{
    Email string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}

type User struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"-"` 
	Role        string    `json:"role"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
}

type CustomClaims struct{
    UserID string `json:"user_id"`
    Role string `json:"role"`
    jwt.RegisteredClaims
}
