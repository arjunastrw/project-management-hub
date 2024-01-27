package repository_mock

import (
	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/shared/shared_model"
	"github.com/stretchr/testify/mock"
)

type TaskRepositoryMock struct {
	mock.Mock
}

func (m *TaskRepositoryMock) GetAll(page int, size int) ([]model.Task, shared_model.Paging, error) {
	args := m.Called(page, size)
	return args.Get(0).([]model.Task), args.Get(1).(shared_model.Paging), args.Error(2)
}

func (m *TaskRepositoryMock) GetById(Id string) (model.Task, error) {
	args := m.Called(Id)
	return args.Get(0).(model.Task), args.Error(1)
}

func (m *TaskRepositoryMock) GetByPersonInCharge(Id string) ([]model.Task, error) {
	args := m.Called(Id)
	return args.Get(0).([]model.Task), args.Error(1)
}

func (m *TaskRepositoryMock) GetByProjectId(Id string) ([]model.Task, error) {
	args := m.Called(Id)
	return args.Get(0).([]model.Task), args.Error(1)
}

func (m *TaskRepositoryMock) CreateTask(payload model.Task) (model.Task, error) {
	args := m.Called(payload)
	return args.Get(0).(model.Task), args.Error(1)
}

func (m *TaskRepositoryMock) UpdateTaskByManager(payload model.Task) (model.Task, error) {
	args := m.Called(payload)
	return args.Get(0).(model.Task), args.Error(1)
}

func (m *TaskRepositoryMock) UpdateTaskByMember(payload model.Task) (model.Task, error) {
	args := m.Called(payload)
	return args.Get(0).(model.Task), args.Error(1)
}

func (m *TaskRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
