package shared_model

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}
