package repository

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"enigma.com/projectmanagementhub/model"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ProjectRepositoryTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    ProjectRepository
}

func (t *ProjectRepositoryTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	t.mockDB, t.mockSql = db, mock
	t.repo = NewProjectRepository(t.mockDB)
}

var projectTest = model.Project{
	Id:        "1",
	Name:      "project1",
	ManagerId: "1",
	Deadline:  "2024-01-01",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
	DeletedAt: nil,
}

var updatedProjectTest = model.Project{
	Id:        "1",
	Name:      "project1",
	ManagerId: "1",
	Deadline:  "2024-01-01",
	CreatedAt: projectTest.CreatedAt,
	UpdatedAt: time.Now(),
	DeletedAt: nil,
}

func (t *ProjectRepositoryTestSuite) TestProjectRepository_CreateProject_Success() {
	// Mock the SQL query expectations for CreateProject with a success outcome.
	t.mockSql.ExpectQuery(`INSERT INTO projects\(name, manager_id, deadline, updated_at\) VALUES \(\$1, \$2, \$3, CURRENT_TIMESTAMP\) RETURNING id, name, manager_id, deadline, created_at, updated_at`).
		WithArgs(projectTest.Name, projectTest.ManagerId, projectTest.Deadline).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "manager_id", "deadline", "created_at", "updated_at"}).
			AddRow(projectTest.Id, projectTest.Name, projectTest.ManagerId, projectTest.Deadline, projectTest.CreatedAt, projectTest.UpdatedAt))

	// Call the CreateProject method.
	resultProject, err := t.repo.CreateProject(projectTest)

	// Assertions
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), projectTest, resultProject)
}

func (t *ProjectRepositoryTestSuite) TestProjectRepository_CreateProject_ErrorOnQuery() {
	// Mock the SQL query expectations for CreateProject with an error.
	t.mockSql.ExpectQuery(`INSERT INTO projects\(name, manager_id, deadline, updated_at\) VALUES \(\$1, \$2, \$3, CURRENT_TIMESTAMP\) RETURNING id, name, manager_id, deadline, created_at, updated_at`).
		WithArgs(projectTest.Name, projectTest.ManagerId, projectTest.Deadline).
		WillReturnError(sql.ErrConnDone)

	// Call the CreateProject method.
	resultProject, err := t.repo.CreateProject(projectTest)

	// Assertions
	assert.Error(t.T(), err)
	assert.True(t.T(), errors.Is(err, sql.ErrConnDone))
	assert.Equal(t.T(), model.Project{}, resultProject)
}

// UpdateProject method testing methods
func (t *ProjectRepositoryTestSuite) TestProjectRepository_UpdateProject_Success() {
	// Mock the SQL query expectations for UpdateProject with a success outcome.
	t.mockSql.ExpectQuery(`UPDATE projects SET name = \$2, manager_id = \$3, deadline = \$4, updated_at = CURRENT_TIMESTAMP WHERE id = \$1 AND deleted_at IS NULL RETURNING id, name, manager_id, deadline, created_at, updated_at`).
		WithArgs(updatedProjectTest.Id, updatedProjectTest.Name, updatedProjectTest.ManagerId, updatedProjectTest.Deadline).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "manager_id", "deadline", "created_at", "updated_at"}).
			AddRow(updatedProjectTest.Id, updatedProjectTest.Name, updatedProjectTest.ManagerId, updatedProjectTest.Deadline, updatedProjectTest.CreatedAt, updatedProjectTest.UpdatedAt))

	// Call the UpdateProject method.
	resultProject, err := t.repo.Update(updatedProjectTest)

	// Assertions
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), updatedProjectTest, resultProject)
}

func (t *ProjectRepositoryTestSuite) TestProjectRepository_UpdateProject_ErrorOnQuery() {
	// Mock the SQL query expectations for UpdateProject with an error.
	t.mockSql.ExpectQuery(`UPDATE projects SET name = \$2, manager_id = \$3, deadline = \$4, updated_at = CURRENT_TIMESTAMP WHERE id = \$1 AND deleted_at IS NULL RETURNING id, name, manager_id, deadline, created_at, updated_at`).
		WithArgs(updatedProjectTest.Id, updatedProjectTest.Name, updatedProjectTest.ManagerId, updatedProjectTest.Deadline).
		WillReturnError(sql.ErrConnDone)

	// Call the UpdateProject method.
	resultProject, err := t.repo.Update(updatedProjectTest)

	// Assertions
	assert.Error(t.T(), err)
	assert.True(t.T(), errors.Is(err, sql.ErrConnDone))
	assert.Equal(t.T(), model.Project{}, resultProject)
}

// DeleteProject method testing methods
func (t *ProjectRepositoryTestSuite) TestProjectRepository_DeleteProject_Success() {
	// Mock the SQL query expectations for DeleteProject with a success outcome.
	t.mockSql.ExpectQuery(`UPDATE projects SET deleted_at = CURRENT_TIMESTAMP WHERE id = \$1 AND deleted_at IS NULL`).
		WithArgs(projectTest.Id).
		WillReturnRows(sqlmock.NewRows([]string{}))

	// Call the DeleteProject method.
	err := t.repo.Delete(projectTest.Id)

	// Assertions
	assert.NoError(t.T(), err)
}

func (t *ProjectRepositoryTestSuite) TestProjectRepository_DeleteProject_ErrorOnQuery() {
	// Mock the SQL query expectations for DeleteProject with an error.
	t.mockSql.ExpectQuery(`UPDATE projects SET deleted_at = CURRENT_TIMESTAMP WHERE id = \$1 AND deleted_at IS NULL`).
		WithArgs(projectTest.Id).
		WillReturnError(sql.ErrConnDone)

	// Call the DeleteProject method.
	err := t.repo.Delete(projectTest.Id)

	// Assertions
	assert.Error(t.T(), err)
	assert.True(t.T(), errors.Is(err, sql.ErrConnDone))
}

func (t *ProjectRepositoryTestSuite) TestProjectRepository_GetAll_ErrorOnQuery() {
	// Mock the SQL query expectations for GetAll with an error.
	t.mockSql.ExpectQuery(`SELECT id, name, manager_id, deadline, created_at, updated_at FROM projects WHERE deleted_at IS NULL ORDER BY deadline DESC LIMIT \$1 OFFSET \$2`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnError(sql.ErrConnDone)

	// Call the GetAll method.
	projects, _, err := t.repo.GetAll(1, 10)

	// Assertions
	assert.Error(t.T(), err)
	assert.True(t.T(), errors.Is(err, sql.ErrConnDone))
	assert.Nil(t.T(), projects)
}

func (t *ProjectRepositoryTestSuite) TestProjectRepository_GetById_ErrorOnQuery() {
	// Mock the SQL query expectations for GetById with an error.
	t.mockSql.ExpectQuery(`SELECT id, name, manager_id, deadline, created_at, updated_at FROM projects WHERE deleted_at IS NULL AND id = \$1`).
		WithArgs(projectTest.Id).
		WillReturnError(sql.ErrConnDone)

	// Call the GetById method.
	resultProject, err := t.repo.GetById(projectTest.Id)

	// Assertions
	assert.Error(t.T(), err)
	assert.True(t.T(), errors.Is(err, sql.ErrConnDone))
	assert.Equal(t.T(), model.Project{}, resultProject)
}

// GetByManagerId method testing methods
func (t *ProjectRepositoryTestSuite) TestProjectRepository_GetByManagerId_Success() {
	// Mock the SQL query expectations for GetByManagerId with a success outcome.
	t.mockSql.ExpectQuery(`SELECT id, name, manager_id, deadline, created_at, updated_at FROM projects WHERE deleted_at IS NULL AND manager_id = \$1`).
		WithArgs(projectTest.ManagerId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "manager_id", "deadline", "created_at", "updated_at"}).
			AddRow(projectTest.Id, projectTest.Name, projectTest.ManagerId, projectTest.Deadline, projectTest.CreatedAt, projectTest.UpdatedAt))

	// Call the GetByManagerId method.
	projects, err := t.repo.GetByManagerId(projectTest.ManagerId)

	// Assertions
	assert.NoError(t.T(), err)
	assert.Len(t.T(), projects, 1)
	assert.Equal(t.T(), projectTest, projects[0])
}

func (t *ProjectRepositoryTestSuite) TestProjectRepository_GetByManagerId_ErrorOnQuery() {
	// Mock the SQL query expectations for GetByManagerId with an error.
	t.mockSql.ExpectQuery(`SELECT id, name, manager_id, deadline, created_at, updated_at FROM projects WHERE deleted_at IS NULL AND manager_id = \$1`).
		WithArgs(projectTest.ManagerId).
		WillReturnError(sql.ErrConnDone)

	// Call the GetByManagerId method.
	projects, err := t.repo.GetByManagerId(projectTest.ManagerId)

	// Assertions
	assert.Error(t.T(), err)
	assert.True(t.T(), errors.Is(err, sql.ErrConnDone))
	assert.Nil(t.T(), projects)
}

// GetByDeadline method testing methods
func (t *ProjectRepositoryTestSuite) TestProjectRepository_GetByDeadline_Success() {
	// Mock the SQL query expectations for GetByDeadline with a success outcome.
	t.mockSql.ExpectQuery(`SELECT id, name, manager_id, deadline, created_at, updated_at FROM projects WHERE deleted_at IS NULL AND deadline = \$1`).
		WithArgs(projectTest.Deadline).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "manager_id", "deadline", "created_at", "updated_at"}).
			AddRow(projectTest.Id, projectTest.Name, projectTest.ManagerId, projectTest.Deadline, projectTest.CreatedAt, projectTest.UpdatedAt))

	// Call the GetByDeadline method.
	projects, err := t.repo.GetByDeadline(projectTest.Deadline)

	// Assertions
	assert.NoError(t.T(), err)
	assert.Len(t.T(), projects, 1)
	assert.Equal(t.T(), projectTest, projects[0])
}

func (t *ProjectRepositoryTestSuite) TestProjectRepository_GetByDeadline_ErrorOnQuery() {
	// Mock the SQL query expectations for GetByDeadline with an error.
	t.mockSql.ExpectQuery(`SELECT id, name, manager_id, deadline, created_at, updated_at FROM projects WHERE deleted_at IS NULL AND deadline = \$1`).
		WithArgs(projectTest.Deadline).
		WillReturnError(sql.ErrConnDone)

	// Call the GetByDeadline method.
	projects, err := t.repo.GetByDeadline(projectTest.Deadline)

	// Assertions
	assert.Error(t.T(), err)
	assert.True(t.T(), errors.Is(err, sql.ErrConnDone))
	assert.Nil(t.T(), projects)
}

// AddProjectMember method testing methods
func (t *ProjectRepositoryTestSuite) TestProjectRepository_AddProjectMember_Success() {
	// Mock the SQL query expectations for AddProjectMember with a success outcome.
	t.mockSql.ExpectQuery(`INSERT INTO project_members\(member_id, project_id\) VALUES \(\$1, \$2\)`).
		WithArgs("member1", projectTest.Id).
		WillReturnRows(sqlmock.NewRows([]string{}))

	// Call the AddProjectMember method.
	err := t.repo.AddProjectMember(projectTest.Id, []string{"member1"})

	// Assertions
	assert.NoError(t.T(), err)
}

func (t *ProjectRepositoryTestSuite) TestProjectRepository_AddProjectMember_ErrorOnQuery() {
	// Mock the SQL query expectations for AddProjectMember with an error.
	t.mockSql.ExpectQuery(`INSERT INTO project_members\(member_id, project_id\) VALUES \(\$1, \$2\)`).
		WithArgs("member1", projectTest.Id).
		WillReturnError(sql.ErrConnDone)

	// Call the AddProjectMember method.
	err := t.repo.AddProjectMember(projectTest.Id, []string{"member1"})

	// Assertions
	assert.Error(t.T(), err)
	assert.True(t.T(), errors.Is(err, sql.ErrConnDone))
}

func (t *ProjectRepositoryTestSuite) TestProjectRepository_GetAllProjectMember_ErrorOnQuery() {
	// Mock the SQL query expectations for GetAllProjectMember with an error.
	t.mockSql.ExpectQuery(`SELECT member_id FROM project_members WHERE project_id = \$1 AND deleted_at IS NULL`).
		WithArgs(projectTest.Id).
		WillReturnError(sql.ErrConnDone)

	// Call the GetAllProjectMember method.
	users, err := t.repo.GetAllProjectMember(projectTest.Id)

	// Assertions
	assert.Error(t.T(), err)
	assert.True(t.T(), errors.Is(err, sql.ErrConnDone))
	assert.Equal(t.T(), users, []model.User{})
}

// DeleteProjectMember method testing methods
func (t *ProjectRepositoryTestSuite) TestProjectRepository_DeleteProjectMember_Success() {
	// Mock the SQL query expectations for DeleteProjectMember with a success outcome.
	t.mockSql.ExpectQuery(`UPDATE project_members SET deleted_at = CURRENT_TIMESTAMP WHERE member_id = \$1 AND project_id = \$2`).
		WithArgs("member1", projectTest.Id).
		WillReturnRows(sqlmock.NewRows([]string{}))

	// Call the DeleteProjectMember method.
	err := t.repo.DeleteProjectMember(projectTest.Id, []string{"member1"})

	// Assertions
	assert.NoError(t.T(), err)
}

func (t *ProjectRepositoryTestSuite) TestProjectRepository_DeleteProjectMember_ErrorOnQuery() {
	// Mock the SQL query expectations for DeleteProjectMember with an error.
	t.mockSql.ExpectQuery(`UPDATE project_members SET deleted_at = CURRENT_TIMESTAMP WHERE member_id = \$1 AND project_id = \$2`).
		WithArgs("member1", projectTest.Id).
		WillReturnError(sql.ErrConnDone)

	// Call the DeleteProjectMember method.
	err := t.repo.DeleteProjectMember(projectTest.Id, []string{"member1"})

	// Assertions
	assert.Error(t.T(), err)
	assert.True(t.T(), errors.Is(err, sql.ErrConnDone))
}

func TestProjectRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ProjectRepositoryTestSuite))
}
