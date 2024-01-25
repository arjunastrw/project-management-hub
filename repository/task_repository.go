package repository

import (
	"database/sql"
	"log"
	"math"

	"enigma.com/projectmanagementhub/config"
	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/shared/shared_model"
)

/*
type Task struct {
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	Status         string    `json:"status"`
	Approval       bool      `json:"approval"`
	ApprovalDate   time.Time `json:"approval_date"`
	Feedback       string    `json:"feedback"`
	PersonInCharge string    `json:"person_in_charge"`
	ProjectId      string    `json:"project_id"`
	Deadline       time.Time `json:"deadline"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}
*/

type TaskRepository interface {
	GetAll(page int, size int) ([]model.Task, shared_model.Paging, error)
	GetById(Id string) (model.Task, error)
	GetByPersonInCharge(Id string) ([]model.Task, error)
	GetByProjectId(Id string) ([]model.Task, error)
	CreateTask(payload model.Task) (model.Task, error)
	UpdateTaskByManager(payload model.Task) (model.Task, error)
	UpdateTaskByMember(payload model.Task) (model.Task, error)
	Delete(id string) error
}

type taskRepository struct {
	db *sql.DB
}

// UpdateTaskByManager implements TaskRepository.
func (t *taskRepository) UpdateTaskByManager(payload model.Task) (model.Task, error) {

	var task model.Task
	err := t.db.QueryRow(config.UpdateTaskByManager, payload.Id, payload.Name, payload.Status, payload.Approval, payload.PersonInCharge, payload.Deadline, payload.Feedback).Scan(&task.Id, &task.Name, &task.Status, &task.Approval, &task.PersonInCharge, &task.Deadline, &task.ProjectId, &task.ApprovalDate, &task.Feedback, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		log.Println("task_repository.QueryRow", err.Error())
		return model.Task{}, err
	}

	return task, nil

}

// UpdateTaskByMember implements TaskRepository.
func (t *taskRepository) UpdateTaskByMember(payload model.Task) (model.Task, error) {

	var task model.Task

	err := t.db.QueryRow(config.UpdateTaskByMember, payload.Id, payload.PersonInCharge, payload.Status).Scan(&task.Id, &task.Name, &task.Status, &task.Approval, &task.PersonInCharge, &task.Deadline, &task.ProjectId, &task.ApprovalDate, &task.Feedback, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		log.Println("task_repository.QueryRow", err.Error())
		return model.Task{}, err
	}

	return task, nil
}

// CreateTask implements TaskRepository.
func (t *taskRepository) CreateTask(payload model.Task) (model.Task, error) {

	var task model.Task

	err := t.db.QueryRow(config.CreateTask, payload.Id, payload.PersonInCharge, payload.Deadline, payload.ProjectId).Scan(&task.Id, &task.Name, &task.PersonInCharge, &task.Deadline, &task.ProjectId, &task.CreatedAt)
	if err != nil {
		log.Println("task_repository.QueryRow", err.Error())
		return model.Task{}, err
	}
	task.Status = "In Progress"
	task.Approval = false
	task.UpdatedAt = task.CreatedAt

	return task, nil
}

// Delete implements TaskRepository.
func (t *taskRepository) Delete(id string) error {

	_, err := t.db.Query(config.DeleteTask, id)
	if err != nil {
		log.Println("task_repository.Query", err.Error())
		return err
	}
	return nil
}

// GetAll implements TaskRepository.
func (t *taskRepository) GetAll(page int, size int) ([]model.Task, shared_model.Paging, error) {

	var tasks []model.Task
	offset := (page - 1) * size
	row, err := t.db.Query(config.GetAllTask, size, offset)
	if err != nil {
		log.Println("task_repository.Query", err.Error())
		return nil, shared_model.Paging{}, err
	}

	for row.Next() {
		task := model.Task{}
		//updated_at cannot be nil
		err := row.Scan(&task.Id, &task.Name, &task.Status, &task.Approval, &task.PersonInCharge, &task.Deadline, &task.ProjectId, &task.ApprovalDate, &task.Feedback, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			log.Println("taskRepository.Rows.Next", err.Error())
			return nil, shared_model.Paging{}, err
		}

		tasks = append(tasks, task)
	}

	totalRows := 0

	if err := t.db.QueryRow(config.CountAllTask).Scan(&totalRows); err != nil {
		return nil, shared_model.Paging{}, err
	}

	paging := shared_model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}

	return tasks, paging, nil
}

// GetById implements TaskRepository.
func (t *taskRepository) GetById(Id string) (model.Task, error) {

	var task model.Task
	err := t.db.QueryRow(config.GetTaskById, Id).Scan(&task.Id, &task.Name, &task.Status, &task.Approval, &task.PersonInCharge, &task.Deadline, &task.ProjectId, &task.ApprovalDate, &task.Feedback, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		log.Println("task_repository.QueryRow", err.Error())
		return model.Task{}, err
	}

	return task, nil
}

// GetByPersonInCharge implements TaskRepository.
func (t *taskRepository) GetByPersonInCharge(Id string) ([]model.Task, error) {

	var tasks []model.Task

	row, err := t.db.Query(config.GetTaskByPersonInCharge, Id)
	if err != nil {
		log.Println("task_repository.Query", err.Error())
		return nil, err
	}
	for row.Next() {
		task := model.Task{}
		err := row.Scan(&task.Id, &task.Name, &task.Status, &task.Approval, &task.PersonInCharge, &task.Deadline, &task.ProjectId, &task.ApprovalDate, &task.Feedback, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			log.Println("taskRepository.Rows.Next", err.Error())
			return nil, err
		}

		tasks = append(tasks, task)
	}
	return tasks, nil
}

// GetByProjectId implements TaskRepository.
func (t *taskRepository) GetByProjectId(Id string) ([]model.Task, error) {

	var tasks []model.Task

	row, err := t.db.Query(config.GetTaskByProjectId, Id)
	if err != nil {
		log.Println("task_repository.Query", err.Error())
		return nil, err
	}

	for row.Next() {
		task := model.Task{}
		err := row.Scan(&task.Id, &task.Name, &task.Status, &task.Approval, &task.PersonInCharge, &task.Deadline, &task.ProjectId, &task.ApprovalDate, &task.Feedback, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			log.Println("taskRepository.Rows.Next", err.Error())
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{
		db: db,
	}
}
