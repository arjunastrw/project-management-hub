package usecase

import (
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
	UpdateTaskByManager(payload model.Task) (model.Task, error)
	UpdateTaskByMember(payload model.Task) (model.Task, error)
	Delete(id string) error
}

type taskUsecase struct {
	taskRepository repository.TaskRepository
}

// CreateTask implements TaskUsecase.
func (t *taskUsecase) CreateTask(payload model.Task) (model.Task, error) {
	//cek role admin
	//cek isi name, person_in_charge, project_id, deadline
	return t.taskRepository.CreateTask(payload)
}

// Delete implements TaskUsecase.
func (t *taskUsecase) Delete(id string) error {
	//cek role admin atau manager project tsb
	//cek task id exist
	return t.taskRepository.Delete(id)
}

// GetAll implements TaskUsecase.
func (t *taskUsecase) GetAll(page int, size int) ([]model.Task, shared_model.Paging, error) {
	//cek role admin
	return t.taskRepository.GetAll(page, size)

}

// GetById implements TaskUsecase.
func (t *taskUsecase) GetById(Id string) (model.Task, error) {
	//cek role admin atau manager projek tsb
	//cek task id nya exist
	return t.taskRepository.GetById(Id)
}

// GetByPersonInCharge implements TaskUsecase.
func (t *taskUsecase) GetByPersonInCharge(Id string) ([]model.Task, error) {
	//cek role admin atau id sendiri
	//cek id user buat pic exist
	return t.taskRepository.GetByPersonInCharge(Id)
}

// GetByProjectId implements TaskUsecase.
func (t *taskUsecase) GetByProjectId(Id string) ([]model.Task, error) {
	//cek role admin atau manager project tsb atau team member tsb

	return t.taskRepository.GetByProjectId(Id)
}

// UpdateTaskByManager implements TaskUsecase.
func (t *taskUsecase) UpdateTaskByManager(payload model.Task) (model.Task, error) {
	//cek role manager project tsb
	//cek payload.Id exist, payload.Name not "", payload.Status antara ('In Progress', 'Blocked', 'Waiting Approval', 'Accepted', 'Rejected', 'On Hold'), payload.Approval ada, payload.PersonInCharge exist, payload.Deadline exist, payload.Feedback not ""
	return t.taskRepository.UpdateTaskByManager(payload)
}

// UpdateTaskByMember implements TaskUsecase.
func (t *taskUsecase) UpdateTaskByMember(payload model.Task) (model.Task, error) {
	//cek role team member, pic dari task tsb
	//cek id exist, status antara ('In Progress', 'Blocked', 'Waiting Approval', 'Accepted', 'Rejected', 'On Hold')

	return t.taskRepository.UpdateTaskByMember(payload)
}

func NewTaskUsecase(taskRepository repository.TaskRepository) TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
	}
}
