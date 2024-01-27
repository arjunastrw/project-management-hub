package usecase

import (
	"fmt"
	"testing"
	"time"

	"enigma.com/projectmanagementhub/mock/repository_mock"
	"enigma.com/projectmanagementhub/model"
	"github.com/stretchr/testify/suite"
)

type UserUseCaseTest struct {
	suite.Suite
	urm *repository_mock.UserRepositoryMock
	uc  UserUseCase
}

func (a *UserUseCaseTest) SetupTest() {
	a.urm = new(repository_mock.UserRepositoryMock)
	a.uc = NewUserUseCase(a.urm)
}

var expectedUsers = []model.User{
	{
		Id:        "1",
		Name:      "User name 1",
		Email:     "useremail1@mail.com",
		Password:  "password1",
		Role:      "ADMIN",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Id:        "2",
		Name:      "User name 2",
		Email:     "authoremail2@mail.com",
		Password:  "password2",
		Role:      "USER",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

// Test Get All User Failed

// Test Get User By ID Success
func (a *UserUseCaseTest) TestFindUserById_Success() {

	a.urm.On("GetById", expectedUsers[0].Id).Return(expectedUsers[0], nil)
	actual, err := a.uc.FindUserById(expectedUsers[0].Id)
	a.NoError(err)
	a.Equal(expectedUsers[0], actual)
}

// Test Get User By ID Failed
func (a *UserUseCaseTest) TestFindUserById_Failed() {

	a.urm.On("GetById", expectedUsers[0].Id).Return(model.User{}, fmt.Errorf("user id not found"))
	_, err := a.uc.FindUserById("1")
	a.Error(err)
}

// Test Get User By Email Success
func (a *UserUseCaseTest) TestFindUserByEmail_Success() {

	a.urm.On("GetByEmail", expectedUsers[0].Email).Return(expectedUsers[1], nil)
	actual, err := a.uc.FindUserByEmail(expectedUsers[0].Email)
	a.NoError(err)
	a.Equal(expectedUsers[1], actual)
}

// Test Get User By Email Failed
func (a *UserUseCaseTest) TestFindUserByEmail_Failed() {

	a.urm.On("GetByEmail", expectedUsers[0].Email).Return(model.User{}, fmt.Errorf("user email not found"))
	_, err := a.uc.FindUserByEmail(expectedUsers[0].Email)
	a.Error(err)
}

// Test Create User Success

func TestUserUsecase(t *testing.T) {

	suite.Run(t, new(UserUseCaseTest))
}
