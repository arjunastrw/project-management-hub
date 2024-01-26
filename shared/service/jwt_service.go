package service

import (
	"fmt"
	"time"

	"enigma.com/projectmanagementhub/config"
	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/model/dto"
	"enigma.com/projectmanagementhub/shared/shared_model"
	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	GenerateToken(user model.User) (dto.AuthResponseDto, error)
	ParseToken(tokenHeader string) (jwt.MapClaims, error)
}

type jwtService struct {
	cfg config.TokenConfig
}

// GenerateToken implements JwtService.
func (j *jwtService) GenerateToken(user model.User) (dto.AuthResponseDto, error) {
	claims := shared_model.CustomClaims{
		UserID: user.Id,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfg.Issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.cfg.JwtExpiresTime)),
		},
	}

	token := jwt.NewWithClaims(j.cfg.JwtSigningMethod, claims)
	tokenString, err := token.SignedString(j.cfg.JwtSignatureKey)
	if err != nil {
		return dto.AuthResponseDto{}, fmt.Errorf("failed to generate token from jwtService.GenerateToken")
	}

	return dto.AuthResponseDto{Token: tokenString}, nil
}

func (j *jwtService) ParseToken(tokenHeader string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
		return j.cfg.JwtSignatureKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("oops, failed to verify token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("oops, failed to claim token")
	}
	return claims, nil
}

func NewJwtService(cfg config.TokenConfig) JwtService {
	return &jwtService{cfg: cfg}
}
