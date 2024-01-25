package controller

import (
	"enigma.com/projectmanagementhub/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUC usecase.UserUseCase
	rg     *gin.RouterGroup
}

func NewUserController(rg *gin.RouterGroup, userUC usecase.UserUseCase) *UserController {
	return &UserController{
		userUC: userUC,
		rg:     rg,
	}
}

func (a *UserController) Route() {
	a.rg.GET("api/users", a.GetAllUser)
	a.rg.GET("api/user/:id", a.GetUserById)
	a.rg.GET("api/email/:email", a.GetUserByEmail)
	a.rg.POST("api/users", a.CreateUser)
	a.rg.PUT("api/user/:id", a.UpdateUser)
	a.rg.DELETE("api/user/:id", a.DeleteUser)
}
