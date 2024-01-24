package repository

import (
	"database/sql"
	"time"

	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/shared/shared_model"
)

type ProjectRepository interface {
	GetAll(page int, size int) ([]model.Project, shared_model.Paging, error)
	GetById(id string) (model.Project, error)
	GetByDeadline(date time.Time) ([]model.Project, shared_model.Paging, error)
	GetByManagerId(id string) ([]model.Project, shared_model.Paging, error)
	GetByMemberId(id string) ([]model.Project, shared_model.Paging, error)
	CreateProject(payload model.Project) (model.Project, error)
	EditProjectMember(id string, members []string) ([]model.User, error)
	GetAllProjectMember(id string) ([]model.User, error)
	Update(payload model.Project) (model.Project, error)
	Delete(string string) error
}

type projectRepository struct {
	sql *sql.DB
}

func NewProjectRepository(sql *sql.DB) ProjectRepository {
	return &projectRepository{
		sql: sql,
	}
}
