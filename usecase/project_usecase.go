package usecase

import (
	"errors"
	"fmt"
	"log"
	"time"

	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/repository"
	"enigma.com/projectmanagementhub/shared/shared_model"
)

type ProjectUseCase interface {
	GetAllProjects(page int, size int) ([]model.Project, shared_model.Paging, error)
	GetProjectById(id string) (model.Project, error)
	GetProjectsByDeadline(date time.Time) ([]model.Project, shared_model.Paging, error)
	GetProjectsByManagerId(id string) ([]model.Project, shared_model.Paging, error)
	GetProjectsByMemberId(id string) ([]model.Project, shared_model.Paging, error)
	CreateNewProject(payload model.Project) (model.Project, error)
	EditProjectMembers(id string, members []string) ([]model.User, error)
	GetAllProjectMembers(id string) ([]model.User, error)
	UpdateProject(payload model.Project) (model.Project, error)
	DeleteProject(id string) error
}
type projectUseCase struct {
	projectRepo repository.ProjectRepository
}

func NewProjectUseCase(projectRepo repository.ProjectRepository) ProjectUseCase {
	return &projectUseCase{
		projectRepo: projectRepo,
	}
}

func (uc *projectUseCase) GetAllProject(page int, size int) ([]model.Project, shared_model.Paging, error) {
	// validate
	// if strings.ToLower(user.Role) != "admin" {

	// 	return nil, shared_model.Paging{}, errors.New("Unauthorized access")
	// }
	projects, paging, err := uc.projectRepo.GetAllProject(page, size)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to get projects: %s", err.Error())
		log.Printf("ERROR: %s", errorMessage)
		return nil, shared_model.Paging{}, errors.New(errorMessage)
	}

	projectList := make([]model.Project, len(projects))
	for i, project := range projects {
		projectList[i] = model.Project{
			Id:        project.Id,
			Name:      project.Name,
			ManagerId: project.ManagerId,
			Deadline:  project.Deadline,
			CreatedAt: project.CreatedAt,
			UpdatedAt: project.UpdatedAt,
			DeletedAt: project.DeletedAt,
			Members:   project.Members,
			Tasks:     project.Tasks,
		}
	}

	return projectList, paging, nil
}

func (uc *projectUseCase) GetProjectById(id string) (model.Project, error) {

	project, err := uc.projectRepo.GetProjectById(id)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to get projects: %s", err.Error())
		log.Printf("ERROR: %s", errorMessage)
		return model.Project{}, errors.New(errorMessage)
	}

	// if !uc.isUserTeamMemberOrAdmin(user, project) {
	// 	return model.Project{}, errors.New("Unauthorized access")
	// }
	return project, nil
}

func (uc *projectUseCase) GetProjectsByDeadline(date time.Time) ([]model.Project, error) {
	//// validate
	//if strings.ToLower(user.Role) != "admin" {
	//	return nil, shared_model.Paging{}, errors.New("Unauthorized access")
	//}
	projects, err := uc.projectRepo.GetProjectByDeadline(date)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to get projects: %s", err.Error())
		log.Printf("ERROR: %s", errorMessage)
		return nil, errors.New(errorMessage)
	}

	successMessage := fmt.Sprintf("Successfully retrieved projects with deadline %s", date.String())
	log.Printf("INFO: %s", successMessage)

	if len(projects) == 0 {
		warningMessage := "No projects found with the given deadline"
		log.Printf("WARNING: %s", warningMessage)
	}

	return projects, nil
}

func (uc *projectUseCase) GetProjectsByManagerId(id string) ([]model.Project, shared_model.Paging, error) {
	// validate
	// if strings.ToLower(user.Role) != "admin" {

	// 	return nil, shared_model.Paging{}, errors.New("Unauthorized access")
	// }

	projects, err := uc.projectRepo.GetByManagerId(id)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to get projects by ManagerId: %s", err.Error())
		log.Printf("ERROR: %s", errorMessage)
		return nil, shared_model.Paging{}, errors.New(errorMessage)
	}

	successMessage := fmt.Sprintf("Successfully retrieved projects with Manager Id %s", id)
	log.Printf("INFO: %s", successMessage)

	if len(projects) == 0 {
		warningMessage := "No projects found with the given Manager Id"
		log.Printf("WARNING: %s", warningMessage)
	}

	return projects, shared_model.Paging{}, nil

}

func (uc *projectUseCase) GetProjectsByMemberId(id string) ([]model.Project, error) {

	// if strings.ToLower(user.Role) != "member" {

	// 	return nil, shared_model.Paging{}, errors.New("Unauthorized access")
	// }

	projects, err := uc.projectRepo.GetProjectByMemberId(id)
	if err != nil {
		return nil, err
	}
	return projects, nil

}

func (uc *projectUseCase) CreateNewProject(payload model.Project) (model.Project, error) {
	// if !uc.isAdminOrProjectManager(user) {

	// 	return nil, errors.New("Unauthorized access")
	// }

	if payload.Name == "" || payload.ManagerId == "" || payload.Deadline.IsZero() || len(payload.Members) == 0 || len(payload.Tasks) == 0 {
		errorMessage := "Fields 'name', 'manager id', 'deadline', 'members', and 'tasks' cannot be empty"
		log.Printf("ERROR: %s", errorMessage)
		return model.Project{}, errors.New(errorMessage)
	}

	createdProject, err := uc.projectRepo.CreateProject(payload)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to create project: %s", err.Error())
		log.Printf("ERROR: %s", errorMessage)
		return model.Project{}, errors.New(errorMessage)
	}

	successMessage := "Successfully created project"
	log.Printf("INFO: %s", successMessage)

	return createdProject, nil

}

func (uc *projectUseCase) EditProjectMembers(id string, members []string) ([]model.User, error) {

	// if !uc.isAdminOrProjectManager(user) {

	// 	return nil, errors.New("Unauthorized access")
	// }

	editedMembers, err := uc.projectRepo.EditProjectMember(id, members)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to edit project members: %s", err.Error())
		log.Printf("ERROR: %s", errorMessage)
		return nil, errors.New(errorMessage)
	}

	successMessage := "Successfully edited project members"
	log.Printf("INFO: %s", successMessage)

	return editedMembers, nil
}

func (uc *projectUseCase) GetAllProjectMembers(id string) ([]model.User, error) {

	// if !uc.isAdminOrProjectManager(user) {

	// 	return nil, errors.New("Unauthorized access")
	// }

	return uc.projectRepo.GetAllProjectMember(id)

}

func (uc *projectUseCase) UpdateProject(payload model.Project) (model.Project, error) {
	//if user.Role != "admin" {
	//
	//	return nil, errors.New("Unauthorized access")
	//}

	return uc.projectRepo.Update(payload)

}

func (uc *projectUseCase) DeleteProject(id string) error {

	// if user.Role != "admin" {

	// 	return nil, errors.New("Unauthorized access")
	// }

	err := uc.projectRepo.Delete(id)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to delete project: %s", err.Error())
		log.Printf("ERROR: %s", errorMessage)
		return errors.New(errorMessage)
	}

	successMessage := "Successfully deleted project"
	log.Printf("INFO: %s", successMessage)

	return nil
}

//func (uc *projectUseCase) isUserTeamMemberOrAdmin(user model.User, project model.Project) bool {
//	//validate
//	for _, member := range project.Members {
//		if member.Id == user.Id || user.Role == "ADMIN" {
//			return true
//		}
//	}
//	return false
//}
//
//func (uc *projectUseCase) isAdminOrProjectManager(user model.User) bool {
//	return user.Role == "ADMIN" || user.Role == "MANAGER"
//}
