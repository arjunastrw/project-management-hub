package usecase

import (
	"fmt"

	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/repository"
	"enigma.com/projectmanagementhub/shared/shared_model"
)

type TaskUsecase interface {
	GetAll(page int, size int) ([]model.Task, shared_model.Paging, error)
	GetById(Id string) (model.Task, error)
	GetByPersonInCharge(Id string) ([]model.Task, error)
	GetByProjectId(Id string) ([]model.Task, error)
	CreateTask(payload model.Task) (model.Task, error)
	UpdateTask(userId string, payload model.Task) (model.Task, error)
	Delete(id string) error
}

type taskUsecase struct {
	taskRepository    repository.TaskRepository
	userRepository    repository.UserRepository
	projectRepository repository.ProjectRepository
}

// CreateTask implements TaskUsecase.
func (t *taskUsecase) CreateTask(payload model.Task) (model.Task, error) {

	if _, err := t.userRepository.GetById(payload.PersonInCharge); err != nil {
		return model.Task{}, fmt.Errorf("failed to create task. person in charge id invalid")
	}
	if _, err := t.projectRepository.GetById(payload.ProjectId); err != nil {
		return model.Task{}, fmt.Errorf("failed to create task. project id invalid")
	}
	if payload.Name == "" || payload.Deadline == "" {
		return model.Task{}, fmt.Errorf("failed to create task. empty field exist")
	}
	return t.taskRepository.CreateTask(payload)

}

// Delete implements TaskUsecase.
func (t *taskUsecase) Delete(id string) error {
	if _, err := t.taskRepository.GetById(id); err != nil {
		return fmt.Errorf("failed to delete task. task id invalid")
	}
	return t.taskRepository.Delete(id)
}

// GetAll implements TaskUsecase.
func (t *taskUsecase) GetAll(page int, size int) ([]model.Task, shared_model.Paging, error) {
	return t.taskRepository.GetAll(page, size)

}

// GetById implements TaskUsecase.
func (t *taskUsecase) GetById(Id string) (model.Task, error) {
	return t.taskRepository.GetById(Id)
}

// GetByPersonInCharge implements TaskUsecase.
func (t *taskUsecase) GetByPersonInCharge(Id string) ([]model.Task, error) {
	pic, err := t.userRepository.GetById(Id)
	if err != nil {
		return []model.Task{}, fmt.Errorf("failed to get task by person in charge. person in charge id invalid")
	}
	if pic.Task == nil {
		return []model.Task{}, fmt.Errorf("this user currently has no tasks")
	}

	return t.taskRepository.GetByPersonInCharge(Id)
}

// GetByProjectId implements TaskUsecase.
func (t *taskUsecase) GetByProjectId(Id string) ([]model.Task, error) {
	project, err := t.projectRepository.GetById(Id)
	if err != nil {
		return []model.Task{}, fmt.Errorf("failed to get task by project id. project id invalid")
	}
	if project.Tasks == nil {
		return []model.Task{}, fmt.Errorf("this project currently has no tasks")
	}
	return t.taskRepository.GetByProjectId(Id)
}

// UpdateTaskByManager implements TaskUsecase.
func (t *taskUsecase) UpdateTask(userId string, payload model.Task) (model.Task, error) {

	if payload.Status != "In Progress" && payload.Status != "Blocked" && payload.Status != "Waiting Approval" && payload.Status != "Accepted" && payload.Status != "Rejected" && payload.Status != "On Hold" {
		return model.Task{}, fmt.Errorf("invalid status type. status type: ('In Progress', 'Blocked', 'Waiting Approval', 'Accepted', 'Rejected', 'On Hold')")
	}

	user, err := t.userRepository.GetById(userId)
	if err != nil {
		return model.Task{}, fmt.Errorf("failed to update task. user id invalid")
	}

	if user.Role == "MANAGER" {
		if payload.Name == "" || payload.Deadline == "" || payload.Feedback == "" {
			return model.Task{}, fmt.Errorf("failed to update task. empty field exist")
		}

		_, err := t.userRepository.GetById(payload.PersonInCharge)
		if err != nil {
			return model.Task{}, fmt.Errorf("failed to update task. person in charge id invalid")
		}

		return t.taskRepository.UpdateTaskByManager(payload)
	} else {
		check, _ := t.taskRepository.GetById(payload.Id)
		if check.PersonInCharge != userId {
			return model.Task{}, fmt.Errorf("only person in charge and project manager can update task")
		}

		return t.taskRepository.UpdateTaskByMember(payload)
	}
}

// UpdateTaskByMember implements TaskUsecase.

func NewTaskUsecase(taskRepository repository.TaskRepository, userRepository repository.UserRepository, projectRepository repository.ProjectRepository) TaskUsecase {
	return &taskUsecase{
		taskRepository:    taskRepository,
		userRepository:    userRepository,
		projectRepository: projectRepository,
	}
}
