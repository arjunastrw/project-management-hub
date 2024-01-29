package controller

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"enigma.com/projectmanagementhub/mock/middleware_mock"
	"enigma.com/projectmanagementhub/mock/usecase_mock"
	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/shared/shared_model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type ProjectControllerTestSuite struct {
	suite.Suite
	ProjectUc      *usecase_mock.ProjectUseCaseMock
	authMiddleware *middleware_mock.AuthMiddlewareMock
	rg             *gin.RouterGroup
}

var ExpectedProject = model.Project{
	Id:        "65fb33dc-a9f8-4a7c-9f07-ab2210b0d535",
	Name:      "Project Web",
	ManagerId: "95a85ac6-a999-4039-ba4f-832ca6f6ed48",
	Deadline:  "2024-02-05",
	CreatedAt: time.Now(),
	UpdatedAt: time.Time{},
	DeletedAt: nil,
	Members:   []model.User{},
	Tasks:     []model.Task{},
}

var ExpectedUserMember = model.User{
	Id:       "95a85ac6-a999-4039-ba4f-832ca6f6ed48",
	Name:     "member",
	Email:    "member@mail",
	Password: "",
	Role:     "TEAM MEMBER",
}

func (a *ProjectControllerTestSuite) SetupTest() {

	a.ProjectUc = new(usecase_mock.ProjectUseCaseMock)
	a.authMiddleware = new(middleware_mock.AuthMiddlewareMock)
	r := gin.Default()
	rg := r.Group("/pmh-api/v1")
	rg.Use(a.authMiddleware.RequireToken("ADMIN", "TIM MEMBER", "MANAGER"))
	a.rg = rg
}

// Test Get All Project Success
func (a *ProjectControllerTestSuite) TestGetAllProjectController_Success() {
	a.ProjectUc.On("GetAll", 1, 10).Return([]model.Project{ExpectedProject}, shared_model.Paging{}, nil)
	projectController := NewProjectController(a.ProjectUc, a.authMiddleware, a.rg)

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/project/list?page=1&size=10", nil)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("ADMIN", true)
	projectController.GetAll(ctx)

	// Assert
	a.Equal(200, w.Code)
	a.ProjectUc.AssertExpectations(a.T())
}

// Test Get All Project Failed
func (a *ProjectControllerTestSuite) TestGetAllProjectController_Failed() {
	expectedError := errors.New("error occurred while getting projects")
	a.ProjectUc.On("GetAll", 1, 10).Return([]model.Project{}, shared_model.Paging{}, expectedError)
	projectController := NewProjectController(a.ProjectUc, a.authMiddleware, a.rg)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/project/list?page=1&size=10", nil)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("ADMIN", true)
	projectController.GetAll(ctx)
	a.Equal(500, w.Code)
	a.ProjectUc.AssertExpectations(a.T())
}

// Test Project Get By Id Success
func (a *ProjectControllerTestSuite) TestGetProjectByIdController_Success() {

	a.ProjectUc.On("GetProjectById", "").Return(ExpectedProject, nil)
	projectController := NewProjectController(a.ProjectUc, a.authMiddleware, a.rg)
	projectController.Route()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/project/:id?id=65fb33dc-a9f8-4a7c-9f07-ab2210b0d535", nil)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("ADMIN", true)
	projectController.GetProjectById(ctx)
	a.Equal(200, w.Code)
	a.ProjectUc.AssertExpectations(a.T())

}

// Test Get By Id Failed
func (a *ProjectControllerTestSuite) TestGetProjectByIdController_Failed() {

	a.ProjectUc.On("GetProjectById", "").Return(model.Project{}, errors.New("error"))
	projectController := NewProjectController(a.ProjectUc, a.authMiddleware, a.rg)
	projectController.Route()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/project/:id?id=65fb33dc-a9f8-4a7c-9f07-ab2210b0d535", nil)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("ADMIN", true)
	projectController.GetProjectById(ctx)
	a.Equal(400, w.Code)
	a.ProjectUc.AssertExpectations(a.T())
}

// Test Get Project By Deadline
func (a *ProjectControllerTestSuite) TestGetProjectByDeadlineController_Success() {

	a.ProjectUc.On("GetProjectsByDeadline", "").Return([]model.Project{ExpectedProject}, nil)
	projectController := NewProjectController(a.ProjectUc, a.authMiddleware, a.rg)
	projectController.Route()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/project/deadline/:deadline?deadline=2024-02-05", nil)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("ADMIN", true)
	projectController.GetProjectsByDeadline(ctx)
	a.Equal(200, w.Code)
	a.ProjectUc.AssertExpectations(a.T())
}

// Test Get Project By Deadline Failed
func (a *ProjectControllerTestSuite) TestGetProjectByDeadlineController_Failed() {

	a.ProjectUc.On("GetProjectsByDeadline", "").Return([]model.Project{}, errors.New("error"))
	projectController := NewProjectController(a.ProjectUc, a.authMiddleware, a.rg)
	projectController.Route()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/project/deadline/:deadline?deadline=2024-02-05", nil)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("ADMIN", true)
	projectController.GetProjectsByDeadline(ctx)
	a.Equal(500, w.Code)
	a.ProjectUc.AssertExpectations(a.T())
}

// Test Get Project By Id Manager Success
func (a *ProjectControllerTestSuite) TestGetProjectByManagerIdController_Success() {

	a.ProjectUc.On("GetProjectsByManagerId", "").Return([]model.Project{ExpectedProject}, nil)
	projectController := NewProjectController(a.ProjectUc, a.authMiddleware, a.rg)
	projectController.Route()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/project/manager/:id?id=65fb33dc-a9f8-4a7c-9f07-ab2210b0d535", nil)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("ADMIN", true)
	projectController.GetProjectsByManagerId(ctx)
	a.Equal(200, w.Code)
	a.ProjectUc.AssertExpectations(a.T())
}

// Test Get Project By Id Manager Failed
func (a *ProjectControllerTestSuite) TestGetProjectByManagerIdController_Failed() {

	a.ProjectUc.On("GetProjectsByManagerId", "").Return([]model.Project{}, errors.New("error"))
	projectController := NewProjectController(a.ProjectUc, a.authMiddleware, a.rg)
	projectController.Route()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/project/manager/:id?id=65fb33dc-a9f8-4a7c-9f07-ab2210b0d535", nil)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("ADMIN", true)
	projectController.GetProjectsByManagerId(ctx)
	a.Equal(500, w.Code)
	a.ProjectUc.AssertExpectations(a.T())
}

// Test Get Project By Member Id Success
func (a *ProjectControllerTestSuite) TestGetProjectByMemberIdController_Success() {

	a.ProjectUc.On("GetProjectsByMemberId", "").Return([]model.Project{ExpectedProject}, nil)
	projectController := NewProjectController(a.ProjectUc, a.authMiddleware, a.rg)
	projectController.Route()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/project/member/:id?id=65fb33dc-a9f8-4a7c-9f07-ab2210b0d535", nil)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("ADMIN", true)
	projectController.GetProjectsByMemberId(ctx)
	a.Equal(200, w.Code)
	a.ProjectUc.AssertExpectations(a.T())
}

// Test Get Project By Member Id Failed
func (a *ProjectControllerTestSuite) TestGetProjectByMemberIdController_Failed() {

	a.ProjectUc.On("GetProjectsByMemberId", "").Return([]model.Project{}, errors.New("error"))
	projectController := NewProjectController(a.ProjectUc, a.authMiddleware, a.rg)
	projectController.Route()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/project/member/:id?id=65fb33dc-a9f8-4a7c-9f07-ab2210b0d535", nil)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("ADMIN", true)
	projectController.GetProjectsByMemberId(ctx)
	a.Equal(500, w.Code)
	a.ProjectUc.AssertExpectations(a.T())
}

// Test Create Project Success
func (a *ProjectControllerTestSuite) TestCreateProjectController_Success() {

	a.ProjectUc.On("CreateNewProject", model.Project{}).Return(ExpectedProject, nil)
	projectController := NewProjectController(a.ProjectUc, a.authMiddleware, a.rg)
	projectController.Route()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/project/create", bytes.NewBufferString(`{}`))
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("ADMIN", true)
	projectController.CreateNewProject(ctx)
	a.Equal(200, w.Code)
	a.ProjectUc.AssertExpectations(a.T())
}

// Test Create Project Failed
func (a *ProjectControllerTestSuite) TestCreateProjectController_Failed() {

	a.ProjectUc.On("CreateNewProject", model.Project{}).Return(model.Project{}, errors.New("error"))
	projectController := NewProjectController(a.ProjectUc, a.authMiddleware, a.rg)
	projectController.Route()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/project/create", bytes.NewBufferString(`{}`))
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("ADMIN", true)
	projectController.CreateNewProject(ctx)
	a.Equal(500, w.Code)
	a.ProjectUc.AssertExpectations(a.T())
}

// Test Get All Project Member

func TestProjectControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ProjectControllerTestSuite))
}
