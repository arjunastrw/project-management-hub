package repository

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"time"

	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/shared/shared_model"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
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
	Id:        "1",
	Name:      "admin1",
	Email:     "admin1#example.com",
	Password:  "",
	Role:      "ADMIN",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
	DeletedAt: nil,
	Project:   nil,
	Task:      nil,
}

var userTestUpdate = model.User{
	Id:        "1",
	Name:      "admin1",
	Email:     "adminganteng1@example.com",
	Password:  "",
	Role:      "ADMIN",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
	DeletedAt: nil,
	Project:   nil,
	Task:      nil,
}

// Test Get All User Success
func (a *UserRepositoryTestSuite) TestUserRepository_GetAll_Success() {
	// Mock the SQL query expectations for GetAll.
	rows := sqlmock.NewRows([]string{"id", "name", "email", "role", "created_at", "updated_at"}).
		AddRow(userTest.Id, userTest.Name, userTest.Email, userTest.Role, userTest.CreatedAt, userTest.UpdatedAt)
	a.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, email, role, created_at, updated_at FROM users WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2`)).
		WithArgs(10, 0).
		WillReturnRows(rows)
	a.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM users WHERE deleted_at IS NULL`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	// Call the GetAll method.
	resultUsers, paging, err := a.repo.GetAll(1, 10)

	// Assertions
	assert.NoError(a.T(), err)
	assert.Len(a.T(), resultUsers, 1)
	assert.Equal(a.T(), userTest, resultUsers[0])
	assert.Equal(a.T(), 1, paging.TotalRows)
}

// Test Get All User Failed
func (a *UserRepositoryTestSuite) TestUserRepository_GetAll_Failed() {
	a.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, email, role, created_at, updated_at FROM users WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2`)).
		WithArgs(10, 0).
		WillReturnError(sql.ErrConnDone)

	resultUsers, paging, err := a.repo.GetAll(1, 10)

	// Assertions
	assert.Error(a.T(), err)
	assert.True(a.T(), errors.Is(err, sql.ErrConnDone))
	assert.Empty(a.T(), resultUsers)
	assert.Equal(a.T(), shared_model.Paging{}, paging)
}

// Test Get All User Error Row Scan
func (a *UserRepositoryTestSuite) TestUserRepository_GetAll_ErrorOnRowScan() {
	// Mock the SQL query expectations for GetAll with an error on row scan.
	rows := sqlmock.NewRows([]string{"id", "name", "email", "role", "created_at", "updated_at"}).
		AddRow("invalid_id", userTest.Name, userTest.Email, userTest.Role, userTest.CreatedAt, userTest.UpdatedAt)
	a.mockSql.ExpectQuery(`ELECT id, name, email, role, created_at, updated_at FROM users WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2`).
		WithArgs(10, 0).
		WillReturnRows(rows)

	// Call the GetAll method.
	resultUsers, paging, err := a.repo.GetAll(1, 10)

	// Assertions
	assert.Error(a.T(), err)
	assert.NotEmpty(a.T(), err.Error())
	assert.Empty(a.T(), resultUsers)
	assert.Equal(a.T(), shared_model.Paging{}, paging)
}

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

// Test Update User Success
func (a *UserRepositoryTestSuite) TestUpdateUser_Success() {
	// Mock the SQL query expectations for UpdateTaskByManager.
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "role", "created_at", "updated_at"}).
		AddRow(userTestUpdate.Id, userTestUpdate.Name, userTestUpdate.Email, userTestUpdate.Password, userTestUpdate.Role, userTestUpdate.CreatedAt, userTestUpdate.UpdatedAt)
	a.mockSql.ExpectQuery(regexp.QuoteMeta(`UPDATE users SET name = $2, email = $3, password = $4, role = $5, updated_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at IS NULL RETURNING id, name, email, password, role, created_at, updated_at`)).
		WithArgs(userTestUpdate.Id, userTestUpdate.Name, userTestUpdate.Email, userTestUpdate.Password, userTestUpdate.Role).
		WillReturnRows(rows)

	updatedUser, err := a.repo.Update(userTestUpdate)
	assert.NoError(a.T(), err)
	assert.Equal(a.T(), userTestUpdate, updatedUser)
}

// Test Update Failed
func (a *UserRepositoryTestSuite) TestUpdateUser_Failed() {
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "role", "created_at", "updated_at"}).
		AddRow(userTestUpdate.Id, "differentName", userTestUpdate.Email, userTestUpdate.Password, userTestUpdate.Role, userTestUpdate.CreatedAt, userTestUpdate.UpdatedAt)
	a.mockSql.ExpectQuery(regexp.QuoteMeta(`UPDATE users SET name = $2, email = $3, password = $4, role = $5, updated_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at IS NULL RETURNING id, name, email, password, role, created_at, updated_at`)).
		WithArgs(userTestUpdate.Id, userTestUpdate.Name, userTestUpdate.Email, userTestUpdate.Password, userTestUpdate.Role).
		WillReturnRows(rows)

	updatedUser, err := a.repo.Update(userTestUpdate)
	assert.NoError(a.T(), err)
	assert.NotEqual(a.T(), userTestUpdate, updatedUser)

}

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

// Test Delete Failedcl
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
