package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
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
	Members:   nil,
	Tasks:     nil,
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
	a.ProjectUc.On("GetAll").Return([]model.Project{ExpectedProject}, shared_model.Paging{}, nil)
	projectController := NewProjectController(a.ProjectUc, a.authMiddleware, a.rg)
	projectController.Route()
	requestBody := `{"name":"Project Web","manager_id":"95a85ac6-a999-4039-ba4f-832ca6f6ed48","deadline":"2024-02-05"}`
	request, err := http.NewRequest("GET", "/pmh-api/v1/project?page=1&size=10", strings.NewReader(requestBody))
	a.Nil(err)

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request
	ctx.Set("ADMIN", true)
	projectController.GetAll(ctx)
	a.Equal(200, record.Code)
}

func TestProjectControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ProjectControllerTestSuite))
}
