package usecasemock

import (
	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/shared/shared_model"
	"github.com/stretchr/testify/mock"
)

type UserUseCaseMock struct {
	mock.Mock
}

func (a *UserUseCaseMock) FindAllUser(page int, size int) ([]model.User, shared_model.Paging, error) {
	args := a.Called(page, size)
	return args.Get(0).([]model.User), args.Get(1).(shared_model.Paging), args.Error(2)
}

func (a *UserUseCaseMock) FindUserById(id string) (model.User, error) {
	args := a.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}

func (a *UserUseCaseMock) FindUserByEmail(email string) (model.User, error) {
	args := a.Called(email)
	return args.Get(0).(model.User), args.Error(1)
}

func (a *UserUseCaseMock) CreateUser(payload model.User) (model.User, error) {
	args := a.Called(payload)
	return args.Get(0).(model.User), args.Error(1)
}

func (a *UserUseCaseMock) UpdateUser(payload model.User) (model.User, error) {
	args := a.Called(payload)
	return args.Get(0).(model.User), args.Error(1)
}

func (a *UserUseCaseMock) DeleteUser(id string) error {
	args := a.Called(id)
	return args.Error(0)
}
