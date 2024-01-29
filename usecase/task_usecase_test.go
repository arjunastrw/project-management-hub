package usecase

import (
	"fmt"
	"testing"
	"time"

	"enigma.com/projectmanagementhub/mock/repository_mock"
	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/shared/shared_model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TaskUsecaseTest struct {
	suite.Suite
	trm *repository_mock.TaskRepositoryMock
	urm *repository_mock.UserRepositoryMock
	prm *repository_mock.ProjectRepositoryMock
	tc  TaskUsecase
}

func (t *TaskUsecaseTest) SetupTest() {
	t.trm = new(repository_mock.TaskRepositoryMock)
	t.urm = new(repository_mock.UserRepositoryMock)
	t.prm = new(repository_mock.ProjectRepositoryMock)
	t.tc = NewTaskUsecase(t.trm, t.urm, t.prm)
}

var expectedTask = model.Task{
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

var expectedTasks = []model.Task{
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

func TestTaskUsecase(t *testing.T) {
	suite.Run(t, new(TaskUsecaseTest))
}

func (t *TaskUsecaseTest) TestCreateTask_Success() {
	t.urm.On("GetById", expectedTask.PersonInCharge).Return(model.User{}, nil)
	t.prm.On("GetById", expectedTask.ProjectId).Return(model.Project{}, nil)
	t.trm.On("CreateTask", expectedTask).Return(expectedTask, nil)

	createdTask, err := t.tc.CreateTask(expectedTask)

	assert.NoError(t.T(), err)
	assert.Equal(t.T(), expectedTask, createdTask)

	t.urm.AssertExpectations(t.T())
	t.prm.AssertExpectations(t.T())
	t.trm.AssertExpectations(t.T())
}

func (t *TaskUsecaseTest) TestCreateTask_PersonInChargeInvalid() {
	t.urm.On("GetById", expectedTask.PersonInCharge).Return(model.User{}, fmt.Errorf("user not found"))

	_, err := t.tc.CreateTask(expectedTask)

	assert.Error(t.T(), err)
	assert.EqualError(t.T(), err, "failed to create task. person in charge id invalid")

	t.urm.AssertExpectations(t.T())
}

func (t *TaskUsecaseTest) TestCreateTask_ProjectIDInvalid() {
	t.urm.On("GetById", expectedTask.PersonInCharge).Return(model.User{}, nil)
	t.prm.On("GetById", expectedTask.ProjectId).Return(model.Project{}, fmt.Errorf("project not found"))

	_, err := t.tc.CreateTask(expectedTask)

	assert.Error(t.T(), err)
	assert.EqualError(t.T(), err, "failed to create task. project id invalid")

	t.urm.AssertExpectations(t.T())
	t.prm.AssertExpectations(t.T())
}

func (t *TaskUsecaseTest) TestCreateTask_EmptyFieldsExist() {
	taskWithEmptyFields := model.Task{
		Id:             "1",
		Name:           "", // Empty name field
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

	t.urm.On("GetById", taskWithEmptyFields.PersonInCharge).Return(model.User{}, nil)
	t.prm.On("GetById", taskWithEmptyFields.ProjectId).Return(model.Project{}, nil)

	_, err := t.tc.CreateTask(taskWithEmptyFields)

	assert.Error(t.T(), err)
	assert.EqualError(t.T(), err, "failed to create task. empty field exist")
}

func (t *TaskUsecaseTest) TestDeleteTask_Success() {
	t.trm.On("GetById", expectedTask.Id).Return(expectedTask, nil)
	t.trm.On("Delete", expectedTask.Id).Return(nil)

	err := t.tc.Delete(expectedTask.Id)

	assert.NoError(t.T(), err)

	t.trm.AssertExpectations(t.T())
}

func (t *TaskUsecaseTest) TestDeleteTask_TaskIDInvalid() {
	t.trm.On("GetById", expectedTask.Id).Return(model.Task{}, fmt.Errorf("task not found"))

	err := t.tc.Delete(expectedTask.Id)

	assert.Error(t.T(), err)
	assert.EqualError(t.T(), err, "failed to delete task. task id invalid")

	t.trm.AssertExpectations(t.T())
}

func (t *TaskUsecaseTest) TestGetAllTasks_Success() {
	page := 1
	size := 10
	expectedPaging := shared_model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   len(expectedTasks),
		TotalPages:  1, // Assuming all tasks fit in one page
	}

	t.trm.On("GetAll", page, size).Return(expectedTasks, expectedPaging, nil)

	tasks, paging, err := t.tc.GetAll(page, size)

	assert.NoError(t.T(), err)
	assert.Equal(t.T(), expectedTasks, tasks)
	assert.Equal(t.T(), expectedPaging, paging)

	t.trm.AssertExpectations(t.T())
}

func (t *TaskUsecaseTest) TestGetAllTasks_Failure() {
	page := 1
	size := 10

	t.trm.On("GetAll", page, size).Return([]model.Task{}, shared_model.Paging{}, fmt.Errorf("failed to get tasks"))

	_, _, err := t.tc.GetAll(page, size)

	assert.Error(t.T(), err)
	assert.EqualError(t.T(), err, "failed to get tasks")

	t.trm.AssertExpectations(t.T())
}

func (t *TaskUsecaseTest) TestGetTaskById_Success() {
	taskID := "1"
	expectedTask := model.Task{
		Id:             taskID,
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

	t.trm.On("GetById", taskID).Return(expectedTask, nil)

	task, err := t.tc.GetById(taskID)

	assert.NoError(t.T(), err)
	assert.Equal(t.T(), expectedTask, task)

	t.trm.AssertExpectations(t.T())
}

func (t *TaskUsecaseTest) TestGetTaskById_Failure() {
	taskID := "1"

	t.trm.On("GetById", taskID).Return(model.Task{}, fmt.Errorf("failed to get task by ID"))

	_, err := t.tc.GetById(taskID)

	assert.Error(t.T(), err)
	assert.EqualError(t.T(), err, "failed to get task by ID")

	t.trm.AssertExpectations(t.T())
}

func (t *TaskUsecaseTest) TestGetTaskByPersonInCharge_Success() {
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

	personInCharge := []model.User{
		{
			Id:   "1",
			Name: "User 1",
			Task: expectedTasks,
		},
		{
			Id:   "2",
			Name: "User 2",
		},
	}

	t.urm.On("GetById", personInChargeID).Return(personInCharge[0], nil)
	t.trm.On("GetByPersonInCharge", personInChargeID).Return(expectedTasks, nil)

	tasks, err := t.tc.GetByPersonInCharge(personInChargeID)

	assert.NoError(t.T(), err)
	assert.Equal(t.T(), expectedTasks, tasks)

	t.urm.AssertExpectations(t.T())
	t.trm.AssertExpectations(t.T())
}

func (t *TaskUsecaseTest) TestGetTaskByPersonInCharge_Failure() {
	personInChargeID := "1"

	t.urm.On("GetById", personInChargeID).Return(model.User{}, nil)

	_, err := t.tc.GetByPersonInCharge(personInChargeID)

	assert.Error(t.T(), err)
	assert.EqualError(t.T(), err, "this user currently has no tasks")

	t.urm.AssertExpectations(t.T())
	t.trm.AssertExpectations(t.T())
}

func (t *TaskUsecaseTest) TestGetTaskByPersonInCharge_InvalidPersonInCharge() {
	personInChargeID := "1"

	t.urm.On("GetById", personInChargeID).Return(model.User{}, fmt.Errorf("failed to get user by ID"))

	_, err := t.tc.GetByPersonInCharge(personInChargeID)

	assert.Error(t.T(), err)
	assert.EqualError(t.T(), err, "failed to get task by person in charge. person in charge id invalid")

	t.urm.AssertExpectations(t.T())

}

// continued from previous code snippet...

func (t *TaskUsecaseTest) TestGetTaskByProjectId_Success() {
	projectID := "1"
	expectedTasks := []model.Task{
		{
			Id:             "1",
			Name:           "Task name 1",
			Status:         "In Progress",
			Approval:       false,
			ApprovalDate:   nil,
			Feedback:       "Feedback",
			PersonInCharge: "1",
			ProjectId:      projectID,
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
			ProjectId:      projectID,
			Deadline:       "2024-05-05",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			DeletedAt:      nil,
		},
	}

	project := model.Project{
		Id:    "1",
		Name:  "Project 1",
		Tasks: expectedTasks,
	}

	t.prm.On("GetById", projectID).Return(project, nil)

	t.trm.On("GetByProjectId", projectID).Return(expectedTasks, nil)

	tasks, err := t.tc.GetByProjectId(projectID)

	assert.NoError(t.T(), err)
	assert.Equal(t.T(), expectedTasks, tasks)

	t.prm.AssertExpectations(t.T())
	t.trm.AssertExpectations(t.T())
}

func (t *TaskUsecaseTest) TestGetTaskByProjectId_Failure() {
	projectID := "1"

	t.prm.On("GetById", projectID).Return(model.Project{}, nil)

	_, err := t.tc.GetByProjectId(projectID)

	assert.Error(t.T(), err)
	assert.EqualError(t.T(), err, "this project currently has no tasks")

	t.prm.AssertExpectations(t.T())
	t.trm.AssertExpectations(t.T())
}

func (t *TaskUsecaseTest) TestGetTaskByProjectId_InvalidProjectID() {
	projectID := "1"

	t.prm.On("GetById", projectID).Return(model.Project{}, fmt.Errorf("failed to get task by project id. project id invalid"))

	_, err := t.tc.GetByProjectId(projectID)

	assert.Error(t.T(), err)
	assert.EqualError(t.T(), err, "failed to get task by project id. project id invalid")

	t.trm.AssertExpectations(t.T())
}

func (t *TaskUsecaseTest) TestUpdateTaskByManager_Success() {
	managerID := "1"
	taskPayload := model.Task{
		Id:             "3",
		Name:           "Updated Task",
		Status:         "In Progress",
		Approval:       false,
		ApprovalDate:   nil,
		Feedback:       "Updated Feedback",
		PersonInCharge: "2",
		ProjectId:      "1",
		Deadline:       "2024-07-07",
	}

	manager := model.User{
		Id:   managerID,
		Role: "MANAGER",
	}
	user := model.User{
		Id:   "2",
		Role: "TEAM MEMBER",
	}

	t.urm.On("GetById", managerID).Return(manager, nil)
	t.urm.On("GetById", taskPayload.PersonInCharge).Return(user, nil)
	t.trm.On("UpdateTaskByManager", taskPayload).Return(taskPayload, nil)

	updatedTask, err := t.tc.UpdateTask(managerID, taskPayload)

	assert.NoError(t.T(), err)
	assert.Equal(t.T(), taskPayload, updatedTask)

	t.urm.AssertExpectations(t.T())
	t.trm.AssertExpectations(t.T())

}

func (t *TaskUsecaseTest) TestUpdateTaskByTeamMember_Success() {
	taskPayload := model.Task{
		Id:             "3",
		Name:           "Updated Task",
		Status:         "In Progress",
		Approval:       false,
		ApprovalDate:   nil,
		Feedback:       "Updated Feedback",
		PersonInCharge: "2",
		ProjectId:      "1",
		Deadline:       "2024-07-07",
	}
	user := model.User{
		Id:   "2",
		Role: "TEAM MEMBER",
	}

	t.urm.On("GetById", user.Id).Return(user, nil)
	t.trm.On("GetById", taskPayload.Id).Return(taskPayload, nil)
	t.trm.On("UpdateTaskByMember", taskPayload).Return(taskPayload, nil)

	updatedTask, err := t.tc.UpdateTask(user.Id, taskPayload)

	assert.NoError(t.T(), err)
	assert.Equal(t.T(), taskPayload, updatedTask)

	t.urm.AssertExpectations(t.T())
	t.trm.AssertExpectations(t.T())

}

func (t *TaskUsecaseTest) TestDeleteTask_Failure() {
	taskID := "1"

	t.trm.On("GetById", taskID).Return(expectedTask, nil)
	t.trm.On("Delete", taskID).Return(fmt.Errorf("failed to delete task"))

	err := t.tc.Delete(taskID)

	assert.Error(t.T(), err)
	assert.EqualError(t.T(), err, "failed to delete task")

	t.trm.AssertExpectations(t.T())
}

func (t *TaskUsecaseTest) TestUpdateTaskByManager_TaskNotFound() {
	managerID := "1"
	taskPayload := model.Task{
		Id:             "3",
		Name:           "Updated Task",
		Status:         "In Progress",
		Approval:       false,
		ApprovalDate:   nil,
		Feedback:       "Updated Feedback",
		PersonInCharge: "1",
		ProjectId:      "1",
		Deadline:       "2024-07-07",
	}

	manager := model.User{
		Id:   managerID,
		Role: "MANAGER",
	}

	t.urm.On("GetById", managerID).Return(manager, nil)
	t.trm.On("UpdateTaskByManager", taskPayload).Return(model.Task{}, fmt.Errorf("task not found"))

	_, err := t.tc.UpdateTask(managerID, taskPayload)

	assert.Error(t.T(), err)
	assert.EqualError(t.T(), err, "task not found")

	t.urm.AssertExpectations(t.T())
	t.trm.AssertExpectations(t.T())
}

func (t *TaskUsecaseTest) TestUpdateTaskByTeamMember_NotPIC() {
	taskPayload := model.Task{
		Id:             "3",
		Name:           "Updated Task",
		Status:         "In Progress",
		Approval:       false,
		ApprovalDate:   nil,
		Feedback:       "Updated Feedback",
		PersonInCharge: "2",
		ProjectId:      "1",
		Deadline:       "2024-07-07",
	}
	user := model.User{
		Id:   "2",
		Role: "TEAM MEMBER",
	}

	t.urm.On("GetById", user.Id).Return(user, nil)
	t.trm.On("GetById", taskPayload.Id).Return(model.Task{}, fmt.Errorf("not pic"))

	_, err := t.tc.UpdateTask(user.Id, taskPayload)

	assert.Error(t.T(), err)
	assert.EqualError(t.T(), err, "only person in charge and project manager can update task")

	t.urm.AssertExpectations(t.T())
	t.trm.AssertExpectations(t.T())
}

func (t *TaskUsecaseTest) TestUpdateTaskByManager_InvalidManager() {
	managerID := "1"
	taskPayload := model.Task{
		Id:             "3",
		Name:           "Updated Task",
		Status:         "In Progress",
		Approval:       false,
		ApprovalDate:   nil,
		Feedback:       "Updated Feedback",
		PersonInCharge: "2",
		ProjectId:      "1",
		Deadline:       "2024-07-07",
	}

	t.urm.On("GetById", managerID).Return(model.User{}, fmt.Errorf("failed to get user by ID"))

	_, err := t.tc.UpdateTask(managerID, taskPayload)

	assert.Error(t.T(), err)
	assert.EqualError(t.T(), err, "failed to update task. user id invalid")

	t.urm.AssertExpectations(t.T())
}

func (t *TaskUsecaseTest) TestUpdateTaskByManager_InvalidPersonInCharge() {
	managerID := "1"
	taskPayload := model.Task{
		Id:             "3",
		Name:           "Updated Task",
		Status:         "In Progress",
		Approval:       false,
		ApprovalDate:   nil,
		Feedback:       "Updated Feedback",
		PersonInCharge: "2",
		ProjectId:      "1",
		Deadline:       "2024-07-07",
	}

	manager := model.User{
		Id:   managerID,
		Role: "MANAGER",
	}

	t.urm.On("GetById", managerID).Return(manager, nil)
	t.urm.On("GetById", taskPayload.PersonInCharge).Return(model.User{}, fmt.Errorf("failed to get user by ID"))

	_, err := t.tc.UpdateTask(managerID, taskPayload)

	assert.Error(t.T(), err)
	assert.EqualError(t.T(), err, "failed to update task. person in charge id invalid")

	t.urm.AssertExpectations(t.T())
}

func (t *TaskUsecaseTest) TestUpdateTaskByTeamMember_InvalidUser() {
	taskPayload := model.Task{
		Id:             "3",
		Name:           "Updated Task",
		Status:         "In Progress",
		Approval:       false,
		ApprovalDate:   nil,
		Feedback:       "Updated Feedback",
		PersonInCharge: "2",
		ProjectId:      "1",
		Deadline:       "2024-07-07",
	}
	user := model.User{
		Id:   "2",
		Role: "TEAM MEMBER",
	}

	t.urm.On("GetById", user.Id).Return(model.User{}, fmt.Errorf("failed to get user by ID"))

	_, err := t.tc.UpdateTask(user.Id, taskPayload)

	assert.Error(t.T(), err)
	assert.EqualError(t.T(), err, "failed to update task. user id invalid")

	t.urm.AssertExpectations(t.T())
}
