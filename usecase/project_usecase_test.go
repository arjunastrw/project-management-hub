package usecase

import (
	"fmt"
	"testing"
	"time"

	"enigma.com/projectmanagementhub/mock/repository_mock"
	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/shared/shared_model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ProjectUsecaseTest struct {
	suite.Suite
	arm *repository_mock.ProjectRepositoryMock
	urm *repository_mock.UserRepositoryMock
	auc ProjectUseCase
}

func (s *ProjectUsecaseTest) SetupTest() {
	s.arm = new(repository_mock.ProjectRepositoryMock)
	s.urm = new(repository_mock.UserRepositoryMock)
	s.auc = NewProjectUseCase(s.arm, s.urm)
}

var projectTest = model.Project{
	Id:        "1",
	Name:      "project1",
	ManagerId: "managerid1",
	Deadline:  "2024-01-01",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
	DeletedAt: nil,
	Members:   nil,
}

// Get All Succes

func (s *ProjectUsecaseTest) TestGetAllProjectSuccess() {
	s.arm.On("GetAll", 1, 2).Return([]model.Project{projectTest}, shared_model.Paging{}, nil)
	// getall usecase
	actual, paging, err := s.auc.GetAll(1, 2)

	// Verif result
	s.NoError(err)
	s.Equal([]model.Project{projectTest}, actual)
	s.NotNil(paging)
	s.arm.AssertExpectations(s.T())
}

// Test Get All Project Fail
func (s *ProjectUsecaseTest) TestGetAllProjectFail() {
	s.arm.On("GetAll", 1, 10).Return([]model.Project{}, shared_model.Paging{}, fmt.Errorf("Query error"))
	// Call the use case method
	_, _, err := s.auc.GetAll(1, 10)
	s.Error(err)
}

// Test get project by Id Succes
func (s *ProjectUsecaseTest) TestGetProjectByIdSuccess() {
	s.arm.On("GetById", projectTest.Id).Return(projectTest, nil)
	actual, err := s.auc.GetProjectById(projectTest.Id)

	// Verify result
	s.NoError(err)
	s.Equal(projectTest, actual)

}

// Test get by Id Fail
func (s *ProjectUsecaseTest) TestGetByIdFail() {
	// Configure
	s.arm.On("GetById", "").Return(model.Project{}, fmt.Errorf("project id not found"))
	_, err := s.auc.GetProjectById("")
	s.Error(err)

}

// Test get project by deadline succes
func (s *ProjectUsecaseTest) TestGetProjectByDeadlineSuccess() {
	s.arm.On("GetByDeadline", projectTest.Deadline).Return([]model.Project{projectTest}, nil)
	actual, err := s.auc.GetProjectsByDeadline(projectTest.Deadline)

	// Verify result
	s.NoError(err)
	s.Equal([]model.Project{projectTest}, actual)

}

// test project get by deadline fail
func (s *ProjectUsecaseTest) TestGetByDeadlineFail() {
	s.arm.On("GetByDeadline", "").Return([]model.Project{projectTest}, fmt.Errorf(" not found"))
	_, err := s.auc.GetProjectsByDeadline("")
	s.Error(err)

}

// Test get project by manager id succes
func (s *ProjectUsecaseTest) TestGetProjectsByManagerIdSuccess() {
	s.urm.On("GetById", projectTest.ManagerId).Return(model.User{Role: "MANAGER"}, nil)
	s.arm.On("GetByManagerId", projectTest.ManagerId).Return([]model.Project{projectTest}, nil)
	actual, err := s.auc.GetProjectsByManagerId(projectTest.ManagerId)
	s.NoError(err)
	s.Equal([]model.Project{projectTest}, actual)
	s.urm.AssertExpectations(s.T())
	s.arm.AssertExpectations(s.T())
}

// Test get project by manager id fail
func (s *ProjectUsecaseTest) TestGetByManagerIdFail() {
	s.urm.On("GetById", projectTest.ManagerId).Return(model.User{Role: "MANAGER"}, nil)
	s.arm.On("GetByManagerId", projectTest.ManagerId).Return([]model.Project{projectTest}, fmt.Errorf("Error"))
	_, err := s.auc.GetProjectsByManagerId(projectTest.ManagerId)

	s.Error(err)
	s.arm.AssertExpectations(s.T())

}

// Test creat new project succes
func (s *ProjectUsecaseTest) TestCreateNewProjectSuccess() {
	s.arm.On("CreateProject", mock.AnythingOfType("model.Project")).Return(projectTest, nil)
	payload := model.Project{
		Name:      "Project One",
		ManagerId: "managerID",
		Deadline:  "2024-01-01",
	}
	createdProject, err := s.arm.CreateProject(payload)
	// Assertions
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), projectTest, createdProject)
	s.urm.AssertExpectations(s.T())
}

// Test Create New Fail
func (s *ProjectUsecaseTest) TestCreateNewProjectFail() {
	s.arm.On("CreateProject", mock.AnythingOfType("model.Project")).Return(model.Project{}, fmt.Errorf("Failed to create project"))
	payload := model.Project{
		Name:      "Project One",
		ManagerId: "managerID",
		Deadline:  "2024-01-01",
	}
	_, err := s.auc.CreateNewProject(payload)
	assert.Error(s.T(), err)

	s.arm.AssertExpectations(s.T())
}

// Test Add Project members succes
func (suite *ProjectUsecaseTest) TestAddProjectMemberSuccess() {
	suite.arm.On("AddProjectMember", projectTest.Id, mock.AnythingOfType("[]string")).Return(nil)

	// Call the use case method
	err := suite.arm.AddProjectMember(projectTest.Id, []string{"member1"})
	assert.NoError(suite.T(), err)
	suite.arm.AssertExpectations(suite.T())
}

// Test Add Project members Fail
func (s *ProjectUsecaseTest) TestAddProjectMemberError() {
	expectedError := fmt.Errorf("Error")
	// Configure the mock repository to return an error for AddProjectMember
	s.arm.On("AddProjectMember", projectTest.Id, mock.AnythingOfType("[]string")).Return(expectedError)
	err := s.arm.AddProjectMember(projectTest.Id, []string{"member1"})
	assert.Error(s.T(), err)
	assert.EqualError(s.T(), err, expectedError.Error())
	s.arm.AssertExpectations(s.T())
}

// Test Get all Project members succes
func (s *ProjectUsecaseTest) TestGetAllProjectMemberSuccess() {
	s.arm.On("GetById", projectTest.Id).Return(projectTest, nil)
	expectedMembers := []model.User{
		{Id: "member1", Name: "Member One"},
		{Id: "member2", Name: "Member Two"},
	}
	s.arm.On("GetAllProjectMember", projectTest.Id).Return(expectedMembers, nil)

	// Call the use case method
	members, err := s.auc.GetAllProjectMember(projectTest.Id)

	// Assertions
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expectedMembers, members)
	s.arm.AssertExpectations(s.T())
}

// Test Get all Project members fail
func (s *ProjectUsecaseTest) TestGetAllProjectMemberInvalidId() {
	s.arm.On("GetById", "").Return(model.Project{}, fmt.Errorf("Invalid project ID"))
	_, err := s.auc.GetAllProjectMember("")
	// Assertions
	assert.Error(s.T(), err)
	s.arm.AssertExpectations(s.T())
}

// Test delete succes
func (s *ProjectUsecaseTest) TestDeleteSuccess() {
	s.arm.On("Delete", projectTest.Id).Return(nil)
	err := s.auc.Delete(projectTest.Id)
	assert.NoError(s.T(), err)
	s.arm.AssertExpectations(s.T())
}

// Test delete fail
func (s *ProjectUsecaseTest) TestDeleteFail() {
	s.arm.On("Delete", projectTest.Id).Return(fmt.Errorf("Error deleting project"))
	err := s.auc.Delete(projectTest.Id)
	// Assertions
	s.Error(err)
}

// Test update succes
func (s *ProjectUsecaseTest) TestUpdateProjectSuccess() {
	// Mocking dependencies
	s.arm.On("GetById", projectTest.Id).Return(projectTest, nil)
	s.urm.On("GetById", projectTest.ManagerId).Return(model.User{Role: "MANAGER"}, nil)
	s.arm.On("Update", projectTest).Return(projectTest, nil)
	updatedProject, err := s.auc.Update(projectTest)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), projectTest, updatedProject)
	// Assert that the expected methods were called
	s.arm.AssertExpectations(s.T())
	s.urm.AssertExpectations(s.T())
}

// Test update fail
func (s *ProjectUsecaseTest) TestUpdateProjectfail() {
	s.arm.On("GetById", projectTest.Id).Return(projectTest, nil)
	s.urm.On("GetById", projectTest.ManagerId).Return(model.User{Role: "MANAGER"}, nil)

	expectedErrorMessage := "Failed updated"
	s.arm.On("Update", mock.AnythingOfType("model.Project")).Return(model.Project{}, fmt.Errorf(expectedErrorMessage))
	_, err := s.auc.Update(projectTest)
	assert.Error(s.T(), err)
	s.arm.AssertExpectations(s.T())
	s.urm.AssertExpectations(s.T())
}

func TestProjectUsecase(t *testing.T) {
	suite.Run(t, new(ProjectUsecaseTest))
}
