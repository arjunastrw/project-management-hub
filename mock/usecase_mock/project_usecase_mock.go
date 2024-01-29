package usecase_mock

import (
	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/shared/shared_model"
	"github.com/stretchr/testify/mock"
)

type ProjectUseCaseMock struct {
	mock.Mock
}

func (m *ProjectUseCaseMock) GetAll(page int, size int) ([]model.Project, shared_model.Paging, error) {
	args := m.Called(page, size)
	return args.Get(0).([]model.Project), args.Get(1).(shared_model.Paging), args.Error(2)
}

func (m *ProjectUseCaseMock) GetProjectById(id string) (model.Project, error) {
	args := m.Called(id)
	return args.Get(0).(model.Project), args.Error(1)
}

func (m *ProjectUseCaseMock) GetProjectsByDeadline(deadline string) ([]model.Project, error) {
	args := m.Called(deadline)
	return args.Get(0).([]model.Project), args.Error(1)
}

func (m *ProjectUseCaseMock) GetProjectsByManagerId(managerID string) ([]model.Project, error) {
	args := m.Called(managerID)
	return args.Get(0).([]model.Project), args.Error(1)
}

func (m *ProjectUseCaseMock) GetProjectsByMemberId(memberID string) ([]model.Project, error) {
	args := m.Called(memberID)
	return args.Get(0).([]model.Project), args.Error(1)
}

func (m *ProjectUseCaseMock) CreateNewProject(project model.Project) (model.Project, error) {
	args := m.Called(project)
	return args.Get(0).(model.Project), args.Error(1)
}

func (m *ProjectUseCaseMock) AddProjectMember(projectID string, members []string) error {
	args := m.Called(projectID, members)
	return args.Error(0)
}

func (m *ProjectUseCaseMock) DeleteProjectMember(projectID string, members []string) error {
	args := m.Called(projectID, members)
	return args.Error(0)
}

func (m *ProjectUseCaseMock) GetAllProjectMember(projectID string) ([]model.User, error) {
	args := m.Called(projectID)
	return args.Get(0).([]model.User), args.Error(1)
}

func (m *ProjectUseCaseMock) Update(project model.Project) (model.Project, error) {
	args := m.Called(project)
	return args.Get(0).(model.Project), args.Error(1)
}

func (m *ProjectUseCaseMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
