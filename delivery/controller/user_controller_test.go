package controller

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"enigma.com/projectmanagementhub/mock/middleware_mock"
	"enigma.com/projectmanagementhub/mock/usecase_mock"
	"enigma.com/projectmanagementhub/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type userControllerTestSuite struct {
	suite.Suite
	UserUc         *usecase_mock.UserUseCaseMock
	authMiddleware *middleware_mock.AuthMiddlewareMock
	rg             *gin.RouterGroup
}

var ExpectedUser = model.User{
	Id:        "66213143-eeb9-427d-bc4c-c9aef4ef5528",
	Name:      "Admin1",
	Email:     "admin1@enigma",
	Password:  "admin1",
	Role:      "ADMIN",
	CreatedAt: time.Now(),
	UpdatedAt: time.Time{},
	DeletedAt: nil,
}

func (a *userControllerTestSuite) SetupTest() {

	a.UserUc = new(usecase_mock.UserUseCaseMock)
	a.authMiddleware = new(middleware_mock.AuthMiddlewareMock)
	r := gin.Default()
	rg := r.Group("/pmh-api/v1")
	rg.Use(a.authMiddleware.RequireToken("ADMIN", "TIM MEMBER", "MANAGER"))
	a.rg = rg
}

// Test Create User Controller Success
func (a *userControllerTestSuite) TestCreateUserController_Success() {
	a.UserUc.On("CreateUser", model.User{}).Return(ExpectedUser, nil)
	userController := NewUserController(a.rg, a.authMiddleware, a.UserUc)
	userController.Route()
	requestBody := `{"name":"Admin1","email":"admin1@enigma","password":"admin1","role":"ADMIN"}`
	request, err := http.NewRequest("POST", "/pmh-api/v1/create", strings.NewReader(requestBody))
	a.Nil(err)

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request
	ctx.Set("ADMIN", true)
	userController.CreateUser(ctx)
	a.Equal(200, record.Code)
}

// Test Create User Controller Failed
func (a *userControllerTestSuite) TestCreateUserController_Failed() {

	a.UserUc.On("CreateUser", model.User{}).Return(ExpectedUser, nil)
	userController := NewUserController(a.rg, a.authMiddleware, a.UserUc)
	userController.Route()
	requestBody := `{"name":"Admin1","email":"admin1@enigma","password":"admin1","role":"ADMIN"}`
	request, err := http.NewRequest("POST", "/pmh-api/v1/create", strings.NewReader(requestBody))
	a.Nil(err)

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request
	ctx.Set("ADMIN", false)
	userController.CreateUser(ctx)
	a.NotEqual(401, record.Code)
}

// Test Update User Success
func (a *userControllerTestSuite) TestUpdateUserController_Success() {
	a.UserUc.On("UpdateUser", model.User{}).Return(ExpectedUser, nil)
	userController := NewUserController(a.rg, a.authMiddleware, a.UserUc)
	userController.Route()
	requestBody := `{"name":"","email":"","password":"","role":""}`
	request, err := http.NewRequest("PUT", "/pmh-api/v1/update/:id", strings.NewReader(requestBody))
	a.Nil(err)

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request
	ctx.Set("ADMIN", true)
	userController.UpdateUser(ctx)
	a.Equal(200, record.Code)
}

// Test Update User Failed
func (a *userControllerTestSuite) TestUpdateUserController_Failed() {
	errorMessage := "update user failed"
	a.UserUc.On("UpdateUser", model.User{}).Return(model.User{}, errors.New(errorMessage))
	userController := NewUserController(a.rg, a.authMiddleware, a.UserUc)
	userController.Route()
	requestBody := `{"name":"","email":"","password":"","role":""}`
	request, err := http.NewRequest("PUT", "/pmh-api/v1/update/:id", strings.NewReader(requestBody))
	a.Nil(err)

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request
	ctx.Set("ADMIN", true)

	userController.UpdateUser(ctx)
	a.Equal(http.StatusInternalServerError, record.Code)
	responseBody, err := ioutil.ReadAll(record.Body)
	a.Nil(err)
	expectedErrorMessage := `{"code":500,"message":"update user failed"}`
	a.Equal(expectedErrorMessage, string(responseBody))
}

// Test Get User By Id Success
func (a *userControllerTestSuite) TestGetUserByIdlController_Success() {

	a.UserUc.On("FindUserById", "").Return(ExpectedUser, nil)

	userController := NewUserController(a.rg, a.authMiddleware, a.UserUc)
	userController.Route()
	request, err := http.NewRequest("GET", "/pmh-api/v1/user/:id?id=66213143-eeb9-427d-bc4c-c9aef4ef5528", nil)
	a.Nil(err)

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request
	ctx.Set("ADMIN", true)
	userController.FindUserById(ctx)
	a.Equal(200, record.Code)
}

// Test Get User By Id as Role Admin Failed
func (a *userControllerTestSuite) TestGetUserByIdController_Failed() {
	errorMessage := "get user by id failed"
	a.UserUc.On("FindUserById", "").Return(model.User{}, errors.New(errorMessage))
	userController := NewUserController(a.rg, a.authMiddleware, a.UserUc)
	userController.Route()
	request, err := http.NewRequest("GET", "/pmh-api/v1/user/:id", nil)
	a.Nil(err)

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request
	ctx.Set("ADMIN", false)
	userController.FindUserById(ctx)
	fmt.Println("Response Code:", record.Code)
	a.NotEqual(401, record.Code)
}

// Test User Get By Id as Role Admin User Not Found
func (a *userControllerTestSuite) TestGetUserByEmailController_Success() {

	a.UserUc.On("FindUserByEmail", "").Return(ExpectedUser, nil)

	userController := NewUserController(a.rg, a.authMiddleware, a.UserUc)
	userController.Route()
	request, err := http.NewRequest("GET", "/pmh-api/v1/user/email/email:?email=admin1@enigma", nil)
	a.Nil(err)

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request
	ctx.Set("ADMIN", true)
	userController.FindUserByEmail(ctx)
	a.Equal(200, record.Code)
}

func (a *userControllerTestSuite) TestGetUserByEmailController_Failed() {
	errorMessage := "get user by id failed"
	a.UserUc.On("FindUserByEmail", "").Return(model.User{}, errors.New(errorMessage))
	userController := NewUserController(a.rg, a.authMiddleware, a.UserUc)
	userController.Route()
	request, err := http.NewRequest("GET", "/pmh-api/v1/user/email/:email", nil)
	a.Nil(err)

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request
	ctx.Set("ADMIN", false)
	userController.FindUserByEmail(ctx)
	fmt.Println("Response Code:", record.Code)
	a.NotEqual(401, record.Code)
}

// Test Delete User Success

// Test Delete User Failed

// Test Delete User Failed

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(userControllerTestSuite))
}
