package repository

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"time"

	"enigma.com/projectmanagementhub/model"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    UserRepository
}

func (a *UserRepositoryTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	a.mockDB, a.mockSql = db, mock
	a.repo = NewUserRepository(a.mockDB)
}

var userTest = model.User{
	Id:        "a1",
	Name:      "admin1",
	Email:     "admin1#example.com",
	Password:  "admin1",
	Role:      "ADMIN",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
	DeletedAt: nil,
	Project:   nil,
	Task:      nil,
}

// Test Get All User Success
// func (a *UserRepositoryTestSuite) TestGetAllUser_Success() {

// }

// Test Get By ID Success
func (a *UserRepositoryTestSuite) TestGeUsertById_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "role", "created_at", "updated_at"}).AddRow(userTest.Id, userTest.Name, userTest.Email, userTest.Password, userTest.Role, userTest.CreatedAt, userTest.UpdatedAt)
	a.mockSql.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE id = $1 AND deleted_at IS NULL")).WithArgs(userTest.Id).WillReturnRows(rows)
	actual, err := a.repo.GetById(userTest.Id)
	a.NoError(err)
	a.Nil(err)
	a.Equal(userTest, actual)
}

// Test Get By ID Failed
func (a *UserRepositoryTestSuite) TestGetUserById_Failed() {

	a.mockSql.ExpectQuery(regexp.QuoteMeta("SELECT * FROM users WHERE id = $1")).WithArgs(userTest.Id).WillReturnError(sql.ErrNoRows)
	_, err := a.repo.GetById(userTest.Id)
	a.Error(err)
}

// Test Get By Email Success
func (a *UserRepositoryTestSuite) TestGetUserByEmail_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "role", "created_at", "updated_at"}).AddRow(userTest.Id, userTest.Name, userTest.Email, userTest.Password, userTest.Role, userTest.CreatedAt, userTest.UpdatedAt)
	a.mockSql.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE email = $1 AND deleted_at IS NULL")).WithArgs(userTest.Email).WillReturnRows(rows)
	actual, err := a.repo.GetByEmail(userTest.Email)
	a.NoError(err)
	a.Nil(err)
	a.Equal(userTest, actual)
}

// Test Get By Email Not Found
func (a *UserRepositoryTestSuite) TestGetUserByEmail_UserNotFound() {
	rows := sqlmock.NewRows([]string{})
	a.mockSql.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE email = $1 AND deleted_at IS NULL")).
		WithArgs(userTest.Email).
		WillReturnRows(rows)

	actual, err := a.repo.GetByEmail(userTest.Email)

	a.Error(err)
	a.Equal("sql: no rows in result set", err.Error())

	expectedUser := model.User{}
	a.Equal(expectedUser, actual)
}

// Test Get By Email Failed
func (a *UserRepositoryTestSuite) TestGetUserByEmail_Failed() {

	a.mockSql.ExpectQuery(regexp.QuoteMeta("SELECT * FROM users WHERE email = $1")).WithArgs(userTest.Id).WillReturnError(sql.ErrNoRows)
	_, err := a.repo.GetById(userTest.Id)
	a.Error(err)
}

// Test Create Success
func (a *UserRepositoryTestSuite) TestCreateUser_Success() {

	a.mockSql.ExpectQuery(regexp.QuoteMeta("INSERT INTO users(name, email, password, role, updated_at) VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP) RETURNING id, name, email, password, role, created_at, updated_at")).WithArgs(userTest.Name, userTest.Email, userTest.Password, userTest.Role).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "role", "created_at", "updated_at"}).AddRow(userTest.Id, userTest.Name, userTest.Email, userTest.Password, userTest.Role, userTest.CreatedAt, userTest.UpdatedAt))
	actual, err := a.repo.CreateUser(userTest)
	a.NoError(err)
	a.Equal(userTest.Name, actual.Name)
}

// Test Create Failed
func (a *UserRepositoryTestSuite) TestCreateUser_Failed() {

	expectedError := errors.New("insert user failed")

	a.mockSql.ExpectQuery(regexp.QuoteMeta("INSERT INTO users(name, email, password, role, updated_at) VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP) RETURNING id, name, email, password, role, created_at, updated_at")).
		WithArgs(userTest.Name, userTest.Email, userTest.Password, userTest.Role).
		WillReturnError(expectedError)

	actual, err := a.repo.CreateUser(userTest)

	a.Error(err)
	a.Equal(expectedError, err)
	a.NotNil(actual)
}

// Test Update Success
// func (a *UserRepositoryTestSuite) TestUpdateUser_Success() {
// 	// Data pengguna yang akan di-update
// 	userToUpdate := model.User{
// 		Id:        "1",
// 		Name:      "New Name",
// 		Email:     "new.email@example.com",
// 		Password:  "new_password",
// 		Role:      "user",
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}

// 	// Expectation: SQL query update
// 	a.mockSql.ExpectQuery("UPDATE users SET name=\\$2, email=\\$3, password=\\$4, role=\\$5, updated_at=\\$6 WHERE id=\\$1 AND deleted_at IS NULL RETURNING id, name, email, password, role, created_at, updated_at").
// 		WithArgs(userToUpdate.Id, userToUpdate.Name, userToUpdate.Email, userToUpdate.Password, userToUpdate.Role, userToUpdate.UpdatedAt).
// 		WillReturnRows(sqlmock.NewResult(1, 1)) // 1 row affected

// 	// Memanggil metode UpdateUser
// 	actual, err := a.repo.Update(userToUpdate)

// 	// Verifikasi bahwa tidak ada error
// 	a.NoError(err)

// 	// Verifikasi bahwa hasil aktual sesuai dengan yang diharapkan
// 	a.Equal(userToUpdate, actual)

// 	// Verifikasi bahwa eksekusi ekspektasi query SQL sesuai
// 	err = a.mockSql.ExpectationsWereMet()
// 	a.NoError(err)
// }

// Test Update Failed
// func (a *UserRepositoryTestSuite) TestUpdateUser_Failed() {

// }

// Test Delete Success
func (a *UserRepositoryTestSuite) TestDeleteUser_Success() {
	a.mockSql.ExpectQuery(regexp.QuoteMeta("UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at IS NULL")).
		WithArgs(userTest.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "role", "created_at", "updated_at"}).AddRow(userTest.Id, userTest.Name, userTest.Email, userTest.Password, userTest.Role, userTest.CreatedAt, userTest.UpdatedAt)) // 1 row affected

	// Memanggil metode DeleteUser
	err := a.repo.Delete(userTest.Id)

	// Verifikasi bahwa tidak ada error
	a.NoError(err)

	// Verifikasi bahwa eksekusi ekspektasi query SQL sesuai
	err = a.mockSql.ExpectationsWereMet()
	a.NoError(err)
}

// Test Delete Failed
func (a *UserRepositoryTestSuite) TestDeleteUser_Failed() {
	// Expected error message
	expectedError := errors.New("delete user failed")

	a.mockSql.ExpectQuery(regexp.QuoteMeta("UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at IS NULL")).
		WithArgs(userTest.Id).
		WillReturnError(expectedError)

	err := a.repo.Delete(userTest.Id)

	a.Error(err)
	a.Equal(expectedError, err)

	err = a.mockSql.ExpectationsWereMet()
	a.NoError(err)
}

// Test Suite
func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
