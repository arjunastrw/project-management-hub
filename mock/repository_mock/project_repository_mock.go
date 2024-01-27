package repository_mock

import (
	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/shared/shared_model"
	"github.com/stretchr/testify/mock"
)

type ProjectRepositoryMock struct {
	mock.Mock
}

func (m *ProjectRepositoryMock) GetAll(page int, size int) ([]model.Project, shared_model.Paging, error) {
	args := m.Called(page, size)
	return args.Get(0).([]model.Project), args.Get(1).(shared_model.Paging), args.Error(2)
}

func (m *ProjectRepositoryMock) GetById(id string) (model.Project, error) {
	args := m.Called(id)
	return args.Get(0).(model.Project), args.Error(1)
}

func (m *ProjectRepositoryMock) GetByDeadline(date string) ([]model.Project, error) {
	args := m.Called(date)
	return args.Get(0).([]model.Project), args.Error(1)
}

func (m *ProjectRepositoryMock) GetByManagerId(id string) ([]model.Project, error) {
	args := m.Called(id)
	return args.Get(0).([]model.Project), args.Error(1)
}

func (m *ProjectRepositoryMock) GetByMemberId(id string) ([]model.Project, error) {
	args := m.Called(id)
	return args.Get(0).([]model.Project), args.Error(1)
}

func (m *ProjectRepositoryMock) CreateProject(payload model.Project) (model.Project, error) {
	args := m.Called(payload)
	return args.Get(0).(model.Project), args.Error(1)
}

func (m *ProjectRepositoryMock) AddProjectMember(id string, members []string) error {
	args := m.Called(id, members)
	return args.Error(0)
}

func (m *ProjectRepositoryMock) DeleteProjectMember(id string, members []string) error {
	args := m.Called(id, members)
	return args.Error(0)
}

func (m *ProjectRepositoryMock) GetAllProjectMember(id string) ([]model.User, error) {
	args := m.Called(id)
	return args.Get(0).([]model.User), args.Error(1)
}

func (m *ProjectRepositoryMock) Update(payload model.Project) (model.Project, error) {
	args := m.Called(payload)
	return args.Get(0).(model.Project), args.Error(1)
}

func (m *ProjectRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
