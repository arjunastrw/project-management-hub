package repository

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/shared/shared_model"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TaskRepositoryTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    TaskRepository
}

func (t *TaskRepositoryTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	t.mockDB, t.mockSql = db, mock
	t.repo = NewTaskRepository(t.mockDB)
}

var originalTask = model.Task{
	Id:             "1",
	Name:           "task1",
	ProjectId:      "1",
	Status:         "In Progress",
	PersonInCharge: "user1",
	ApprovalDate:   nil,
	Feedback:       "",
	Deadline:       "2024-01-01",
	Approval:       false,
	CreatedAt:      time.Now(),
	UpdatedAt:      time.Now(),
	DeletedAt:      nil,
}

var updatedTask = model.Task{
	Id:             "1",
	Name:           "task1",
	ProjectId:      "1",
	Status:         "Complete",
	PersonInCharge: "user1",
	ApprovalDate:   nil,
	Feedback:       "approved",
	Deadline:       "2024-01-01",
	Approval:       true,
	CreatedAt:      originalTask.CreatedAt,
	UpdatedAt:      time.Now(),
	DeletedAt:      nil,
}

func (t *TaskRepositoryTestSuite) TestTaskRepository_GetAll_Success() {
	// Mock the SQL query expectations for GetAll.
	rows := sqlmock.NewRows([]string{"id", "name", "status", "approval", "person_in_charge", "deadline", "project_id", "approval_date", "feedback", "created_at", "updated_at"}).
		AddRow(originalTask.Id, originalTask.Name, originalTask.Status, originalTask.Approval, originalTask.PersonInCharge, originalTask.Deadline, originalTask.ProjectId, originalTask.ApprovalDate, originalTask.Feedback, originalTask.CreatedAt, originalTask.UpdatedAt)
	t.mockSql.ExpectQuery(`SELECT id, name, status, approval, person_in_charge, deadline, project_id, approval_date, CASE WHEN feedback IS NULL THEN '-' ELSE feedback END, created_at, updated_at FROM tasks WHERE deleted_at IS NULL ORDER BY deadline DESC LIMIT \$1 OFFSET \$2`).
		WithArgs(10, 0).
		WillReturnRows(rows)
	t.mockSql.ExpectQuery(`SELECT COUNT\(\*\) FROM tasks WHERE deleted_at IS NULL`).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	// Call the GetAll method.
	resultTasks, paging, err := t.repo.GetAll(1, 10)

	// Assertions
	assert.NoError(t.T(), err)
	assert.Len(t.T(), resultTasks, 1)
	assert.Equal(t.T(), originalTask, resultTasks[0])
	assert.Equal(t.T(), 1, paging.TotalRows)
}

func (t *TaskRepositoryTestSuite) TestTaskRepository_GetAll_ErrorOnQuery() {
	// Mock the SQL query expectations for GetAll with an error.
	t.mockSql.ExpectQuery(`SELECT id, name, status, approval, person_in_charge, deadline, project_id, approval_date, CASE WHEN feedback IS NULL THEN '-' ELSE feedback END, created_at, updated_at FROM tasks WHERE deleted_at IS NULL ORDER BY deadline DESC LIMIT \$1 OFFSET \$2`).
		WithArgs(10, 0).
		WillReturnError(sql.ErrConnDone)

	// Call the GetAll method.
	resultTasks, paging, err := t.repo.GetAll(1, 10)

	// Assertions
	assert.Error(t.T(), err)
	assert.True(t.T(), errors.Is(err, sql.ErrConnDone))
	assert.Empty(t.T(), resultTasks)
	assert.Equal(t.T(), shared_model.Paging{}, paging)
}

func (t *TaskRepositoryTestSuite) TestTaskRepository_GetAll_ErrorOnRowScan() {
	// Mock the SQL query expectations for GetAll with an error on row scan.
	rows := sqlmock.NewRows([]string{"id", "name", "status", "approval", "person_in_charge", "deadline", "project_id", "approval_date", "feedback", "created_at", "updated_at"}).
		AddRow("invalid_id", originalTask.Name, originalTask.Status, originalTask.Approval, originalTask.PersonInCharge, originalTask.Deadline, originalTask.ProjectId, originalTask.ApprovalDate, originalTask.Feedback, originalTask.CreatedAt, originalTask.UpdatedAt)
	t.mockSql.ExpectQuery(`SELECT id, name, status, approval, person_in_charge, deadline, project_id, approval_date, CASE WHEN feedback IS NULL THEN '-' ELSE feedback END, created_at, updated_at FROM tasks WHERE deleted_at IS NULL ORDER BY deadline DESC LIMIT \$1 OFFSET \$2`).
		WithArgs(10, 0).
		WillReturnRows(rows)

	// Call the GetAll method.
	resultTasks, paging, err := t.repo.GetAll(1, 10)

	// Assertions
	assert.Error(t.T(), err)
	assert.NotEmpty(t.T(), err.Error()) // Specific error message may vary based on your implementation.
	assert.Empty(t.T(), resultTasks)
	assert.Equal(t.T(), shared_model.Paging{}, paging)
}

func (t *TaskRepositoryTestSuite) TestTaskRepository_GetById_Success() {
	// Mock the SQL query expectations for GetById.
	rows := sqlmock.NewRows([]string{"id", "name", "status", "approval", "person_in_charge", "deadline", "project_id", "approval_date", "feedback", "created_at", "updated_at"}).
		AddRow(originalTask.Id, originalTask.Name, originalTask.Status, originalTask.Approval, originalTask.PersonInCharge, originalTask.Deadline, originalTask.ProjectId, originalTask.ApprovalDate, originalTask.Feedback, originalTask.CreatedAt, originalTask.UpdatedAt)
	t.mockSql.ExpectQuery(`SELECT id, name, status, approval, person_in_charge, deadline, project_id, approval_date, CASE WHEN feedback IS NULL THEN '-' ELSE feedback END, created_at, updated_at FROM tasks WHERE id = \$1 AND deleted_at IS NULL`).
		WithArgs(originalTask.Id).
		WillReturnRows(rows)

	// Call the GetById method.
	resultTask, err := t.repo.GetById(originalTask.Id)

	// Assertions
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), originalTask, resultTask)
}

func (t *TaskRepositoryTestSuite) TestTaskRepository_GetById_NotFound() {
	// Mock the SQL query expectations for GetById with no result.
	t.mockSql.ExpectQuery(`SELECT id, name, status, approval, person_in_charge, deadline, project_id, approval_date, CASE WHEN feedback IS NULL THEN '-' ELSE feedback END, created_at, updated_at FROM tasks WHERE id = \$1 AND deleted_at IS NULL`).
		WithArgs(originalTask.Id).
		WillReturnRows(sqlmock.NewRows([]string{}))

	// Call the GetById method.
	resultTask, err := t.repo.GetById(originalTask.Id)

	// Assertions
	assert.Error(t.T(), err)
	assert.True(t.T(), errors.Is(err, sql.ErrNoRows))
	assert.Equal(t.T(), model.Task{}, resultTask)
}

func (t *TaskRepositoryTestSuite) TestTaskRepository_GetByPersonInCharge_Success() {
	// Mock the SQL query expectations for GetByPersonInCharge.
	rows := sqlmock.NewRows([]string{"id", "name", "status", "approval", "person_in_charge", "deadline", "project_id", "approval_date", "feedback", "created_at", "updated_at"}).
		AddRow(originalTask.Id, originalTask.Name, originalTask.Status, originalTask.Approval, originalTask.PersonInCharge, originalTask.Deadline, originalTask.ProjectId, originalTask.ApprovalDate, originalTask.Feedback, originalTask.CreatedAt, originalTask.UpdatedAt)
	t.mockSql.ExpectQuery(`SELECT id, name, status, approval, person_in_charge, deadline, project_id, approval_date, CASE WHEN feedback IS NULL THEN '-' ELSE feedback END, created_at, updated_at FROM tasks WHERE person_in_charge=\$1 AND deleted_at IS NULL`).
		WithArgs(originalTask.PersonInCharge).
		WillReturnRows(rows)

	// Call the GetByPersonInCharge method.
	resultTasks, err := t.repo.GetByPersonInCharge(originalTask.PersonInCharge)

	// Assertions
	assert.NoError(t.T(), err)
	assert.Len(t.T(), resultTasks, 1)
	assert.Equal(t.T(), originalTask, resultTasks[0])
}

func (t *TaskRepositoryTestSuite) TestTaskRepository_GetByPersonInCharge_EmptyResult() {
	// Mock the SQL query expectations for GetByPersonInCharge with no result.
	t.mockSql.ExpectQuery(`SELECT id, name, status, approval, person_in_charge, deadline, project_id, approval_date, CASE WHEN feedback IS NULL THEN '-' ELSE feedback END, created_at, updated_at FROM tasks WHERE person_in_charge=\$1 AND deleted_at IS NULL`).
		WithArgs(originalTask.PersonInCharge).
		WillReturnRows(sqlmock.NewRows([]string{}))

	// Call the GetByPersonInCharge method.
	resultTasks, err := t.repo.GetByPersonInCharge(originalTask.PersonInCharge)

	// Assertions
	assert.NoError(t.T(), err)
	assert.Empty(t.T(), resultTasks)
}

// Similar tests can be created for GetByProjectId, CreateTask, UpdateTaskByManager, UpdateTaskByMember, and Delete methods.
func (t *TaskRepositoryTestSuite) TestTaskRepository_GetByProjectId_Success() {
	// Mock the SQL query expectations for GetByProjectId.
	rows := sqlmock.NewRows([]string{"id", "name", "status", "approval", "person_in_charge", "deadline", "project_id", "approval_date", "feedback", "created_at", "updated_at"}).
		AddRow(originalTask.Id, originalTask.Name, originalTask.Status, originalTask.Approval, originalTask.PersonInCharge, originalTask.Deadline, originalTask.ProjectId, originalTask.ApprovalDate, originalTask.Feedback, originalTask.CreatedAt, originalTask.UpdatedAt)
	t.mockSql.ExpectQuery(`SELECT id, name, status, approval, person_in_charge, deadline, project_id, approval_date, CASE WHEN feedback IS NULL THEN '-' ELSE feedback END, created_at, updated_at FROM tasks WHERE project_id=\$1 AND deleted_at IS NULL`).
		WithArgs(originalTask.ProjectId).
		WillReturnRows(rows)

	// Call the GetByProjectId method.
	resultTasks, err := t.repo.GetByProjectId(originalTask.ProjectId)

	// Assertions
	assert.NoError(t.T(), err)
	assert.Len(t.T(), resultTasks, 1)
	assert.Equal(t.T(), originalTask, resultTasks[0])
}

func (t *TaskRepositoryTestSuite) TestTaskRepository_GetByProjectId_EmptyResult() {
	// Mock the SQL query expectations for GetByProjectId with no result.
	t.mockSql.ExpectQuery(`SELECT id, name, status, approval, person_in_charge, deadline, project_id, approval_date, CASE WHEN feedback IS NULL THEN '-' ELSE feedback END, created_at, updated_at FROM tasks WHERE project_id=\$1 AND deleted_at IS NULL`).
		WithArgs(originalTask.ProjectId).
		WillReturnRows(sqlmock.NewRows([]string{}))

	// Call the GetByProjectId method.
	resultTasks, err := t.repo.GetByProjectId(originalTask.ProjectId)

	// Assertions
	assert.NoError(t.T(), err)
	assert.Empty(t.T(), resultTasks)
}

func (t *TaskRepositoryTestSuite) TestTaskRepository_CreateTask_Success() {
	// Mock the SQL query expectations for CreateTask.
	rows := sqlmock.NewRows([]string{"id", "name", "person_in_charge", "deadline", "project_id", "created_at"}).
		AddRow(originalTask.Id, originalTask.Name, originalTask.PersonInCharge, originalTask.Deadline, originalTask.ProjectId, originalTask.CreatedAt)
	t.mockSql.ExpectQuery(`INSERT INTO tasks\(name, status, approval, person_in_charge, deadline, project_id, updated_at\) VALUES \(\$1, 'In Progress', false, \$2, \$3, \$4, CURRENT_TIMESTAMP\) RETURNING id, name, person_in_charge, deadline, project_id, created_at`).
		WithArgs(originalTask.Name, originalTask.PersonInCharge, originalTask.Deadline, originalTask.ProjectId).
		WillReturnRows(rows)

	// Call the CreateTask method.
	resultTask, err := t.repo.CreateTask(originalTask)

	// Assertions
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), "In Progress", resultTask.Status)
	assert.False(t.T(), resultTask.Approval)
}

func (t *TaskRepositoryTestSuite) TestTaskRepository_UpdateTaskByManager_Success() {
	// Mock the SQL query expectations for UpdateTaskByManager.
	rows := sqlmock.NewRows([]string{"id", "name", "status", "approval", "person_in_charge", "deadline", "project_id", "approval_date", "feedback", "created_at", "updated_at"}).
		AddRow(updatedTask.Id, updatedTask.Name, updatedTask.Status, updatedTask.Approval, updatedTask.PersonInCharge, updatedTask.Deadline, updatedTask.ProjectId, updatedTask.ApprovalDate, updatedTask.Feedback, updatedTask.CreatedAt, updatedTask.UpdatedAt)
	t.mockSql.ExpectQuery(`UPDATE tasks SET name = \$2, status = \$3, approval = \$4, person_in_charge = \$5, deadline = \$6, approval_date = CURRENT_TIMESTAMP, feedback = \$7, updated_at = CURRENT_TIMESTAMP WHERE id = \$1 AND deleted_at IS NULL RETURNING id, name, status, approval, person_in_charge, deadline, project_id, approval_date, CASE WHEN feedback IS NULL THEN '-' ELSE feedback END, created_at, updated_at`).
		WithArgs(updatedTask.Id, updatedTask.Name, updatedTask.Status, updatedTask.Approval, updatedTask.PersonInCharge, updatedTask.Deadline, updatedTask.Feedback).
		WillReturnRows(rows)

	// Call the UpdateTaskByManager method.
	resultTask, err := t.repo.UpdateTaskByManager(updatedTask)

	// Assertions
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), updatedTask, resultTask)
}

func (t *TaskRepositoryTestSuite) TestTaskRepository_UpdateTaskByMember_Success() {
	// Mock the SQL query expectations for UpdateTaskByMember.
	rows := sqlmock.NewRows([]string{"id", "name", "status", "approval", "person_in_charge", "deadline", "project_id", "approval_date", "feedback", "created_at", "updated_at"}).
		AddRow(updatedTask.Id, updatedTask.Name, updatedTask.Status, updatedTask.Approval, updatedTask.PersonInCharge, updatedTask.Deadline, updatedTask.ProjectId, updatedTask.ApprovalDate, updatedTask.Feedback, updatedTask.CreatedAt, updatedTask.UpdatedAt)
	t.mockSql.ExpectQuery(`UPDATE tasks SET status = \$3, updated_at = CURRENT_TIMESTAMP WHERE id = \$1 AND person_in_charge = \$2 AND deleted_at IS NULL RETURNING id, name, status, approval, person_in_charge, deadline, project_id, approval_date, CASE WHEN feedback IS NULL THEN '-' ELSE feedback END, created_at, updated_at`).
		WithArgs(updatedTask.Id, updatedTask.PersonInCharge, updatedTask.Status).
		WillReturnRows(rows)

	// Call the UpdateTaskByMember method.
	resultTask, err := t.repo.UpdateTaskByMember(updatedTask)

	// Assertions
	assert.NoError(t.T(), err)
	assert.Equal(t.T(), updatedTask, resultTask)
}

func (t *TaskRepositoryTestSuite) TestTaskRepository_DeleteTask_Success() {
	// Mock the SQL query expectations for DeleteTask.
	t.mockSql.ExpectQuery(`UPDATE tasks SET deleted_at = CURRENT_TIMESTAMP WHERE id = \$1 AND deleted_at IS NULL`).
		WithArgs(originalTask.Id).
		WillReturnRows(sqlmock.NewRows([]string{}))

	// Call the DeleteTask method.
	err := t.repo.Delete(originalTask.Id)

	// Assertions
	assert.NoError(t.T(), err)
}

func (t *TaskRepositoryTestSuite) TestTaskRepository_DeleteTask_AlreadyDeleted() {
	// Mock the SQL query expectations for DeleteTask when the task is already deleted.
	t.mockSql.ExpectQuery(`UPDATE tasks SET deleted_at = CURRENT_TIMESTAMP WHERE id = \$1 AND deleted_at IS NULL`).
		WithArgs(originalTask.Id).
		WillReturnError(sql.ErrNoRows)

	// Call the DeleteTask method.
	err := t.repo.Delete(originalTask.Id)

	// Assertions
	assert.Error(t.T(), err)
	assert.True(t.T(), errors.Is(err, sql.ErrNoRows))
}

func TestTaskRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepositoryTestSuite))
}
