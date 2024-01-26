package usecase

import (
	"fmt"

	"enigma.com/projectmanagementhub/model/dto"
	"enigma.com/projectmanagementhub/shared/service"
)

type AuthUsecase interface {
	Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error)
}

type authUsecase struct {
	userUC     UserUseCase
	jwtService service.JwtService
}

// Login implements AuthUsecase.
func (a *authUsecase) Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error) {
	user, err := a.userUC.FindUserByEmail(payload.Email)
	if err != nil {
		return dto.AuthResponseDto{}, fmt.Errorf("email not found")
	}

	if user.Password != payload.Password {
		return dto.AuthResponseDto{}, fmt.Errorf("wrong password")
	}

	tokenDto, err := a.jwtService.GenerateToken(user)
	if err != nil {
		return dto.AuthResponseDto{}, fmt.Errorf("failed to generate token from authUsecase.Login")
	}

	return tokenDto, nil
}

func NewAuthUsecase(userUC UserUseCase, jwtService service.JwtService) AuthUsecase {
	return &authUsecase{
		userUC:     userUC,
		jwtService: jwtService,
	}
}
