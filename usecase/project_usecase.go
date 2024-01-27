package usecase

import (
	"fmt"

	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/repository"
	"enigma.com/projectmanagementhub/shared/shared_model"
)

type ProjectUseCase interface {
	GetAll(page int, size int) ([]model.Project, shared_model.Paging, error)
	GetProjectById(id string) (model.Project, error)
	GetProjectsByDeadline(date string) ([]model.Project, error)
	GetProjectsByManagerId(id string) ([]model.Project, error)
	GetProjectsByMemberId(id string) ([]model.Project, error)
	CreateNewProject(payload model.Project) (model.Project, error)
	AddProjectMember(id string, members []string) error
	DeleteProjectMember(id string, members []string) error
	GetAllProjectMember(id string) ([]model.User, error)
	Update(payload model.Project) (model.Project, error)
	Delete(id string) error
}

type projectUseCase struct {
	projectRepo repository.ProjectRepository
	userRepo    repository.UserRepository
}

func NewProjectUseCase(projectRepo repository.ProjectRepository, userRepo repository.UserRepository) ProjectUseCase {
	return &projectUseCase{
		projectRepo: projectRepo,
		userRepo:    userRepo,
	}
}

func (uc *projectUseCase) GetAll(page int, size int) ([]model.Project, shared_model.Paging, error) {

	projects, paging, err := uc.projectRepo.GetAll(page, size)
	if err != nil {
		errorMessage := fmt.Errorf(" Failed to get projects: %s", err.Error())
		return nil, shared_model.Paging{}, errorMessage
	}

	return projects, paging, nil
}

func (uc *projectUseCase) GetProjectById(id string) (model.Project, error) {

	project, err := uc.projectRepo.GetById(id)
	if err != nil {
		errorMessage := fmt.Errorf(" Failed to get projects: %s", err.Error())
		return model.Project{}, errorMessage
	}

	return project, nil
}

func (uc *projectUseCase) GetProjectsByDeadline(date string) ([]model.Project, error) {

	projects, err := uc.projectRepo.GetByDeadline(date)
	if err != nil {
		errorMessage := fmt.Errorf(" Failed to get projects: %s", err.Error())
		return nil, errorMessage
	}

	if len(projects) == 0 {
		errorMessage := fmt.Errorf(" No projects found with the given deadline")
		return nil, errorMessage
	}

	return projects, nil
}

func (uc *projectUseCase) GetProjectsByManagerId(id string) ([]model.Project, error) {
	manager, err := uc.userRepo.GetById(id)
	if err != nil {
		errorMessage := fmt.Errorf(" Failed to get projects by ManagerId: invalid id")
		return nil, errorMessage
	}
	if manager.Role == "TEAM MEMBER" {
		errorMessage := fmt.Errorf(" Failed to get projects by ManagerId: Unauthorized")
		return nil, errorMessage
	}
	projects, err := uc.projectRepo.GetByManagerId(id)
	if err != nil {
		errorMessage := fmt.Errorf(" Failed to get projects by ManagerId: %s", err.Error())
		return nil, errorMessage
	}

	if len(projects) == 0 {
		errorMessage := fmt.Errorf(" No projects found with the given Manager Id")
		return nil, errorMessage
	}

	return projects, nil

}

func (uc *projectUseCase) GetProjectsByMemberId(id string) ([]model.Project, error) {

	_, err := uc.userRepo.GetById(id)
	if err != nil {

		errorMessage := fmt.Errorf(" Failed to get projects by MemberId: invalid id")
		return nil, errorMessage
	}

	projects, err := uc.projectRepo.GetByMemberId(id)
	if err != nil {

		errorMessage := fmt.Errorf(" Failed to get projects by MemberId: no project currently")

		return nil, errorMessage
	}

	return projects, nil

}

func (uc *projectUseCase) CreateNewProject(payload model.Project) (model.Project, error) {

	if payload.Name == "" || payload.ManagerId == "" || payload.Deadline == "" {
		errorMessage := fmt.Errorf(" Fields 'name', 'manager id', 'deadline' cannot be empty")

		return model.Project{}, errorMessage
	}

	createdProject, err := uc.projectRepo.CreateProject(payload)
	if err != nil {
		errorMessage := fmt.Errorf(" Failed to create project: %s", err.Error())

		return model.Project{}, errorMessage
	}

	return createdProject, nil

}

func (uc *projectUseCase) AddProjectMember(id string, members []string) error {
	memberscheck, err := uc.projectRepo.GetAllProjectMember(id)
	if err != nil {
		errorMessage := fmt.Errorf(" Failed to add project members: %s", err.Error())
		return errorMessage
	}
	for _, membercheck := range memberscheck {
		for _, member := range members {
			if membercheck.Id == member {
				errorMessage := fmt.Errorf(" Failed to add member. Some member(s) already in project")
				return errorMessage
			}
		}
	}
	err = uc.projectRepo.AddProjectMember(id, members)
	if err != nil {
		errorMessage := fmt.Errorf(" Failed to add project members: %s", err.Error())

		return errorMessage
	}

	return nil

}

func (uc *projectUseCase) DeleteProjectMember(id string, members []string) error {

	for _, member := range members {
		x, err := uc.userRepo.GetById(member)
		if err != nil {
			errorMessage := fmt.Errorf(" Failed to delete project members: %s", err.Error())
			return errorMessage
		}
		for _, task := range x.Task {
			if task.ProjectId == id {
				errorMessage := fmt.Errorf(" Failed to delete project members: member with id %s stile have task in project id %s", x.Id, id)
				return errorMessage
			}
		}
	}
	err := uc.projectRepo.DeleteProjectMember(id, members)
	if err != nil {
		errorMessage := fmt.Errorf(" Failed to delete project members: %s", err.Error())

		return errorMessage
	}

	return nil
}

func (uc *projectUseCase) GetAllProjectMember(id string) ([]model.User, error) {

	_, err := uc.projectRepo.GetById(id)
	if err != nil {
		errorMessage := fmt.Errorf(" Failed to get project members: invalid id")
		return []model.User{}, errorMessage
	}

	users, err := uc.projectRepo.GetAllProjectMember(id)
	if err != nil {
		errorMessage := fmt.Errorf(" Failed to get project members: %s", err.Error())

		return []model.User{}, errorMessage
	}
	return users, nil
}

func (uc *projectUseCase) Delete(id string) error {

	err := uc.projectRepo.Delete(id)
	if err != nil {
		errorMessage := fmt.Errorf(" Failed to delete project: %s", err.Error())

		return errorMessage
	}

	return nil
}

func (uc *projectUseCase) Update(payload model.Project) (model.Project, error) {

	_, err := uc.projectRepo.GetById(payload.Id)
	if err != nil {

		errorMessage := fmt.Errorf(" Failed to update project: invalid id")
		return model.Project{}, errorMessage
	}
	manager, err := uc.userRepo.GetById(payload.ManagerId)
	if err != nil {

		errorMessage := fmt.Errorf(" Failed to update project: invalid manager id")
		return model.Project{}, errorMessage
	}
	if manager.Role != "MANAGER" {

		errorMessage := fmt.Errorf(" Failed to update project: manager id is not manager")
		return model.Project{}, errorMessage
	}

	if payload.Name == "" || payload.Deadline == "" {

		errorMessage := fmt.Errorf(" Fields 'name', 'deadline' cannot be empty")
		return model.Project{}, errorMessage
	}

	updatedProject, err := uc.projectRepo.Update(payload)
	if err != nil {
		errorMessage := fmt.Errorf(" Failed to update project: %s", err.Error())

		return model.Project{}, errorMessage
	}

	return updatedProject, nil
}
