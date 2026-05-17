package service

import (
    "errors"
    "time"
)


type JWTService struct{
    secretKey []byte
}

func NewJwtService(secret string) *JWTService {
    return &JWTService{
        secretKey: []byte(secret),
    }
}
