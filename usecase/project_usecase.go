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
	GetAll(page int, size int) ([]model.Project, shared_model.Paging, error)
	GetProjectById(id string) (model.Project, error)
	GetProjectsByDeadline(date time.Time) ([]model.Project, error)
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
}

func NewProjectUseCase(projectRepo repository.ProjectRepository) ProjectUseCase {
	return &projectUseCase{
		projectRepo: projectRepo,
	}
}

func (uc *projectUseCase) GetAll(page int, size int) ([]model.Project, shared_model.Paging, error) {
	// validate
	// if strings.ToLower(user.Role) != "admin" {

	// 	return nil, shared_model.Paging{}, errors.New("Unauthorized access")
	// }
	projects, paging, err := uc.projectRepo.GetAll(page, size)
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

	project, err := uc.projectRepo.GetById(id)
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
	projects, err := uc.projectRepo.GetByDeadline(date)
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

func (uc *projectUseCase) GetProjectsByManagerId(id string) ([]model.Project, error) {
	// validate
	// if strings.ToLower(user.Role) != "admin" {

	// 	return nil, shared_model.Paging{}, errors.New("Unauthorized access")
	// }

	projects, err := uc.projectRepo.GetByManagerId(id)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to get projects by ManagerId: %s", err.Error())
		log.Printf("ERROR: %s", errorMessage)
		return nil, errors.New(errorMessage)
	}

	successMessage := fmt.Sprintf("Successfully retrieved projects with Manager Id %s", id)
	log.Printf("INFO: %s", successMessage)

	if len(projects) == 0 {
		warningMessage := "No projects found with the given Manager Id"
		log.Printf("WARNING: %s", warningMessage)
	}

	return projects, nil

}

func (uc *projectUseCase) GetProjectsByMemberId(id string) ([]model.Project, error) {

	// if strings.ToLower(user.Role) != "member" {

	// 	return nil, shared_model.Paging{}, errors.New("Unauthorized access")
	// }

	projects, err := uc.projectRepo.GetByMemberId(id)
	if err != nil {

		errorMessage := fmt.Sprintf("Failed to get projects by MaemberId: %s", err.Error())
		log.Printf("ERROR: %s", errorMessage)
		return nil, errors.New(errorMessage)
	}

	successMessage := fmt.Sprintf("Successfully retrieved projects with Member Id %s", id)
	log.Printf("INFO: %s", successMessage)

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

func (uc *projectUseCase) AddProjectMember(id string, members []string) error {
	// if !uc.isAdminOrProjectManager(user) {

	// 	return nil, errors.New("Unauthorized access")
	// }

	// Call the corresponding method from the repository

	err := uc.projectRepo.AddProjectMember(id, members)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to add project members: %s", err.Error())
		log.Printf("ERROR: %s", errorMessage)
		return errors.New(errorMessage)
	}

	successMessage := "Successfully added project members"
	log.Printf("INFO: %s", successMessage)

	return nil

}

func (uc *projectUseCase) DeleteProjectMember(id string, members []string) error {
	// Add validation logic here if needed
	// if !uc.isAdminOrProjectManager(user) {
	//     return errors.New("Unauthorized access")
	// }

	// Call the corresponding method from the repository
	err := uc.projectRepo.DeleteProjectMember(id, members)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to delete project members: %s", err.Error())
		log.Printf("ERROR: %s", errorMessage)
		return errors.New(errorMessage)
	}

	successMessage := "Successfully deleted project members"
	log.Printf("INFO: %s", successMessage)

	return nil
}

func (uc *projectUseCase) GetAllProjectMember(id string) ([]model.User, error) {

	// if !uc.isAdminOrProjectManager(user) {

	// 	return nil, errors.New("Unauthorized access")
	// }

	users, err := uc.projectRepo.GetAllProjectMember(id)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to get project members: %s", err.Error())
		log.Printf("ERROR: %s", errorMessage)
		return []model.User{}, errors.New(errorMessage)
	}
	return users, nil
}

func (uc *projectUseCase) Delete(id string) error {

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

func (uc *projectUseCase) Update(payload model.Project) (model.Project, error) {
	// Add validation or logic here if needed
	// if !uc.isAdminOrProjectManager(user) {
	//     return model.Project{}, errors.New("Unauthorized access")
	// }

	updatedProject, err := uc.projectRepo.Update(payload)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to update project: %s", err.Error())
		log.Printf("ERROR: %s", errorMessage)
		return model.Project{}, errors.New(errorMessage)
	}

	successMessage := "Successfully updated project"
	log.Printf("INFO: %s", successMessage)

	return updatedProject, nil
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
