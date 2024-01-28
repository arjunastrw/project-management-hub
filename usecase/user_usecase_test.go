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
		Email:     "userremail2@mail.com",
		Password:  "password2",
		Role:      "USER",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

var expectedPaging = shared_model.Paging{Page: 1, RowsPerPage: 10, TotalPages: 2}

// Test Get All User Success
func (a *UserUseCaseTest) TestFindAllUser_Success() {

	a.urm.On("GetAll", 1, 10).Return(expectedUsers, expectedPaging, nil)
	users, paging, err := a.uc.FindAllUser(1, 10)

	assert.NoError(a.T(), err)
	assert.Equal(a.T(), expectedUsers, users)
	assert.Equal(a.T(), expectedPaging, paging)
	a.urm.AssertExpectations(a.T())
}

// Test Get All User Failed
func (a *UserUseCaseTest) TestFindAllUser_Failed() {

	a.urm.On("GetAll", 1, 10).Return([]model.User{}, shared_model.Paging{}, fmt.Errorf("failed to get users"))
	_, _, err := a.uc.FindAllUser(1, 10)
	a.Error(err)
}

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

	a.urm.On("GetByEmail", expectedUsers[1].Email).Return(expectedUsers[1], nil)
	actual, err := a.uc.FindUserByEmail(expectedUsers[1].Email)
	a.NoError(err)
	a.Equal(expectedUsers[1], actual)
}

// Test Get User By Email Not Found
func (a *UserUseCaseTest) TestFindUserByEmail_Failed() {

	a.urm.On("GetByEmail", expectedUsers[0].Email).Return(model.User{}, fmt.Errorf("user email not found"))
	_, err := a.uc.FindUserByEmail(expectedUsers[0].Email)
	a.Error(err)
}

// Test Create user Success

// Test Delete User Success
func (a *UserUseCaseTest) TestDeleteUser_Success() {

	a.urm.On("GetById", "1").Return(expectedUsers[0], nil)
	a.urm.On("Delete", "1").Return(nil)
	actual := a.uc.DeleteUser("1")
	a.NoError(actual)
	a.urm.AssertExpectations(a.T())
}

// Test Delete User Failed
func (a *UserUseCaseTest) TestDeleteUser_Failed() {
	a.urm.On("GetById", "1").Return(model.User{}, fmt.Errorf("user id not found"))
	a.urm.On("Delete", "1").Return(fmt.Errorf("failed to delete user"))
	err := a.uc.DeleteUser("1")
	a.Error(err)

}

func TestUserUsecase(t *testing.T) {

	suite.Run(t, new(UserUseCaseTest))
}
