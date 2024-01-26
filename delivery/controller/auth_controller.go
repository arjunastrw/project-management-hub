package controller

import (
	"net/http"

	"enigma.com/projectmanagementhub/model/dto"
	"enigma.com/projectmanagementhub/shared/common"
	"enigma.com/projectmanagementhub/usecase"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authUC usecase.AuthUsecase
	rg     *gin.RouterGroup
}

func (a *AuthController) loginHandler(c *gin.Context) {
	var payload dto.AuthRequestDto
	if err := c.ShouldBind(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := a.authUC.Login(payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreatedResponse(c, response, "Login Success")
}

func (a *AuthController) Route() {
	a.rg.POST("/login", a.loginHandler)
}

func NewAuthController(authUC usecase.AuthUsecase, rg *gin.RouterGroup) *AuthController {
	return &AuthController{
		authUC: authUC,
		rg:     rg,
	}
}
