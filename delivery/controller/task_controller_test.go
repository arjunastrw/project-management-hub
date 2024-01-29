package controller

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"enigma.com/projectmanagementhub/mock/middleware_mock"
	"enigma.com/projectmanagementhub/mock/usecase_mock"
	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/shared/shared_model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TaskControllerTestSuite struct {
	suite.Suite
	rg  *gin.RouterGroup
	tum *usecase_mock.TaskUsecaseMock
	amm *middleware_mock.AuthMiddlewareMock
}

func (s *TaskControllerTestSuite) SetupTest() {
	s.amm = new(middleware_mock.AuthMiddlewareMock)
	s.tum = new(usecase_mock.TaskUsecaseMock)

	router := gin.Default()
	gin.SetMode(gin.TestMode)
	rg := router.Group("/pmh-api/v1")
	rg.Use(s.amm.RequireToken("Admin", "MANAGER", "TEAM MEMBER"))
	s.rg = rg
}

func TestTaskControllerTestSuite(t *testing.T) {
	suite.Run(t, new(TaskControllerTestSuite))
}

func (s *TaskControllerTestSuite) TestCreateTask() {
	// Arrange
	taskController := NewTaskController(s.tum, s.amm, s.rg)
	s.tum.On("CreateTask", mock.Anything).Return(model.Task{Id: "1", Name: "Test Task"}, nil)

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/tasks/create", bytes.NewBufferString(`{"name": "Test Task"}`))
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("ADMIN", true)
	taskController.CreateTask(ctx)

	// Assert
	s.Equal(http.StatusOK, w.Code)
	s.Contains(w.Body.String(), "Test Task")
	s.tum.AssertExpectations(s.T())
}

func (s *TaskControllerTestSuite) TestCreateTask_Failure() {
	// Arrange
	s.tum.On("CreateTask", mock.Anything).Return(model.Task{}, fmt.Errorf("failed to create task"))
	taskController := NewTaskController(s.tum, s.amm, s.rg)

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/tasks/create", bytes.NewBufferString(`{"name": "Test Task"}`))
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("ADMIN", true)
	taskController.CreateTask(ctx)
	// Assert
	s.Equal(http.StatusInternalServerError, w.Code)
	s.Contains(w.Body.String(), "failed to create task")
	s.tum.AssertExpectations(s.T())
}

func (s *TaskControllerTestSuite) TestGetAllTasks_Success() {
	// Arrange
	expectedTasks := []model.Task{
		{
			Id:             "1",
			Name:           "Task name 1",
			Status:         "In Progress",
			Approval:       false,
			ApprovalDate:   nil,
			Feedback:       "Feedback",
			PersonInCharge: "1",
			ProjectId:      "1",
			Deadline:       "2024-05-05",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			DeletedAt:      nil,
		},
		{
			Id:             "2",
			Name:           "Task name 2",
			Status:         "In Progress",
			Approval:       false,
			ApprovalDate:   nil,
			Feedback:       "Feedback",
			PersonInCharge: "1",
			ProjectId:      "1",
			Deadline:       "2024-05-05",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			DeletedAt:      nil,
		},
	}
	s.tum.On("GetAll", 1, 10).Return(expectedTasks, shared_model.Paging{}, nil)
	taskController := NewTaskController(s.tum, s.amm, s.rg)

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/tasks/list?page=1&size=10", nil)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("ADMIN", true)
	taskController.GetAllTask(ctx)

	// Assert
	s.Equal(http.StatusOK, w.Code)
	s.Contains(w.Body.String(), "Task name 1")
	s.Contains(w.Body.String(), "Task name 2")
	s.tum.AssertExpectations(s.T())
}

func (s *TaskControllerTestSuite) TestGetAllTasks_Fail() {
	// Arrange
	taskController := NewTaskController(s.tum, s.amm, s.rg)
	s.tum.On("GetAll", 1, 10).Return([]model.Task{}, shared_model.Paging{}, fmt.Errorf("failed to get tasks"))
	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/tasks/list?page=1&size=10", nil)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("ADMIN", true)
	taskController.GetAllTask(ctx)
	// Assert
	s.Equal(http.StatusBadRequest, w.Code)
	s.Contains(w.Body.String(), "no task found")
	s.tum.AssertExpectations(s.T())
}

func (s *TaskControllerTestSuite) TestGetTaskById_Success() {
	// Arrange
	expectedTask := model.Task{
		Id:             "1",
		Name:           "Task name 1",
		Status:         "In Progress",
		Approval:       false,
		ApprovalDate:   nil,
		Feedback:       "Feedback",
		PersonInCharge: "1",
		ProjectId:      "1",
		Deadline:       "2024-05-05",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		DeletedAt:      nil,
	}
	s.tum.On("GetById", "1").Return(expectedTask, nil)
	taskController := NewTaskController(s.tum, s.amm, s.rg)
	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/tasks/getbyid/1", nil)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.AddParam("id", "1")
	ctx.Set("ADMIN", true)
	taskController.GetTaskById(ctx)
	// Assert
	s.Equal(http.StatusOK, w.Code)
	s.Contains(w.Body.String(), "Task name 1")
	s.tum.AssertExpectations(s.T())
}

func (s *TaskControllerTestSuite) TestGetTaskById_Fail() {
	// Arrange
	taskController := NewTaskController(s.tum, s.amm, s.rg)
	s.tum.On("GetById", "1").Return(model.Task{}, fmt.Errorf("not found"))
	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/tasks/getbyid/1", nil)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.AddParam("id", "1")
	ctx.Set("ADMIN", true)
	taskController.GetTaskById(ctx)
	// Assert
	s.Equal(http.StatusBadRequest, w.Code)
	s.Contains(w.Body.String(), "not found")
	s.tum.AssertExpectations(s.T())
}

func (s *TaskControllerTestSuite) TestGetTaskByPersonInCharge_Success() {
	// Arrange
	personInChargeID := "1"
	expectedTasks := []model.Task{
		{
			Id:             "1",
			Name:           "Task name 1",
			Status:         "In Progress",
			Approval:       false,
			ApprovalDate:   nil,
			Feedback:       "Feedback",
			PersonInCharge: personInChargeID,
			ProjectId:      "1",
			Deadline:       "2024-05-05",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			DeletedAt:      nil,
		},
		{
			Id:             "2",
			Name:           "Task name 2",
			Status:         "In Progress",
			Approval:       false,
			ApprovalDate:   nil,
			Feedback:       "Feedback",
			PersonInCharge: personInChargeID,
			ProjectId:      "1",
			Deadline:       "2024-05-05",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			DeletedAt:      nil,
		},
	}
	s.tum.On("GetByPersonInCharge", personInChargeID).Return(expectedTasks, nil)
	taskController := NewTaskController(s.tum, s.amm, s.rg)
	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/tasks/getbypic/1", nil)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.AddParam("id", personInChargeID)
	ctx.Set("ADMIN", true)
	taskController.GetTaskByPersonInCharge(ctx)
	// Assert
	s.Equal(http.StatusOK, w.Code)
	s.Contains(w.Body.String(), "Task name 1")
	s.Contains(w.Body.String(), "Task name 2")
	s.tum.AssertExpectations(s.T())
}

func (s *TaskControllerTestSuite) TestGetTaskByPersonInCharge_Fail() {
	// Arrange
	personInChargeID := "1"
	s.tum.On("GetByPersonInCharge", personInChargeID).Return([]model.Task{}, fmt.Errorf("not found"))
	taskController := NewTaskController(s.tum, s.amm, s.rg)
	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/tasks/getbypic/1", nil)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.AddParam("id", personInChargeID)
	ctx.Set("ADMIN", true)
	taskController.GetTaskByPersonInCharge(ctx)
	// Assert
	s.Equal(http.StatusBadRequest, w.Code)
	s.Contains(w.Body.String(), "not found")
	s.tum.AssertExpectations(s.T())
}

func (s *TaskControllerTestSuite) TestDeleteTask_Success() {
	// Arrange
	taskID := "1"

	s.tum.On("Delete", taskID).Return(nil)
	taskController := NewTaskController(s.tum, s.amm, s.rg)

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/tasks/delete/1", nil)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("ADMIN", true)
	ctx.AddParam("id", taskID)
	taskController.DeleteTask(ctx)

	// Assert
	s.Equal(http.StatusOK, w.Code)
	s.Contains(w.Body.String(), "Success")
	s.tum.AssertExpectations(s.T())
}

func (s *TaskControllerTestSuite) TestDeleteTask_Failure() {
	// Arrange
	taskID := "1"
	expectedError := errors.New("failed to delete task")

	s.tum.On("Delete", taskID).Return(expectedError)
	taskController := NewTaskController(s.tum, s.amm, s.rg)

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/tasks/delete/1", nil)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Set("ADMIN", true)
	ctx.AddParam("id", taskID)
	taskController.DeleteTask(ctx)

	// Assert
	s.Equal(http.StatusInternalServerError, w.Code)
	s.Contains(w.Body.String(), expectedError.Error())
	s.tum.AssertExpectations(s.T())
}

func (s *TaskControllerTestSuite) TestGetTaskByProjectId_Success() {
	// Arrange
	projectId := "1"
	expectedTasks := []model.Task{
		{
			Id:             "1",
			Name:           "Task name 1",
			Status:         "In Progress",
			Approval:       false,
			ApprovalDate:   nil,
			Feedback:       "Feedback",
			PersonInCharge: "1",
			ProjectId:      projectId,
			Deadline:       "2024-05-05",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			DeletedAt:      nil,
		},
		{
			Id:             "2",
			Name:           "Task name 2",
			Status:         "In Progress",
			Approval:       false,
			ApprovalDate:   nil,
			Feedback:       "Feedback",
			PersonInCharge: "1",
			ProjectId:      projectId,
			Deadline:       "2024-05-05",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			DeletedAt:      nil,
		},
	}
	s.tum.On("GetByProjectId", projectId).Return(expectedTasks, nil)
	taskController := NewTaskController(s.tum, s.amm, s.rg)
	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/tasks/getbyprojectid/1", nil)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.AddParam("id", projectId)
	ctx.Set("ADMIN", true)
	taskController.GetTaskByProjectId(ctx)
	// Assert
	s.Equal(http.StatusOK, w.Code)
	s.Contains(w.Body.String(), "Task name 1")
	s.Contains(w.Body.String(), "Task name 2")
	s.tum.AssertExpectations(s.T())
}

func (s *TaskControllerTestSuite) TestGetTaskByProjectId_Fail() {
	// Arrange
	projectId := "1"
	s.tum.On("GetByProjectId", projectId).Return([]model.Task{}, fmt.Errorf("not found"))
	taskController := NewTaskController(s.tum, s.amm, s.rg)
	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/tasks/getbyprojectid/1", nil)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.AddParam("id", projectId)
	ctx.Set("ADMIN", true)
	taskController.GetTaskByProjectId(ctx)
	// Assert
	s.Equal(http.StatusBadRequest, w.Code)
	s.Contains(w.Body.String(), "not found")
	s.tum.AssertExpectations(s.T())
}
