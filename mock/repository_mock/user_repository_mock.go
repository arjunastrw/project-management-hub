package repository_mock

import (
	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/shared/shared_model"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) GetAll(page int, size int) ([]model.User, shared_model.Paging, error) {
	args := m.Called(page, size)
	return args.Get(0).([]model.User), args.Get(1).(shared_model.Paging), args.Error(2)
}

func (m *UserRepositoryMock) GetById(id string) (model.User, error) {
	args := m.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *UserRepositoryMock) GetByEmail(email string) (model.User, error) {
	args := m.Called(email)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *UserRepositoryMock) CreateUser(payload model.User) (model.User, error) {
	args := m.Called(payload)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *UserRepositoryMock) Update(payload model.User) (model.User, error) {
	args := m.Called(payload)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *UserRepositoryMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
