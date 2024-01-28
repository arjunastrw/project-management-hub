package controller

import (
	"testing"

	"enigma.com/projectmanagementhub/mock/middleware_mock"
	"enigma.com/projectmanagementhub/mock/usecase_mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type UserControllerTestSuite struct {
	suite.Suite
	rg  *gin.RouterGroup
	uum *usecase_mock.UserUseCaseMock
	umm *middleware_mock.UserMiddlewareMock
}

func (a *UserControllerTestSuite) SetupTest() {
	a.umm = new(middleware_mock.UserMiddlewareMock)
	a.uum = new(usecase_mock.UserUseCaseMock)

	router := gin.Default()
	gin.SetMode(gin.TestMode)
	rg := router.Group("/pmh-api/v1")
	rg.Use(a.umm.RequireToken("ADMIN"))
	a.rg = rg
}

// Test GetAllUser
func (a *UserControllerTestSuite) TestFindAllUser_Success() {

}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
