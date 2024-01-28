package usecase_mock

import (
	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/shared/shared_model"
	"github.com/stretchr/testify/mock"
)

type TaskUsecaseMock struct {
	mock.Mock
}

// UpdateTaskByManager implements repository.TaskRepository.
func (*TaskUsecaseMock) UpdateTaskByManager(payload model.Task) (model.Task, error) {
	panic("unimplemented")
}

// UpdateTaskByMember implements repository.TaskRepository.
func (*TaskUsecaseMock) UpdateTaskByMember(payload model.Task) (model.Task, error) {
	panic("unimplemented")
}

func (m *TaskUsecaseMock) GetAll(page int, size int) ([]model.Task, shared_model.Paging, error) {
	args := m.Called(page, size)
	return args.Get(0).([]model.Task), args.Get(1).(shared_model.Paging), args.Error(2)
}

func (m *TaskUsecaseMock) GetById(Id string) (model.Task, error) {
	args := m.Called(Id)
	return args.Get(0).(model.Task), args.Error(1)
}

func (m *TaskUsecaseMock) GetByPersonInCharge(Id string) ([]model.Task, error) {
	args := m.Called(Id)
	return args.Get(0).([]model.Task), args.Error(1)
}

func (m *TaskUsecaseMock) GetByProjectId(Id string) ([]model.Task, error) {
	args := m.Called(Id)
	return args.Get(0).([]model.Task), args.Error(1)
}

func (m *TaskUsecaseMock) CreateTask(payload model.Task) (model.Task, error) {
	args := m.Called(payload)
	return args.Get(0).(model.Task), args.Error(1)
}

func (m *TaskUsecaseMock) UpdateTask(userId string, payload model.Task) (model.Task, error) {
	args := m.Called(userId, payload)
	return args.Get(0).(model.Task), args.Error(1)
}

func (m *TaskUsecaseMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
