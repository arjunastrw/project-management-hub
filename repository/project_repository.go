package repository

import (
	"database/sql"
	"log"
	"math"
	"slices"
	"time"

	"enigma.com/projectmanagementhub/config"
	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/shared/shared_model"
)

type ProjectRepository interface {
	GetAll(page int, size int) ([]model.Project, shared_model.Paging, error)
	GetById(id string) (model.Project, error)
	GetByDeadline(date time.Time) ([]model.Project, error)
	GetByManagerId(id string) ([]model.Project, error)
	GetByMemberId(id string) ([]model.Project, error)
	CreateProject(payload model.Project) (model.Project, error)
	AddProjectMember(id string, members []string) error
	DeleteProjectMember(id string, members []string) error
	GetAllProjectMember(id string) ([]model.User, error)
	Update(payload model.Project) (model.Project, error)
	Delete(id string) error
}

type projectRepository struct {
	db *sql.DB
}

// DeleteProjectMember implements ProjectRepository.
func (p *projectRepository) DeleteProjectMember(id string, members []string) error {

	row, err := p.db.Query(config.GetAllProjectMember, id)
	if err != nil {
		log.Println("project_repository.Query", err.Error())
		return err
	}

	for row.Next() {
		var memberid string
		//updated_at cannot be nil
		err1 := row.Scan(&memberid)
		if err1 != nil {
			log.Println("projectRepository.Rows.Next", err1.Error())
			return err1
		}

		if slices.Contains(members, memberid) {
			_, err := p.db.Query(config.DeleteProjectMember, memberid, id)
			if err != nil {
				log.Println("project_repository.Query", err.Error())
				return err
			}
		}
	}

	return nil
}

// CreateProject implements ProjectRepository.
func (p *projectRepository) CreateProject(payload model.Project) (model.Project, error) {
	var project model.Project

	err := p.db.QueryRow(config.CreateProject, payload).Scan(&project.Id, &project.Name, &project.ManagerId, &project.Deadline, &project.CreatedAt, &project.UpdatedAt)
	if err != nil {
		log.Println("project_repository.QueryRow", err.Error())
		return model.Project{}, err
	}
	return project, nil
}

// Delete implements ProjectRepository.
func (p *projectRepository) Delete(id string) error {
	_, err := p.db.Query(config.DeleteProject, id)
	if err != nil {
		log.Println("project_repository.Query", err.Error())
		return err
	}

	return nil
}

// AddProjectMember implements ProjectRepository.
func (p *projectRepository) AddProjectMember(id string, members []string) error {

	row, err := p.db.Query(config.GetAllProjectMember, id)
	if err != nil {
		log.Println("project_repository.Query", err.Error())
		return err
	}

	for row.Next() {
		var memberid string
		//updated_at cannot be nil
		err1 := row.Scan(&memberid)
		if err1 != nil {
			log.Println("projectRepository.Rows.Next", err1.Error())
			return err1
		}

		if !slices.Contains(members, memberid) {
			_, err := p.db.Query(config.AddProjectMember, memberid, id)
			if err != nil {
				log.Println("project_repository.Query", err.Error())
				return err
			}
		}
	}

	return nil

}

// GetAll implements ProjectRepository.
func (p *projectRepository) GetAll(page int, size int) ([]model.Project, shared_model.Paging, error) {
	var projects []model.Project
	offset := (page - 1) * size
	row, err := p.db.Query(config.GetAllProject, size, offset)
	if err != nil {
		log.Println("project_repository.QueryRow", err.Error())
		return nil, shared_model.Paging{}, err
	}

	for row.Next() {
		project := model.Project{}
		//updated_at cannot be nil
		err := row.Scan(&project.Id, &project.Name, &project.ManagerId, &project.Deadline, &project.CreatedAt, &project.UpdatedAt)
		if err != nil {
			log.Println("projectRepository.Rows.Next", err.Error())
			return nil, shared_model.Paging{}, err
		}

		projects = append(projects, project)
	}

	totalRows := 0

	if err := p.db.QueryRow(config.CountAllUser).Scan(&totalRows); err != nil {
		return nil, shared_model.Paging{}, err
	}

	paging := shared_model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}

	return projects, paging, nil
}

// GetAllProjectMember implements ProjectRepository.
func (p *projectRepository) GetAllProjectMember(id string) ([]model.User, error) {

	var users []model.User
	row, err := p.db.Query(config.GetAllProjectMember, id)
	if err != nil {
		log.Println("project_repository.Query", err.Error())
		return []model.User{}, err
	}

	for row.Next() {
		var memberid string
		var user model.User
		//updated_at cannot be nil
		err1 := row.Scan(&memberid)
		if err1 != nil {
			log.Println("projectRepository.Rows.Next", err1.Error())
			return []model.User{}, err1
		}

		err := p.db.QueryRow(config.GetUserByID, id).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Println("user not found", err.Error())
			return []model.User{}, err
		}
		users = append(users, user)

	}

	return users, err
}

// GetByDeadline implements ProjectRepository.
func (p *projectRepository) GetByDeadline(date time.Time) ([]model.Project, error) {
	var projects []model.Project
	row, err := p.db.Query(config.GetProjectByDeadline, date)
	if err != nil {
		log.Println("project_repository.QueryRow", err.Error())
		return nil, err
	}

	for row.Next() {
		project := model.Project{}
		//updated_at cannot be nil
		err := row.Scan(&project.Id, &project.Name, &project.Deadline, &project.CreatedAt, &project.UpdatedAt)
		if err != nil {
			log.Println("projectRepository.Rows.Next", err.Error())
			return nil, err
		}

		projects = append(projects, project)
	}

	return projects, nil
}

// GetById implements ProjectRepository.
func (p *projectRepository) GetById(id string) (model.Project, error) {
	var project model.Project

	err := p.db.QueryRow(config.GetProjectByID, id).Scan(&project.Id, &project.Name, &project.Deadline, &project.CreatedAt, &project.UpdatedAt)
	if err != nil {
		log.Println("project_repository.QueryRow", err.Error())
		return model.Project{}, err
	}
	return project, nil
}

// GetByManagerId implements ProjectRepository.
func (p *projectRepository) GetByManagerId(id string) ([]model.Project, error) {
	var projects []model.Project
	row, err := p.db.Query(config.GetProjectByManagerID, id)
	if err != nil {
		log.Println("project_repository.QueryRow", err.Error())
		return nil, err
	}

	for row.Next() {
		project := model.Project{}
		//updated_at cannot be nil
		err := row.Scan(&project.Id, &project.Name, &project.Deadline, &project.CreatedAt, &project.UpdatedAt)
		if err != nil {
			log.Println("projectRepository.Rows.Next", err.Error())
			return nil, err
		}

		projects = append(projects, project)
	}

	return projects, nil
}

// GetByMemberId implements ProjectRepository.
func (p *projectRepository) GetByMemberId(id string) ([]model.Project, error) {
	var projects []model.Project
	row, err := p.db.Query(config.GetAllProjectByMemberID, id)
	if err != nil {
		log.Println("project_repository.QueryRow", err.Error())
		return nil, err
	}

	for row.Next() {
		project := model.Project{}
		//updated_at cannot be nil
		err := row.Scan(&project.Id, &project.Name, &project.Deadline, &project.CreatedAt, &project.UpdatedAt)
		if err != nil {
			log.Println("projectRepository.Rows.Next", err.Error())
			return nil, err
		}

		projects = append(projects, project)
	}

	return projects, nil
}

// Update implements ProjectRepository.
func (p *projectRepository) Update(payload model.Project) (model.Project, error) {
	var project model.Project

	err := p.db.QueryRow(config.UpdateProject, payload.Id).Scan(&project.Id, &project.Name, &project.Deadline, &project.CreatedAt, &project.UpdatedAt)
	if err != nil {
		log.Println("user_repository.QueryRow", err.Error())
		return model.Project{}, err
	}
	return project, nil
}

func NewProjectRepository(db *sql.DB) ProjectRepository {
	return &projectRepository{
		db: db,
	}
}
